package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	requiredEnvVars := []string{"HOST_DB", "USER_DB", "PASSWORD_DB", "NAME_DB", "PORT_DB"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return fmt.Errorf("env var %s not set", envVar)
		}
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST_DB"), os.Getenv("USER_DB"), os.Getenv("PASSWORD_DB"),
		os.Getenv("NAME_DB"), os.Getenv("PORT_DB"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}
