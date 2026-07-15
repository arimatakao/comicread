Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

$Repo = "arimatakao/comicread"
$BinName = "comicread"
$VersionInput = "latest"
$AutoYes = $false
$InstallDir = Join-Path -Path $env:LOCALAPPDATA -ChildPath "Programs\comicread"
$ReinstallConfirmed = $false
$LocaleBaseUrl = "https://raw.githubusercontent.com/$Repo/main/installer/locales"
$Messages = @{}

function ConvertFrom-Properties {
    param([string]$Content)

    $properties = @{}
    foreach ($line in ($Content -split "`r?`n")) {
        if ([string]::IsNullOrWhiteSpace($line) -or $line.StartsWith("#")) {
            continue
        }
        $separator = $line.IndexOf("=")
        if ($separator -lt 1) {
            continue
        }
        $properties[$line.Substring(0, $separator)] = $line.Substring($separator + 1)
    }
    return $properties
}

function Get-RemoteLocale {
    param([string]$Language)

    try {
        return (Invoke-WebRequest -Uri "$script:LocaleBaseUrl/$Language.properties" -UseBasicParsing).Content
    }
    catch {
        return $null
    }
}

function Initialize-InstallerLocale {
    $language = $env:LANG
    if ([string]::IsNullOrWhiteSpace($language)) {
        $language = [Globalization.CultureInfo]::CurrentUICulture.TwoLetterISOLanguageName
    }
    $language = ($language -split "[_@.]")[0].ToLowerInvariant()
    $content = $null
    if ($PSScriptRoot) {
        $localPath = Join-Path -Path $PSScriptRoot -ChildPath "installer\locales\$language.properties"
        if (Test-Path -Path $localPath -PathType Leaf) {
            $content = Get-Content -Path $localPath -Raw -Encoding utf8
        }
        if (-not $content) {
            $fallbackPath = Join-Path -Path $PSScriptRoot -ChildPath "installer\locales\en.properties"
            if (Test-Path -Path $fallbackPath -PathType Leaf) {
                $content = Get-Content -Path $fallbackPath -Raw -Encoding utf8
            }
        }
    }
    if (-not $content) {
        $content = Get-RemoteLocale -Language $language
    }
    if (-not $content -and $language -ne "en") {
        $content = Get-RemoteLocale -Language "en"
    }
    if ($content) {
        $script:Messages = ConvertFrom-Properties -Content $content
    }
}

function T {
    param(
        [string]$Key,
        [object[]]$Values = @()
    )

    $message = $script:Messages[$Key]
    if ([string]::IsNullOrEmpty($message)) {
        return $Key
    }
    if ($Values.Count -eq 0) {
        return $message
    }
    return [string]::Format($message, $Values)
}

function Show-Usage {
@"
$(T "usage" "powershell -File install.ps1 [--install-dir <path>] [-y|--yes] [version]")

$(T "usage.description")

$(T "usage.examples")
  powershell -File install.ps1
  powershell -File install.ps1 v1.2.3
  powershell -File install.ps1 --install-dir "`$env:USERPROFILE\bin"
  powershell -File install.ps1 --yes
"@ | Write-Host
}

function Parse-Args {
    param([string[]]$ArgsList)

    $positional = @()
    for ($i = 0; $i -lt $ArgsList.Count; $i++) {
        $arg = $ArgsList[$i]
        switch ($arg) {
            "-h" { Show-Usage; exit 0 }
            "--help" { Show-Usage; exit 0 }
            "-y" { $script:AutoYes = $true }
            "--yes" { $script:AutoYes = $true }
            "--install-dir" {
                if ($i + 1 -ge $ArgsList.Count) {
                    throw (T "error.option_value" "--install-dir")
                }
                $i++
                $script:InstallDir = $ArgsList[$i]
            }
            default {
                if ($arg.StartsWith("-")) {
                    throw (T "error.unknown_option" $arg)
                }
                $positional += $arg
            }
        }
    }

    if ($positional.Count -gt 1) {
        throw (T "error.too_many_versions")
    }
    if ($positional.Count -eq 1) {
        $script:VersionInput = $positional[0]
    }
}

function Confirm-Install {
    param([string]$Message)

    if ($script:AutoYes) {
        return
    }

    $answer = Read-Host "$Message [y/N]"
    if ($answer -notmatch "^(y|yes)$") {
        Write-Host (T "cancelled")
        exit 0
    }
}

function Ask-YesNo {
    param([string]$Message)

    $answer = Read-Host "$Message [y/N]"
    return $answer -match "^(y|yes)$"
}

function Resolve-Version {
    if ($script:VersionInput -eq "latest") {
        $release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
        if (-not $release.tag_name) {
            throw (T "error.latest_version")
        }
        return [string]$release.tag_name
    }

    if ($script:VersionInput.StartsWith("v")) {
        return $script:VersionInput
    }
    return "v$($script:VersionInput)"
}

function Convert-ToVersionObject {
    param([string]$Value)

    if ([string]::IsNullOrWhiteSpace($Value)) {
        return $null
    }

    $normalized = ($Value.Trim() -replace "^[vV]", "").Split("-")[0]
    $parts = $normalized.Split(".")
    if ($parts.Count -eq 0 -or $parts.Count -gt 4 -or ($parts | Where-Object { $_ -notmatch "^\d+$" })) {
        return $null
    }

    $padded = @($parts)
    while ($padded.Count -lt 4) {
        $padded += "0"
    }
    return [version]::Parse(($padded -join "."))
}

function Get-InstalledVersion {
    $candidates = @()
    $localExe = Join-Path -Path $script:InstallDir -ChildPath "$BinName.exe"
    if (Test-Path -Path $localExe -PathType Leaf) {
        $candidates += $localExe
    }
    $command = Get-Command -Name $BinName -ErrorAction SilentlyContinue
    if ($command) {
        $candidates += $command.Source
    }

    foreach ($candidate in ($candidates | Select-Object -Unique)) {
        try {
            $output = & $candidate --version 2>$null
            $match = [regex]::Match(($output | Out-String), "v?\d+(\.\d+){1,3}([-.+][0-9A-Za-z.-]+)?")
            if ($match.Success) {
                return $match.Value
            }
        }
        catch {
            continue
        }
    }
    return $null
}

function Confirm-UpgradeIfNeeded {
    param([string]$TargetVersion)

    $installedVersion = Get-InstalledVersion
    if (-not $installedVersion) {
        return
    }

    $target = Convert-ToVersionObject -Value $TargetVersion
    $installed = Convert-ToVersionObject -Value $installedVersion
    if (-not $target -or -not $installed) {
        return
    }
    if ($target -gt $installed) {
        Confirm-Install (T "confirm.update" $installedVersion $TargetVersion)
        $script:ReinstallConfirmed = $true
    }
    elseif ($target -eq $installed) {
        Confirm-Install (T "confirm.reinstall" $installedVersion $TargetVersion)
        $script:ReinstallConfirmed = $true
    }
}

function Get-WindowsArch {
    $arch = $env:PROCESSOR_ARCHITEW6432
    if ([string]::IsNullOrWhiteSpace($arch)) {
        $arch = $env:PROCESSOR_ARCHITECTURE
    }

    switch ($arch) {
        "AMD64" { return "amd64" }
        "ARM64" { return "arm64" }
        "x86" { return "386" }
        default { throw (T "error.unsupported_arch" $arch) }
    }
}

function Find-ZipAssetUrl {
    param([string]$Version)

    $release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/tags/$Version"
    $arch = Get-WindowsArch
    $expected = "${BinName}_${Version}_windows_${arch}.zip"

    foreach ($asset in $release.assets) {
        if ($asset.name -eq $expected) {
            return [string]$asset.browser_download_url
        }
    }
    throw (T "error.zip_asset" $arch $Version)
}

function Download-File {
    param(
        [string]$Url,
        [string]$OutFile
    )

    $previousProgressPreference = $ProgressPreference
    $ProgressPreference = "SilentlyContinue"
    try {
        Invoke-WebRequest -Uri $Url -OutFile $OutFile -UseBasicParsing
    }
    finally {
        $ProgressPreference = $previousProgressPreference
    }
}

function Ensure-UserPathContains {
    param([string]$Dir)

    $current = [Environment]::GetEnvironmentVariable("Path", "User")
    if ([string]::IsNullOrWhiteSpace($current)) {
        [Environment]::SetEnvironmentVariable("Path", $Dir, "User")
        return $true
    }

    $parts = $current.Split(";") | Where-Object { $_ -ne "" }
    if ($parts -contains $Dir) {
        return $false
    }

    [Environment]::SetEnvironmentVariable("Path", "$current;$Dir", "User")
    return $true
}

function Configure-EnvironmentValue {
    param(
        [string]$Name,
        [string[]]$AllowedValues
    )

    while ($true) {
        $prompt = T -Key "prompt.value" -Values @($Name, ($AllowedValues -join ' '))
        $value = Read-Host $prompt
        if ([string]::IsNullOrWhiteSpace($value)) {
            return $false
        }
        if ($AllowedValues -contains $value) {
            [Environment]::SetEnvironmentVariable($Name, $value, "User")
            return $true
        }
        Write-Host (T "error.invalid_value" ($AllowedValues -join ' '))
    }
}

function Configure-Environment {
    if ($script:AutoYes -or -not (Ask-YesNo (T "environment.configure"))) {
        return $false
    }

    Write-Host (T "environment.saved.user")
    Write-Host (T "environment.graphics")
    Write-Host (T "environment.view")
    Write-Host (T "environment.language")

    $changed = $false
    if (Configure-EnvironmentValue -Name "COMICREAD_GRAPHICS" -AllowedValues @("auto", "ascii", "dots", "kitty", "sixel", "iterm2")) { $changed = $true }
    if (Configure-EnvironmentValue -Name "COMICREAD_VIEW" -AllowedValues @("book-view", "right-view", "circle-view", "right-circle-view")) { $changed = $true }
    if (Configure-EnvironmentValue -Name "COMICREAD_LANG" -AllowedValues @("en", "uk", "pl", "de", "fr", "es", "cs", "ro", "it", "ko", "ja", "id", "hi", "el", "tr")) { $changed = $true }
    return $changed
}

function Main {
    param([string[]]$CliArgs)

    Initialize-InstallerLocale
    Parse-Args -ArgsList $CliArgs
    $version = Resolve-Version
    $url = Find-ZipAssetUrl -Version $version
    $fileName = Split-Path -Path $url -Leaf

    Confirm-UpgradeIfNeeded -TargetVersion $version
    if (-not $script:ReinstallConfirmed) {
        Confirm-Install (T "confirm.install" $version $InstallDir)
    }

    $tempDir = Join-Path -Path ([System.IO.Path]::GetTempPath()) -ChildPath ("comicread-install-" + [guid]::NewGuid().ToString("N"))
    $zipPath = Join-Path -Path $tempDir -ChildPath $fileName
    $extractDir = Join-Path -Path $tempDir -ChildPath "extract"

    try {
        New-Item -ItemType Directory -Path $tempDir -Force | Out-Null
        Write-Host (T "status.downloading" $fileName)
        Download-File -Url $url -OutFile $zipPath
        Unblock-File -Path $zipPath -ErrorAction SilentlyContinue

        Write-Host (T "status.extracting" $fileName)
        Expand-Archive -Path $zipPath -DestinationPath $extractDir -Force
        $sourceExe = Join-Path -Path $extractDir -ChildPath "$BinName.exe"
        if (-not (Test-Path -Path $sourceExe -PathType Leaf)) {
            throw (T "error.binary_missing" "$BinName.exe")
        }

        New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
        $targetExe = Join-Path -Path $InstallDir -ChildPath "$BinName.exe"
        Copy-Item -Path $sourceExe -Destination $targetExe -Force
        $pathChanged = Ensure-UserPathContains -Dir $InstallDir
        $environmentChanged = Configure-Environment

        Write-Host (T "status.installed" $version $targetExe)
        if ($pathChanged -or $environmentChanged) {
            Write-Host (T "status.restart")
        }
    }
    finally {
        Remove-Item -Path $tempDir -Recurse -Force -ErrorAction SilentlyContinue
    }
}

Main -CliArgs $args
