package helpers

import "github.com/pavlovic265/265-gt/config"

func IsProtectedBranch(branch string) bool {
	for _, pBranch := range config.Config.LocalConfig.Protected {
		if branch == pBranch {
			return true
		}
	}
	return false
}
