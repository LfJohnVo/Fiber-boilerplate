package databases

import (
	"fmt"
	"log"

	"fiber-boilerplate/internal/models"
	"fiber-boilerplate/utils"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Database   *gorm.DB
	ConfigData utils.Config
)

func LoadDatabaseConfig() error {
	var err error
	ConfigData, err = utils.LoadConfig(".")
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}
	return nil
}

func BuildDatabaseURI() string {
	switch ConfigData.DB_CONNECTION {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			ConfigData.DB_USER,
			ConfigData.DB_PASSWORD,
			ConfigData.DB_HOST,
			ConfigData.DB_PORT,
			ConfigData.DB_NAME,
		)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			ConfigData.DB_HOST,
			ConfigData.DB_PORT,
			ConfigData.DB_USER,
			ConfigData.DB_NAME,
			ConfigData.DB_PASSWORD,
			ConfigData.DB_SSL_MODE,
		)
	case "sqlite":
		return ConfigData.DB_NAME
	default:
		log.Fatalf("Unsupported database type: %s", ConfigData.DB_CONNECTION)
		return ""
	}
}

func Connect() error {
	var err error

	if err := LoadDatabaseConfig(); err != nil {
		return err
	}

	dbURI := BuildDatabaseURI()
	dbConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}

	switch ConfigData.DB_CONNECTION {
	case "mysql":
		Database, err = gorm.Open(mysql.Open(dbURI), dbConfig)
	case "postgres":
		Database, err = gorm.Open(postgres.Open(dbURI), dbConfig)
	case "sqlite":
		Database, err = gorm.Open(sqlite.Open(dbURI), dbConfig)
	default:
		log.Fatalf("Unsupported database type")
		return fmt.Errorf("unsupported database type")
	}

	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	Database.Logger = logger.Default.LogMode(logger.Info)

	// Migrate models
	if err := Database.AutoMigrate(&models.User{}, &models.PasswordReset{}, &models.Contact{}); err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}

	return nil
}
