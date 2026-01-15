package helpers

func (gh *GitHelperImpl) SetParent(parent string, child string) error {
	return gh.runner.Git("config", "--local", "gt.branch."+child+".parent", parent)
}

func (gh *GitHelperImpl) GetParent(branch string) (string, error) {
	return gh.runner.GitOutput("config", "--local", "--get", "gt.branch."+branch+".parent")
}

func (gh *GitHelperImpl) DeleteParent(branch string) error {
	return gh.runner.Git("config", "--local", "--unset", "gt.branch."+branch+".parent")
}
