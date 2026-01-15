package version

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
	"github.com/pavlovic265/265-gt/utils/timeutils"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

func GetLatestVersion() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	version, _, err := getLatestVersionWithContext(ctx)
	return version, err
}

func getLatestVersionWithContext(ctx context.Context) (string, string, error) {
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

func CheckGTVersion(ctx context.Context) {
	cfg, ok := config.GetConfig(ctx)
	if !ok || cfg.Global == nil || cfg.Global.Version == nil {
		return
	}

	v := *cfg.Global.Version
	if !shouldCheckVersion(v) {
		return
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	latestVersion, latestURL, err := getLatestVersionWithContext(timeoutCtx)
	if err != nil {
		return
	}

	storedVersion := v.CurrentVersion

	cfg.Global.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)
	cfg.MarkDirty()

	if len(storedVersion) == 0 {
		showVersionNotification("unknown", latestVersion, latestURL)
		return
	}

	if storedVersion == latestVersion {
		return
	}

	showVersionNotification(storedVersion, latestVersion, latestURL)
}

func shouldCheckVersion(v config.Version) bool {
	if v.LastChecked == "" {
		return true
	}

	lastChecked, err := time.Parse(time.RFC3339, v.LastChecked)
	if err != nil {
		return true
	}

	return time.Since(lastChecked) >= 24*time.Hour
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
