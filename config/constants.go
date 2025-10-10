package config

var FileName = ".gtconfig.yaml"

type Platform string

var (
	GitHubPlatform Platform = "GitHub"
	GitLabPlatform Platform = "GitLab"
)

func (p Platform) String() string {
	switch p {
	case GitHubPlatform:
		return "GitHub"
	case GitLabPlatform:
		return "GitLab"
	default:
		return ""
	}
}
