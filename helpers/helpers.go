package helpers

import (
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/helpers/accounthelper"
	"github.com/pavlovic265/265-gt/helpers/githelper"
	"github.com/pavlovic265/265-gt/helpers/sshhelper"
	"github.com/pavlovic265/265-gt/runner"
)

type GitHelper = githelper.GitHelper

type SSHKeyType = sshhelper.SSHKeyType

const (
	SSHKeyTypeEd25519 = sshhelper.SSHKeyTypeEd25519
	SSHKeyTypeRSA     = sshhelper.SSHKeyTypeRSA
)

type SSHHelper = sshhelper.SSHHelper

func NewGitHelper(runner runner.Runner) GitHelper {
	return githelper.NewGitHelper(runner)
}

func NewSSHHelper(runner runner.Runner) SSHHelper {
	return sshhelper.NewSSHHelper(runner)
}

func BuildSSHHost(platform, username string) string {
	return sshhelper.BuildSSHHost(platform, username)
}

func DefaultSSHKeyPath(platform, username string) string {
	return sshhelper.DefaultSSHKeyPath(platform, username)
}

func AttachAccountToDir(account *config.Account, targetDir string) error {
	return accounthelper.AttachAccountToDir(account, targetDir)
}
