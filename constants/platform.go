package constants

var FileName = ".gtconfig.yaml"

type Platform string

var (
	GitHubPlatform Platform = "GitHub"
	GitLabPlatform Platform = "GitLab"
)

const (
	GitHubHost        = "github.com"
	GitLabHost        = "gitlab.com"
	GitHubNoReplyMail = "@users.noreply.github.com"
	GitLabNoReplyMail = "@users.noreply.gitlab.com"
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
