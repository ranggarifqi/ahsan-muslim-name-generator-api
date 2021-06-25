package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
)

func init() {
	goose.AddMigration(upSeedbaseusers, downSeedbaseusers)
}

var id string = "9e6ed714-82ec-4fa7-b016-fbe32d7f4915"

func upSeedbaseusers(tx *sql.Tx) error {
	godotenv.Load()

	email := os.Getenv("SUPERADMIN_EMAIL")
	password := os.Getenv("SUPERADMIN_PASSWORD")
	fullName := os.Getenv("SUPERADMIN_FULLNAME")

	password, err := helper.HashPassword(password)
	helper.HandleError("Error hashing password:", err)

	query := `INSERT INTO "users"(id, email, password, "fullName")
		VALUES
			('%v', '%v', '%v', '%v')
	`
	query = fmt.Sprintf(query, id, email, password, fullName)

	_, err = tx.Exec(query)
	helper.HandleError("Error seeding base user:", err)

	// This code is executed when the migration is applied.
	return nil
}

func downSeedbaseusers(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	query := fmt.Sprintf(`DELETE FROM "users" WHERE id=%v`, id)
	_, err := tx.Exec(query)
	helper.HandleError("Error deleting base user:", err)
	return nil
}
