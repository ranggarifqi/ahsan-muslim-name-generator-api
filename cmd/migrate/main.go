package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper/database/postgresql"

	_ "github.com/lib/pq"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Must include goose command")
		return
	}

	dir := "./cmd/migrate"

	dbString, err := postgresql.GetDBString()
	helper.HandleError("GetDBString Error", err)

	db, err := goose.OpenDBWithDriver("postgres", dbString)
	helper.HandleError("goose: failed to open DB", err)

	defer func() {
		err := db.Close()
		helper.HandleError("goose: failed to close DB", err)
	}()

	command := os.Args[1]
	arguments := []string{}
	if len(os.Args) > 2 {
		arguments = append(arguments, os.Args[2:]...)
	}
	err = goose.Run(command, db, dir, arguments...)
	helper.HandleError(fmt.Sprintf("goose %v", command), err)
}
