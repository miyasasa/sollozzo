package main

import (
	"github.com/yasinKIZILKAYA/sollozzo/sollozzoctl"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
	"os"
)

func main() {
	store := boltdb.NewStore()

	err := store.Open()

	if err != nil {
		os.Exit(1)
	}

	defer store.Close()

	cli := sollozzoctl.NewSollozzoCli(store)

	cli.Execute()
}
