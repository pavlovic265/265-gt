package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	timeutils "github.com/pavlovic265/265-gt/utils/timeutils"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

func CheckGTVersion(exe executor.Executor) {
	if !shouldCheckVersion() {
		return // Silently fail if check is not needed
	}

	// Create context with timeout to avoid blocking
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	latestVersion, latestURL, err := getLatestGTVersionWithContext(ctx)
	if err != nil {
		return // Silently fail if we can't get latest version
	}

	storedVersion := config.Config.GlobalConfig.Version.LastVersion
	if err := config.UpdateLastChecked(); err != nil {
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

	// Compare only date (day, month, year), ignore time
	now := timeutils.Now()
	lastCheckedDate := time.Date(lastChecked.Year(), lastChecked.Month(), lastChecked.Day(), 0, 0, 0, 0, lastChecked.Location())
	currentDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Check if it's a different day
	return !lastCheckedDate.Equal(currentDate)
}

func getLatestGTVersionWithContext(ctx context.Context) (string, string, error) {
	// Get repository from environment variable
	repository := os.Getenv("GT_REPOSITORY")
	if repository == "" {
		return "", "", fmt.Errorf("GT_REPOSITORY environment variable not set")
	}

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repository)

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return "", "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
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

	fmt.Printf("\n%s %s %s %s %s\n",
		config.GetInfoStyle().Render("🔄"),
		config.GetInfoStyle().Render("A new release of gt is available:"),
		config.GetWarningStyle().Render(currentDisplay),
		config.GetDebugStyle().Render(config.ArrowRightIcon),
		config.GetWarningStyle().Render(latestDisplay))

	fmt.Printf("To upgrade, run: %s\n", config.GetSuccessStyle().Render("gt upgrade"))
	fmt.Printf("%s\n\n", config.GetInfoStyle().Render(url))
}
