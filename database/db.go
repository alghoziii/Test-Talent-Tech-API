package database

import (
	"log"
	"os"
	"e-ticketing/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Create connection string
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + 
	       " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLMode

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	// Auto migrate tables
	err = db.AutoMigrate(&models.User{}, &models.Terminal{})
	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}

	// Create default admin user
	createDefaultAdmin(db)

	DB = db
	log.Println("✅ Database terhubung sukses")
	log.Println("✅ User admin default dibuat: username='admin', password='admin123'")
}

func createDefaultAdmin(db *gorm.DB) {
	// Check if admin already exists
	var existingUser models.User
	result := db.Where("username = ?", "admin").First(&existingUser)

	// Jika admin belum ada, buat baru
	if result.Error != nil {
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Gagal hash password:", err)
		}

		admin := models.User{
			Username:     "admin",
			PasswordHash: string(hashedPassword),
		}

		// Create admin user
		if err := db.Create(&admin).Error; err != nil {
			log.Fatal("Gagal membuat user admin:", err)
		}

		log.Printf("✅ User admin dibuat dengan ID: %d", admin.ID)
	} else {
		log.Printf("✅ User admin sudah ada dengan ID: %d", existingUser.ID)
	}
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan instance database:", err)
	}
	sqlDB.Close()
	log.Println("✅ Koneksi database ditutup")
}