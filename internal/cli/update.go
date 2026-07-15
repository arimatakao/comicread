package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const latestReleaseURL = "https://api.github.com/repos/arimatakao/comicread/releases/latest"

var updateHTTPClient = &http.Client{Timeout: 10 * time.Second}

type release struct {
	TagName string `json:"tag_name"`
}

func checkForUpdate() error {
	req, err := http.NewRequest(http.MethodGet, latestReleaseURL, nil)
	if err != nil {
		return fmt.Errorf("create update request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("User-Agent", "comicread/"+Version)

	resp, err := updateHTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("check latest release: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("check latest release: GitHub API returned %s", resp.Status)
	}

	var latest release
	if err := json.NewDecoder(resp.Body).Decode(&latest); err != nil {
		return fmt.Errorf("decode latest release: %w", err)
	}
	if latest.TagName == "" {
		return fmt.Errorf("check latest release: response has no tag name")
	}

	fmt.Printf("Current version: %s\nLatest version: %s\n", Version, latest.TagName)
	if Version == "dev" {
		fmt.Println("This is a development build; install the latest release to update.")
		printUpdateCommand()
		return nil
	}
	if newerVersion(latest.TagName, Version) {
		fmt.Println("An update is available.")
		printUpdateCommand()
		return nil
	}

	fmt.Println("You are up to date.")
	return nil
}

func printUpdateCommand() {
	command := updateCommand()
	if command == "" {
		return
	}
	fmt.Println("Run this command to update:")
	fmt.Println(command)
}

func updateCommand() string {
	language := installerLanguage()
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf(`powershell -ExecutionPolicy Bypass -Command "$env:LANG = '%s'; iwr -useb https://raw.githubusercontent.com/arimatakao/comicread/main/install.ps1 | iex"`, language)
	case "darwin", "linux":
		return fmt.Sprintf("curl -fsSL https://raw.githubusercontent.com/arimatakao/comicread/main/install.sh | LANG=%s bash", language)
	default:
		return ""
	}
}

func installerLanguage() string {
	return os.Getenv("COMICREAD_LANG")
}

func newerVersion(latest, current string) bool {
	latestParts, latestPrerelease, ok := parseVersion(latest)
	if !ok {
		return false
	}
	currentParts, currentPrerelease, ok := parseVersion(current)
	if !ok {
		return false
	}
	for i := range latestParts {
		if latestParts[i] != currentParts[i] {
			return latestParts[i] > currentParts[i]
		}
	}
	if latestPrerelease == currentPrerelease {
		return false
	}
	return latestPrerelease == ""
}

func parseVersion(value string) ([3]int, string, bool) {
	value = strings.TrimSpace(value)
	value = strings.TrimPrefix(value, "v")
	value = strings.TrimPrefix(value, "V")
	value, _, _ = strings.Cut(value, "+")
	value, prerelease, _ := strings.Cut(value, "-")
	fields := strings.Split(value, ".")
	if len(fields) == 0 || len(fields) > 3 {
		return [3]int{}, "", false
	}

	var parts [3]int
	for i, field := range fields {
		if field == "" {
			return [3]int{}, "", false
		}
		part, err := strconv.Atoi(field)
		if err != nil || part < 0 {
			return [3]int{}, "", false
		}
		parts[i] = part
	}
	return parts, prerelease, true
}
