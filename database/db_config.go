package database

import (
	"be_deteksi_kalimat/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/db_deteksi_kalimat?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	log.Println("database connected!")

	// ROLLBACK
	// ========
	// db.Migrator().DropTable(&models.UserAccess{})
	// db.Migrator().DropTable(&models.User{})
	// db.Migrator().DropTable(&models.Category{})
	// db.Migrator().DropTable(&models.Dataset{})

	// MIGRATE
	// ========
	db.AutoMigrate(&models.UserAccess{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Dataset{})

	// SEEDER
	// ========
	// RunSeeder(db)
}
