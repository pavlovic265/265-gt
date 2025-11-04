package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/executor"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/pavlovic265/265-gt/utils/pointer"
	"github.com/pavlovic265/265-gt/utils/timeutils"
	"github.com/spf13/cobra"
)

type upgradeCommand struct {
	exe           executor.Executor
	configManager config.ConfigManager
}

func NewUpgradeCommand(
	exe executor.Executor,
	configManager config.ConfigManager,
) upgradeCommand {
	return upgradeCommand{
		exe:           exe,
		configManager: configManager,
	}
}

func (svc upgradeCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade of current build",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			svc.configManager.InitGlobalConfig()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			version, isLatest, err := svc.isLatestVersion()
			if err != nil {
				return err
			}
			if isLatest {
				log.Infof("You are already on the latest version: %s", pointer.Deref(version))
				return nil
			}

			binary := svc.checkWhichBinary()
			if binary == nil {
				return log.ErrorMsg("Failed to check if homebrew is installed")
			}

			switch pointer.Deref(binary) {
			case BinaryHomebrew:
				if err := svc.upgradeWithHomebrew(); err != nil {
					return err
				}
			case BinaryScript:
				if err := svc.upgradeWithScript(); err != nil {
					return err
				}
			}

			globalConfig, err := svc.configManager.LoadGlobalConfig()
			if err != nil {
				return log.Error("Global config not found. Run 'gt config global' to create it first", err)
			}

			globalConfig.Version.LastChecked = timeutils.Now().Format(timeutils.LayoutISOWithTime)
			globalConfig.Version.CurrentVersion = pointer.Deref(version)

			if err := svc.configManager.SaveGlobalConfig(*globalConfig); err != nil {
				return log.Error("Failed to update version in config", err)
			}

			log.Success("Tool upgraded successfully")
			return nil
		},
	}
}

type Binary string

const (
	BinaryHomebrew Binary = "homebrew"
	BinaryScript   Binary = "script"
)

func (svc upgradeCommand) checkWhichBinary() *Binary {
	exeArgs := []string{"-v", "gt"}
	out, err := svc.exe.WithName("command").WithArgs(exeArgs).RunWithOutput()
	if err != nil {
		log.Warningf("Failed to check if homebrew is installed: %v", err)
		return nil
	}
	if strings.Contains(out.String(), "homebrew") {
		return pointer.From(BinaryHomebrew)
	}

	return pointer.From(BinaryScript)
}

func (svc upgradeCommand) upgradeWithHomebrew() error {
	exeArgs := []string{"brew", "upgrade", "gt"}
	err := svc.exe.WithName("bash").WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}

	return nil
}

func (svc upgradeCommand) upgradeWithScript() error {
	installURL := "https://raw.githubusercontent.com/pavlovic265/265-gt/main/scripts/install.sh"
	exeArgs := []string{"-c", fmt.Sprintf("curl -fsSL %s | bash", installURL)}
	err := svc.exe.WithName("bash").WithArgs(exeArgs).Run()
	if err != nil {
		return err
	}

	return nil
}

func (svc upgradeCommand) isLatestVersion() (*string, bool, error) {
	// Get latest version from GitHub API
	url := "https://api.github.com/repos/pavlovic265/265-gt/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		return nil, false, err
	}
	defer func() { _ = resp.Body.Close() }()

	var result struct {
		TagName string `json:"tag_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, false, err
	}

	// Get current version from config manager
	globalConfig, err := svc.configManager.LoadGlobalConfig()
	if err != nil {
		return nil, false, err
	}
	currentVersion := globalConfig.Version.CurrentVersion
	if currentVersion == result.TagName {
		return pointer.From(result.TagName), true, nil
	}

	return pointer.From(result.TagName), false, nil
}
