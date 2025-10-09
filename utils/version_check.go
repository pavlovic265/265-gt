package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

// VersionCheck struct is now defined in config/config.go
func CheckGTVersion(exe executor.Executor) {
	// Only check once per day to avoid spam
	if !shouldCheckVersion() {
		return // Silently fail if check is not needed
	}

	latestVersion, latestURL, err := getLatestGTVersion()
	if err != nil {
		return // Silently fail if we can't get latest version
	}

	storedVersion := config.Config.GlobalConfig.Version.LastVersion
	if storedVersion == latestVersion {
		return // Silently fail if we have the latest version
	}

	currentVersion, err := getCurrentGTVersion(exe)
	if err != nil {
		currentVersion = "unknown"
	}

	showVersionNotification(currentVersion, latestVersion, latestURL)
	updateVersionCheck(latestVersion)
}

func shouldCheckVersion() bool {
	version := config.Config.GlobalConfig.Version

	// If no last checked time, should check
	if version.LastChecked == "" {
		return true
	}

	lastChecked, err := time.Parse(time.RFC3339, version.LastChecked)
	if err != nil {
		return true // If we can't parse time, check anyway
	}

	// Check if it's been more than 24 hours since last check
	return time.Since(lastChecked) > 24*time.Hour
}

func getCurrentGTVersion(exe executor.Executor) (string, error) {
	output, err := exe.WithName("gt").WithArgs([]string{"version"}).RunWithOutput()
	if err != nil {
		return "", err
	}

	versionStr := strings.TrimSpace(output.String())

	if versionStr != "" {
		return versionStr, nil
	}

	return "", fmt.Errorf("could not parse version from: %s", versionStr)
}

func getLatestGTVersion() (string, string, error) {
	// Get repository from environment variable
	repository := os.Getenv("GT_REPOSITORY")
	if repository == "" {
		return "", "", fmt.Errorf("GT_REPOSITORY environment variable not set")
	}

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repository)
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var release GitHubRelease
	if err := json.Unmarshal(body, &release); err != nil {
		return "", "", err
	}

	return release.TagName, release.HTMLURL, nil
}

func showVersionNotification(current, latest, url string) {
	currentDisplay := strings.TrimPrefix(current, "v")
	latestDisplay := strings.TrimPrefix(latest, "v")

	fmt.Printf("\n🔄 A new release of gt is available: %s → %s\n", currentDisplay, latestDisplay)
	fmt.Printf("To upgrade, run: gt upgrade\n")
	fmt.Printf("%s\n\n", url)
}

func updateVersionCheck(latestVersion string) {
	config.Config.GlobalConfig.Version.LastChecked = time.Now().Format(time.RFC3339)
	config.Config.GlobalConfig.Version.LastVersion = latestVersion

	config.SaveGlobalConfig()
}
