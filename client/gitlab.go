package client

type gitLabCli struct{}

func NewGitLabCli() CliClient {
	return &gitLabCli{}
}

func (svc gitLabCli) CreatePullRequest(args []string) error {
	return nil
}
