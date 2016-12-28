package sollozzoctl_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
	"github.com/yasinKIZILKAYA/sollozzo/sollozzoctl"
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
