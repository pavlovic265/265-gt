package account

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/config"
	"github.com/pavlovic265/265-gt/ui/theme"
	"github.com/pavlovic265/265-gt/utils/log"
	"github.com/spf13/cobra"
)

var (
	headerStyle = lipgloss.NewStyle().
			Foreground(theme.Blue).
			Bold(true)

	userStyle = lipgloss.NewStyle().
			Foreground(theme.Green)

	platformStyle = lipgloss.NewStyle().
			Foreground(theme.Cyan)

	emailStyle = lipgloss.NewStyle().
			Foreground(theme.BrightBlack)

	activeStyle = lipgloss.NewStyle().
			Foreground(theme.Yellow).
			Bold(true)
)

type listCommand struct {
	configManager config.ConfigManager
}

func NewListCommand(
	configManager config.ConfigManager,
) listCommand {
	return listCommand{
		configManager: configManager,
	}
}

func (lc listCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all accounts",
		Long:    "List all configured GitHub and GitLab accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.RequireGlobal(cmd.Context())
			if err != nil {
				return err
			}

			if len(cfg.Global.Accounts) == 0 {
				log.Info("No accounts configured")
				fmt.Println("\nRun 'gt account add' to add an account")
				return nil
			}

			fmt.Println(headerStyle.Render("Accounts"))
			fmt.Println()

			activeAccount := cfg.Global.ActiveAccount

			for i, account := range cfg.Global.Accounts {
				isActive := activeAccount != nil &&
					activeAccount.User == account.User &&
					activeAccount.Platform == account.Platform

				prefix := "  "
				if isActive {
					prefix = activeStyle.Render("* ")
				}

				fmt.Printf("%s%s %s %s",
					prefix,
					userStyle.Render(account.User),
					platformStyle.Render("("+account.Platform.String()+")"),
					emailStyle.Render(account.Email))

				if account.Name != "" {
					fmt.Printf(" - %s", account.Name)
				}

				fmt.Println()

				// Add spacing between accounts
				if i < len(cfg.Global.Accounts)-1 {
					fmt.Println()
				}
			}

			fmt.Println()
			log.Infof("Total: %d account(s)", len(cfg.Global.Accounts))

			return nil
		},
	}
}
