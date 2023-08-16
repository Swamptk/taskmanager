package main

import (
	"Task/cmd"
	"Task/db"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(db.ClearDone())
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
