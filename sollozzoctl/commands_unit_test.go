package sollozzoctl_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
	"github.com/yasinKIZILKAYA/sollozzo/sollozzoctl"
	"strconv"
)

const db = "test.db"

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func NewTestStore() *boltdb.Store {

	store := boltdb.NewStore(db)

	//Removing existing test db
	exist, err := exists(store.Path())

	if exist {
		err = os.Remove(store.Path())
	}

	if err != nil {
		log.Panic(err)
	}

	//Open new db
	err = store.Open()

	if err != nil {
		log.Panic(err)
	}

	return store
}

// add command tests

func TestAddCommandWithNoArgumentExpectArgumentError(t *testing.T) {

	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewAddCommand(store)

	err := cmd.RunE(cmd, []string{})

	assert.Error(t, err)
	assert.Equal(t, "\"sollozzo add\" accepts only project name and project version(optional) arguments.", err.Error())
}

func TestNewAddCommandOnlyProjectNameArgument(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewAddCommand(store)

	err := cmd.RunE(cmd, []string{"abc"})

	assert.NoError(t, err)
}

func TestNewAddCommandProjectNameArgumentWithVersion(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewAddCommand(store)

	err := cmd.RunE(cmd, []string{"abc", "1.2.3"})

	assert.NoError(t, err)
}

func TestNewAddCommandWithMoreArgument(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewAddCommand(store)

	err := cmd.RunE(cmd, []string{"abc", "1.2.3", "kml"})

	assert.Error(t, err)
	assert.Equal(t, "\"sollozzo add\" accepts only project name and project version(optional) arguments.", err.Error())
}

func TestNewAddCommandProjectNameArgumentWithInvalidVersion(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewAddCommand(store)

	err := cmd.RunE(cmd, []string{"abc", "1.2."})

	assert.Error(t, err)
	assert.Equal(t, "Project can not created. Invalid version number, try as 1.2.3", err.Error())
}

func TestNewAddCommandProjectIsExist(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewAddCommand(store)

	err := cmd.RunE(cmd, []string{"abc"})
	assert.NoError(t, err)

	err = cmd.RunE(cmd, []string{"abc", "1.2.3"})

	assert.Error(t, err)
	assert.Equal(t, "Project exist. Please remove availabile project firstly", err.Error())
}

// current command tests

func TestNewCurrentCommand(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)
	err := addCmd.RunE(addCmd, []string{"abc"})

	assert.NoError(t, err)

	currentCmd := sollozzoctl.NewCurrentCommand(store)

	err = currentCmd.RunE(currentCmd, []string{"abc"})

	assert.NoError(t, err)
}

func TestNewCurrentCommandWithNoArgument(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewCurrentCommand(store)

	err := cmd.RunE(cmd, []string{})

	assert.Error(t, err)
	assert.Equal(t, "\"sollozzo current\" accepts only project name argument.", err.Error())
}

func TestNewCurrentCommandWithMoreArguments(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewCurrentCommand(store)

	err := cmd.RunE(cmd, []string{"abc", "def"})

	assert.Error(t, err)
	assert.Equal(t, "\"sollozzo current\" accepts only project name argument.", err.Error())
}

func TestNewCurrentCommandGetNotExistProject(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	currentCmd := sollozzoctl.NewCurrentCommand(store)

	err := currentCmd.RunE(currentCmd, []string{"abc"})

	assert.Error(t, err)
	assert.Equal(t, "Project not available", err.Error())
}

// list command tests

func TestNewListCommand(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)

	err := addCmd.RunE(addCmd, []string{"abc", "4.5.6"})
	assert.NoError(t, err)

	err = addCmd.RunE(addCmd, []string{"def"})
	assert.NoError(t, err)

	listCmd := sollozzoctl.NewListCommand(store)
	err = listCmd.RunE(listCmd, []string{})
	assert.NoError(t, err)
}

func TestNewListCommandWithArgument(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)

	err := addCmd.RunE(addCmd, []string{"abc", "4.5.6"})
	assert.NoError(t, err)

	err = addCmd.RunE(addCmd, []string{"def"})
	assert.NoError(t, err)

	listCmd := sollozzoctl.NewListCommand(store)
	err = listCmd.RunE(listCmd, []string{"abc"})

	assert.Error(t, err)
	assert.Equal(t, "\"sollozzo list\" accepts no argument(s).", err.Error())
}

func TestNewListCommandNoProject(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	listCmd := sollozzoctl.NewListCommand(store)
	err := listCmd.RunE(listCmd, []string{})

	assert.Error(t, err)
	assert.Equal(t, "You do not have a project yet", err.Error())
}

// remove command test

func TestNewRemoveCommand(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)

	err := addCmd.RunE(addCmd, []string{"abc", "4.5.6"})
	assert.NoError(t, err)

	removeCmd := sollozzoctl.NewRemoveCommand(store)
	err = removeCmd.RunE(removeCmd, []string{"abc"})
	assert.NoError(t, err)
}

func TestNewRemoveCommandWithMoreArguments(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	removeCmd := sollozzoctl.NewRemoveCommand(store)
	err := removeCmd.RunE(removeCmd, []string{"abc", "def"})

	assert.Error(t, err)
	assert.Equal(t, "\"sollozzo remove\" accepts only project name argument.", err.Error())
}

func TestNewRemoveCommandRemoveNotExistProject(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	removeCmd := sollozzoctl.NewRemoveCommand(store)
	err := removeCmd.RunE(removeCmd, []string{"abc"})

	assert.Error(t, err)
	assert.Equal(t, "Project not available", err.Error())
}

// release tests

func TestNewReleaseCommandMissingFlag(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)

	err := addCmd.RunE(addCmd, []string{"abc", "1.2.3"})
	assert.NoError(t, err)

	releaseCmd := sollozzoctl.NewReleaseCommand(store)
	err = releaseCmd.RunE(releaseCmd, []string{"abc"})
	assert.NoError(t, err)
}

func TestNewReleaseCommandMissingProjectName(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)

	err := addCmd.RunE(addCmd, []string{"abc", "1.2.3"})
	assert.NoError(t, err)

	releaseCmd := sollozzoctl.NewReleaseCommand(store)
	err = releaseCmd.RunE(releaseCmd, []string{})

	assert.Error(t, err)
	assert.Equal(t, "\"sollozzo release\" accepts project name and version parameter(optional) arguments.", err.Error())
}

func TestNewReleaseCommandWithMoreArguments(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)

	err := addCmd.RunE(addCmd, []string{"abc", "1.2.3"})
	assert.NoError(t, err)

	releaseCmd := sollozzoctl.NewReleaseCommand(store)
	err = releaseCmd.RunE(releaseCmd, []string{"abc", "klm"})

	assert.Error(t, err)
	assert.Equal(t, "\"sollozzo release\" accepts project name and version parameter(optional) arguments.", err.Error())
}

func TestNewReleaseCommandWithMajorFlag(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)

	err := addCmd.RunE(addCmd, []string{"abc", "1.2.3"})
	assert.NoError(t, err)

	releaseCmd := sollozzoctl.NewReleaseCommand(store)
	flag := releaseCmd.Flag("major")
	flag.Value.Set(strconv.FormatBool(true));

	err = releaseCmd.RunE(releaseCmd, []string{"abc"})
	assert.NoError(t, err)
}

func TestNewReleaseCommandWithAllFlag(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	addCmd := sollozzoctl.NewAddCommand(store)

	err := addCmd.RunE(addCmd, []string{"abc", "1.2.3"})
	assert.NoError(t, err)

	releaseCmd := sollozzoctl.NewReleaseCommand(store)

	major := releaseCmd.Flag("major")
	major.Value.Set(strconv.FormatBool(true));

	minor := releaseCmd.Flag("minor")
	minor.Value.Set(strconv.FormatBool(true));

	build := releaseCmd.Flag("build")
	build.Value.Set(strconv.FormatBool(true));

	err = releaseCmd.RunE(releaseCmd, []string{"abc"})
	assert.NoError(t, err)
}

func TestNewReleaseCommandNotExistProject(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	releaseCmd := sollozzoctl.NewReleaseCommand(store)
	releaseCmd.HasFlags()
	err := releaseCmd.RunE(releaseCmd, []string{"abc"})

	assert.Error(t, err)
	assert.Equal(t, "Project can not available", err.Error())
}