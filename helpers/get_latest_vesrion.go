package helpers

import (
	"encoding/json"
	"net/http"
)

func GetLatestVersion() (string, error) {
	url := "https://api.github.com/repos/pavlovic265/265-gt/releases/latest"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	var result struct {
		TagName string `json:"tag_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.TagName, nil
}
