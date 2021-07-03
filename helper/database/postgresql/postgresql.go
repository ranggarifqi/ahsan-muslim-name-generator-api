package postgresql

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn, err := GetDBString()
	helper.HandleError("GetDBString Error", err)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.HandleError("DB Error", err)

	return db
}

func GetDBString() (string, error) {
	godotenv.Load()

	if os.Getenv("DATABASE_URL") != "" {
		return os.Getenv("DATABASE_URL"), nil
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"), os.Getenv("PG_PORT"), os.Getenv("PG_SSLMODE"))
	return dsn, nil
}
