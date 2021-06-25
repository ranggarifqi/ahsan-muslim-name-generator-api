package main

import (
	"database/sql"

	"github.com/pressly/goose"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
)

func init() {
	goose.AddMigration(upCreatetablenames, downCreatetablenames)
}

func upCreatetablenames(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	query := `CREATE TABLE IF NOT EXISTS "names" (
		id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		gender VARCHAR(1) NOT NULL DEFAULT 'm',
		isVerified BOOLEAN NOT NULL DEFAULT FALSE,
		createdBy uuid NULL,
		updatedBy uuid NULL,
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deletedAt TIMESTAMP NULL
)`
	_, err := tx.Exec(query)
	helper.HandleError("Error when creating table names", err)

	return nil
}

func downCreatetablenames(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS "names"`)
	helper.HandleError("Error when dropping table names", err)
	return nil
}
