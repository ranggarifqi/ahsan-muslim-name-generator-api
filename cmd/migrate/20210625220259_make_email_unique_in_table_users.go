package main

import (
	"database/sql"

	"github.com/pressly/goose"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
)

func init() {
	goose.AddMigration(upMakeEmailUniqueInTableUsers, downMakeEmailUniqueInTableUsers)
}

func upMakeEmailUniqueInTableUsers(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`ALTER TABLE "users" ADD CONSTRAINT uniq_email UNIQUE ("email")`)
	helper.HandleError("Error add unique constraint:", err)
	return nil
}

func downMakeEmailUniqueInTableUsers(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`ALTER TABLE "users" DROP CONSTRAINT uniq_email`)
	helper.HandleError("Error dropping unique constraing:", err)
	return nil
}
