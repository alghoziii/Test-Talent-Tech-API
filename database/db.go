package database

import (
	"e-ticketing/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    // Gunakan environment variable atau default value
    dbHost := getEnv("DB_HOST", "localhost")
    dbUser := getEnv("DB_USER", "postgres")
    dbPassword := getEnv("DB_PASSWORD", "")
    dbName := getEnv("DB_NAME", "e_ticketing")
    dbPort := getEnv("DB_PORT", "5432")
    dbSSLMode := getEnv("DB_SSLMODE", "disable")

    dsn := "host=" + dbHost + 
           " user=" + dbUser + 
           " password=" + dbPassword + 
           " dbname=" + dbName + 
           " port=" + dbPort + 
           " sslmode=" + dbSSLMode

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto migrate tables
    err = db.AutoMigrate(&models.User{}, &models.Terminal{})
    if err != nil {
        log.Fatal("Failed to auto-migrate tables:", err)
    }

    DB = db
    log.Println("Database connected successfully")
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func GetDB() *gorm.DB {
    return DB
}

func CloseDB() {
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatal("Failed to get database instance:", err)
    }
    sqlDB.Close()
}