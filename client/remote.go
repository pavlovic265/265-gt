package client

import (
	"fmt"
	"regexp"
	"strings"
)

type RepoInfo struct {
	Owner    string
	Repo     string
	Platform string
}

func ParseRemoteURL(remoteURL string) (*RepoInfo, error) {
	// SSH format: git@github.com:owner/repo.git
	sshRegex := regexp.MustCompile(`git@([^:]+):([^/]+)/(.+?)(?:\.git)?$`)
	if matches := sshRegex.FindStringSubmatch(remoteURL); matches != nil {
		return &RepoInfo{
			Platform: matches[1],
			Owner:    matches[2],
			Repo:     strings.TrimSuffix(matches[3], ".git"),
		}, nil
	}

	// HTTPS format: https://github.com/owner/repo.git
	httpsRegex := regexp.MustCompile(`https?://([^/]+)/([^/]+)/(.+?)(?:\.git)?$`)
	if matches := httpsRegex.FindStringSubmatch(remoteURL); matches != nil {
		return &RepoInfo{
			Platform: matches[1],
			Owner:    matches[2],
			Repo:     strings.TrimSuffix(matches[3], ".git"),
		}, nil
	}

	return nil, fmt.Errorf("unable to parse remote URL: %s", remoteURL)
}
