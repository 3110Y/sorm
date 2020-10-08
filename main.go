package main

import (
	"fmt"
	"os"
	"sorm/src/config"
	"sorm/src/creator"
)

var err error = nil

func main() {
	config.GetInstance()
	switch os.Args[1] {
	case "start":
	case "create":
		switch os.Args[2] {
		case "entity":
			//err = creator.CreateEntity()
		case "migration":
			//err = creator.CreateMigration()
		case "structure":
			err = creator.CreateStucture()
		}
	case "migrate":
	default:
		fmt.Println("Please use start, make or migration")
	}
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR: %v", err))
		os.Exit(1)
	}
}
