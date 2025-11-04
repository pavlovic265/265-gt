package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/constants"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/pavlovic265/265-gt/utils/timeutils"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

func (gh *GitHelperImpl) CheckGTVersion() {
	version := gh.configManager.GetVersion()
	if !shouldCheckVersion(version) {
		return // Silently fail if check is not needed
	}

	// Create context with timeout to avoid blocking (reduced to 1s for faster commands)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	latestVersion, latestURL, err := getLatestGTVersionWithContext(ctx)
	if err != nil {
		return // Silently fail if we can't get latest version
	}

	storedVersion := version.CurrentVersion
	globalConfig, err := gh.configManager.LoadGlobalConfig()
	if err != nil {
		_ = log.Error("Global config not found. Run 'gt config global' to create it first", err)
		return
	}

	globalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)

	if err := gh.configManager.SaveGlobalConfig(*globalConfig); err != nil {
		return // Silently fail if we can't update last checked
	}

	// If no version is stored, show notification to upgrade
	if len(storedVersion) == 0 {
		showVersionNotification("unknown", latestVersion, latestURL)
		return
	}

	// If we have the latest version, don't show notification
	if storedVersion == latestVersion {
		return
	}

	showVersionNotification(storedVersion, latestVersion, latestURL)
}

func shouldCheckVersion(version config.Version) bool {
	// If no last checked time, should check
	if version.LastChecked == "" {
		return true
	}

	lastChecked, err := time.Parse(time.RFC3339, version.LastChecked)
	if err != nil {
		return true // If we can't parse time, check anyway
	}

	// Check if it's been at least 24 hours since last check
	return time.Since(lastChecked) >= 24*time.Hour
}

func getLatestGTVersionWithContext(ctx context.Context) (string, string, error) {
	apiURL := "https://api.github.com/repos/pavlovic265/265-gt/releases/latest"

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return "", "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer func() { _ = resp.Body.Close() }()

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

	fmt.Printf("\n%s %s %s %s %s\n",
		"🔄",
		"A new release of gt is available:",
		currentDisplay,
		constants.ArrowRightIcon,
		latestDisplay)

	fmt.Printf("To upgrade, run: %s\n", "gt upgrade")
	fmt.Printf("%s\n\n", url)
}
