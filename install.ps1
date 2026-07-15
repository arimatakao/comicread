Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

$Repo = "arimatakao/comicread"
$BinName = "comicread"
$VersionInput = "latest"
$AutoYes = $false
$InstallDir = Join-Path -Path $env:LOCALAPPDATA -ChildPath "Programs\comicread"
$ReinstallConfirmed = $false

function Show-Usage {
@"
Usage: powershell -File install.ps1 [--install-dir <path>] [-y|--yes] [version]

Install the latest comicread release, or a specified version.

Examples:
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
                    throw "Error: option '--install-dir' requires a value."
                }
                $i++
                $script:InstallDir = $ArgsList[$i]
            }
            default {
                if ($arg.StartsWith("-")) {
                    throw "Error: unknown option '$arg'."
                }
                $positional += $arg
            }
        }
    }

    if ($positional.Count -gt 1) {
        throw "Error: too many version arguments."
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
        Write-Host "Installation cancelled."
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
            throw "Error: unable to resolve latest release tag."
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
        Confirm-Install "$BinName is already installed (version $installedVersion). Do you want to update to $TargetVersion?"
        $script:ReinstallConfirmed = $true
    }
    elseif ($target -eq $installed) {
        Confirm-Install "$BinName is already installed (version $installedVersion). Do you want to reinstall $TargetVersion?"
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
        default { throw "Error: unsupported Windows architecture '$arch'." }
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
    throw "Error: no Windows zip asset found for arch '$arch' in release '$Version'."
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
        $value = Read-Host "  $Name ($($AllowedValues -join ' '); leave blank to skip)"
        if ([string]::IsNullOrWhiteSpace($value)) {
            return $false
        }
        if ($AllowedValues -contains $value) {
            [Environment]::SetEnvironmentVariable($Name, $value, "User")
            return $true
        }
        Write-Host "Invalid value. Choose one of: $($AllowedValues -join ' ')"
    }
}

function Configure-Environment {
    if ($script:AutoYes -or -not (Ask-YesNo "Configure comicread environment variables?")) {
        return $false
    }

    Write-Host "Values will be saved in your user environment."
    Write-Host "COMICREAD_GRAPHICS chooses how pages are rendered; auto detects terminal support."
    Write-Host "COMICREAD_VIEW chooses the default page layout; leave it blank for single-page view."
    Write-Host "COMICREAD_LANG chooses the language of the interface; the default is en."

    $changed = $false
    if (Configure-EnvironmentValue -Name "COMICREAD_GRAPHICS" -AllowedValues @("auto", "ascii", "dots", "kitty", "sixel", "iterm2")) { $changed = $true }
    if (Configure-EnvironmentValue -Name "COMICREAD_VIEW" -AllowedValues @("book-view", "right-view", "circle-view", "right-circle-view")) { $changed = $true }
    if (Configure-EnvironmentValue -Name "COMICREAD_LANG" -AllowedValues @("en", "uk", "pl", "de", "fr", "es", "cs", "ro", "it", "ko", "ja", "id", "hi", "el", "tr")) { $changed = $true }
    return $changed
}

function Main {
    param([string[]]$CliArgs)

    Parse-Args -ArgsList $CliArgs
    $version = Resolve-Version
    $url = Find-ZipAssetUrl -Version $version
    $fileName = Split-Path -Path $url -Leaf

    Confirm-UpgradeIfNeeded -TargetVersion $version
    if (-not $script:ReinstallConfirmed) {
        Confirm-Install "Install $BinName $version to $InstallDir?"
    }

    $tempDir = Join-Path -Path ([System.IO.Path]::GetTempPath()) -ChildPath ("comicread-install-" + [guid]::NewGuid().ToString("N"))
    $zipPath = Join-Path -Path $tempDir -ChildPath $fileName
    $extractDir = Join-Path -Path $tempDir -ChildPath "extract"

    try {
        New-Item -ItemType Directory -Path $tempDir -Force | Out-Null
        Write-Host "Downloading $fileName..."
        Download-File -Url $url -OutFile $zipPath
        Unblock-File -Path $zipPath -ErrorAction SilentlyContinue

        Write-Host "Extracting $fileName..."
        Expand-Archive -Path $zipPath -DestinationPath $extractDir -Force
        $sourceExe = Join-Path -Path $extractDir -ChildPath "$BinName.exe"
        if (-not (Test-Path -Path $sourceExe -PathType Leaf)) {
            throw "Error: '$BinName.exe' was not found in archive '$fileName'."
        }

        New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
        $targetExe = Join-Path -Path $InstallDir -ChildPath "$BinName.exe"
        Copy-Item -Path $sourceExe -Destination $targetExe -Force
        $pathChanged = Ensure-UserPathContains -Dir $InstallDir
        $environmentChanged = Configure-Environment

        Write-Host "comicread $version installed to $targetExe"
        if ($pathChanged -or $environmentChanged) {
            Write-Host "Restart your terminal to use comicread."
        }
    }
    finally {
        Remove-Item -Path $tempDir -Recurse -Force -ErrorAction SilentlyContinue
    }
}

Main -CliArgs $args
