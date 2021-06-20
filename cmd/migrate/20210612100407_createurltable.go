package main

import (
	"database/sql"

	"github.com/pressly/goose"
	"github.com/ranggarifqi/qashir-api/helper"
)

func init() {
	goose.AddMigration(upCreateurltable, downCreateurltable)
}

func upCreateurltable(tx *sql.Tx) error {
	// This code is executed when the migration is applied.

	query := `CREATE TABLE IF NOT EXISTS urls (
		id VARCHAR(6) PRIMARY KEY,
		url VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL
)`

	_, err := tx.Exec(query)
	helper.HandleError("Migration error", err)

	return nil
}

func downCreateurltable(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	query := `DROP TABLE IF EXISTS urls`

	_, err := tx.Exec(query)
	helper.HandleError("Rollback migration error", err)
	return nil
}
