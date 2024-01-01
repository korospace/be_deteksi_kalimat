package database

import (
	"be_deteksi_kalimat/helpers"
	"be_deteksi_kalimat/models"
	"fmt"

	"gorm.io/gorm"
)

func SeedData(db *gorm.DB, data interface{}) error {
	if err := db.Create(data).Error; err != nil {
		return fmt.Errorf("error seeding data: %v", err)
	}
	fmt.Println("Data seeded successfully")
	return nil
}

func RunSeeder(db *gorm.DB) {
	// User Access
	usersAccess := []models.UserAccess{
		{Name: "superadmin", Description: "Super Admin"},
		{Name: "pakar", Description: "Pakar"},
	}

	if err := SeedData(db, &usersAccess); err != nil {
		panic(fmt.Sprintf("failed to seed UserAccess: %v", err))
	}

	// Users
	users := []models.User{
		{Name: "Super Admin", Email: "superadmin@test.com", Password: helpers.HashPassword("superadmin"), UserAccessID: 1},
		{Name: "Pakar 1", Email: "pakar1@test.com", Password: helpers.HashPassword("pakar1"), UserAccessID: 2},
	}

	if err := SeedData(db, &users); err != nil {
		panic(fmt.Sprintf("failed to seed User: %v", err))
	}

	// Categories
	category := []models.Category{
		{Name: "netral", Description: "Kalimat Netral"},
		{Name: "bullying", Description: "Kalimat Bullying"},
	}

	if err := SeedData(db, &category); err != nil {
		panic(fmt.Sprintf("failed to seed Category: %v", err))
	}
}
