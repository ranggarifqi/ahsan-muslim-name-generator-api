package main

import (
	"database/sql"

	"github.com/pressly/goose"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
)

func init() {
	goose.AddMigration(upCreatetableusers, downCreatetableusers)
}

func upCreatetableusers(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	helper.HandleError("Error create uuid-ossp extension", err)

	query := `CREATE TABLE IF NOT EXISTS users (
		id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		fullName VARCHAR(255) NOT NULL,
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deletedAt TIMESTAMP NULL
)`
	_, err = tx.Exec(query)
	helper.HandleError("Error create table users", err)

	return nil
}

func downCreatetableusers(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("DROP TABLE IF EXISTS users")
	helper.HandleError("Error on dropping table users", err)

	_, err = tx.Exec("DROP EXTENSION IF EXISTS \"uuid-ossp\"")
	helper.HandleError("Error on dropping uuid-ossp extension", err)

	return nil
}
