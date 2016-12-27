package sollozzoctl

import (
	"testing"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestAddCommand(t *testing.T) {

	var testRootCmd = &cobra.Command{}
	testRootCmd.AddCommand(addCmd)

	assert.NotEmpty(t, testRootCmd)
	assert.True(t, testRootCmd.HasAvailableSubCommands())
	assert.Equal(t, 1, len(testRootCmd.Commands()))

	addCommand := testRootCmd.Commands()[0]

	assert.NotEmpty(t, addCommand)
	assert.Equal(t, addCmd.Use, addCommand.Use)
	assert.Equal(t, addCmd.Short, addCommand.Short)
	assert.Equal(t, addCmd.Long, addCommand.Long)
}

func TestListCommand(t *testing.T) {

	var testRootCmd = &cobra.Command{}
	testRootCmd.AddCommand(listCmd)

	assert.NotEmpty(t, testRootCmd)
	assert.True(t, testRootCmd.HasAvailableSubCommands())
	assert.Equal(t, 1, len(testRootCmd.Commands()))

	listCommand := testRootCmd.Commands()[0]

	assert.NotEmpty(t, listCommand)
	assert.Equal(t, listCmd.Use, listCommand.Use)
	assert.Equal(t, listCmd.Short, listCommand.Short)
	assert.Equal(t, listCmd.Long, listCommand.Long)
}

func TestRemoveCommand(t *testing.T) {

	var testRootCmd = &cobra.Command{}
	testRootCmd.AddCommand(removeCmd)

	assert.NotEmpty(t, testRootCmd)
	assert.True(t, testRootCmd.HasAvailableSubCommands())
	assert.Equal(t, 1, len(testRootCmd.Commands()))

	removeCommand := testRootCmd.Commands()[0]

	assert.NotEmpty(t, removeCommand)
	assert.Equal(t, removeCmd.Use, removeCommand.Use)
	assert.Equal(t, removeCmd.Short, removeCommand.Short)
	assert.Equal(t, removeCmd.Long, removeCommand.Long)
}

func TestCurrentCommand(t *testing.T)  {
	var testRootCmd  =&cobra.Command{}
	testRootCmd.AddCommand(currentCmd)

	assert.NotEmpty(t,testRootCmd)
	assert.True(t, testRootCmd.HasAvailableSubCommands())
	assert.Equal(t, 1, len(testRootCmd.Commands()))

	currentCommand:=testRootCmd.Commands()[0]

	assert.NotEmpty(t,currentCommand)
	assert.Equal(t,currentCmd.Use,currentCommand.Use)
	assert.Equal(t,currentCmd.Short,currentCommand.Short)
	assert.Equal(t,currentCmd.Long,currentCommand.Long)
}

func TestReleaseCommand(t *testing.T)  {
	var testRootCmd  =&cobra.Command{}
	testRootCmd.AddCommand(releaseCmd)

	assert.NotEmpty(t,testRootCmd)
	assert.True(t, testRootCmd.HasAvailableSubCommands())
	assert.Equal(t, 1, len(testRootCmd.Commands()))

	releaseCommand:=testRootCmd.Commands()[0]

	assert.NotEmpty(t,releaseCommand)
	assert.Equal(t,releaseCmd.Use,releaseCommand.Use)
	assert.Equal(t,releaseCmd.Short,releaseCommand.Short)
	assert.Equal(t,releaseCmd.Long,releaseCommand.Long)
}
func TestAddAllCommand(t *testing.T){
	var testRootCmd  =&cobra.Command{}
	testRootCmd.AddCommand(addCmd)
	testRootCmd.AddCommand(currentCmd)
	testRootCmd.AddCommand(listCmd)
	testRootCmd.AddCommand(releaseCmd)
	testRootCmd.AddCommand(removeCmd)

	assert.NotEmpty(t,testRootCmd)
	assert.True(t, testRootCmd.HasAvailableSubCommands())
	assert.Equal(t, 5, len(testRootCmd.Commands()))
}