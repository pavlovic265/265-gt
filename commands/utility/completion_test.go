package utility_test

import (
	"testing"

	"github.com/pavlovic265/265-gt/commands/utility"
	"github.com/stretchr/testify/assert"
)

func TestCompletionCommand_Command(t *testing.T) {
	completionCmd := utility.NewCompletionCommand()
	cmd := completionCmd.Command()

	assert.Equal(t, "completion [bash|zsh|fish|powershell]", cmd.Use)
	assert.Equal(t, "Generate shell completion scripts", cmd.Short)
	assert.Len(t, cmd.ValidArgs, 4)
	assert.Contains(t, cmd.ValidArgs, "bash")
	assert.Contains(t, cmd.ValidArgs, "zsh")
	assert.Contains(t, cmd.ValidArgs, "fish")
	assert.Contains(t, cmd.ValidArgs, "powershell")
}

func TestNewCompletionCommand(t *testing.T) {
	completionCmd := utility.NewCompletionCommand()
	cmd := completionCmd.Command()

	assert.NotNil(t, cmd)
	assert.Equal(t, "completion [bash|zsh|fish|powershell]", cmd.Use)
}
