package main

import (
	"database/sql"

	"github.com/pressly/goose"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
)

func init() {
	goose.AddMigration(upCreatetablenamemappings, downCreatetablenamemappings)
}

func upCreatetablenamemappings(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	query := `CREATE TABLE IF NOT EXISTS "nameMappings" (
		nameId uuid NOT NULL,
		nameTypeId uuid NOT NULL,
		PRIMARY KEY(nameId, nameTypeId),
		CONSTRAINT fk_name_id FOREIGN KEY(nameId) REFERENCES "names"(id)
			ON DELETE CASCADE,
		CONSTRAINT fk_name_type_id FOREIGN KEY(nameTypeId) REFERENCES "nameTypes"(id)
			ON DELETE CASCADE
	)`
	_, err := tx.Exec(query)
	helper.HandleError("Error creating nameMappings table", err)
	return nil
}

func downCreatetablenamemappings(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS "nameMappings"`)
	helper.HandleError("Error when dropping table nameMappings", err)
	return nil
}
