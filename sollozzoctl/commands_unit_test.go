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

func TestNewAddCommand(t *testing.T) {
	store := NewTestStore()

	defer store.Close()

	cmd := sollozzoctl.NewAddCommand(store)

	err := cmd.RunE(cmd, []string{"xxx"})

	assert.NoError(t, err)
}
