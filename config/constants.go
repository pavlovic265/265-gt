package config

var FileName = ".gtconfig.yaml"

type platform string

var (
	GitHubPlatform platform = "GitHub"
	GitLabPlatform platform = "GitLab"
)

func (p platform) String() string {
	switch p {
	case GitHubPlatform:
		return "GitHub"
	case GitLabPlatform:
		return "GitLab"
	default:
		return ""
	}
}
