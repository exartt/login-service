package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	requiredEnvVars := []string{"HOST_DB", "USER_DB", "PASSWORD_DB", "DATABASE_DB", "PORT_DB"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return nil, fmt.Errorf("env var %s not set", envVar)
		}
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST_DB"), os.Getenv("USER_DB"), os.Getenv("PASSWORD_DB"),
		os.Getenv("DATABASE_DB"), os.Getenv("PORT_DB"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}
