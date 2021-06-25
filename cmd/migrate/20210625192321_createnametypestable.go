package main

import (
	"database/sql"

	"github.com/pressly/goose"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
)

func init() {
	goose.AddMigration(upCreatenametypestable, downCreatenametypestable)
}

func upCreatenametypestable(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	query := `CREATE TABLE IF NOT EXISTS "nameTypes" (
		id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		"createdBy" uuid NULL,
		"updatedBy" uuid NULL,
		"createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		"updatedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		"deletedAt" TIMESTAMP NULL
	)`

	_, err := tx.Exec(query)
	helper.HandleError("Error creating nameTypes table", err)
	return nil
}

func downCreatenametypestable(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS "nameTypes"`)
	helper.HandleError("Error when dropping table nameTypes", err)
	return nil
}
