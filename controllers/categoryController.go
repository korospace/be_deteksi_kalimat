package controllers

import (
	"be_deteksi_kalimat/database"
	"be_deteksi_kalimat/helpers"
	"be_deteksi_kalimat/models"
	"encoding/json"
	"net/http"
	"time"
)

func ListingCategory(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category

	if err := database.DB.Where("na = ?", "N").Find(&categories).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, "Failed to fetch categories", nil)
		return
	}

	helpers.Response(w, 200, "Category List", categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.CategoryCreateReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// make create
	currentTime := time.Now()
	category := models.Category{
		Name:         request.Name,
		Description:  request.Description,
		UseridCreate: userinfo.ID,
		DateCreate:   &currentTime,
	}

	// create row
	if err := database.DB.Create(&category).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Category Created", category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.CategoryUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// cek category id
	var category models.Category
	if err := database.DB.Where("id = ? AND NA = ?", request.ID, "N").First(&category).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make update
	currentTime := time.Now()
	category.Name = request.Name
	category.Description = request.Description
	category.UseridUpdate = userinfo.ID
	category.DateUpdate = &currentTime

	// update row
	if err := database.DB.Save(&category).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Category Updated", category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	categoryID := r.URL.Query().Get("id")

	// cek category id
	var category models.Category
	if err := database.DB.Where("id = ? AND NA = ?", categoryID, "N").First(&category).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make delete
	currentTime := time.Now()
	category.NA = "Y"
	category.UseridUpdate = userinfo.ID
	category.DateUpdate = &currentTime

	// delete row
	if err := database.DB.Save(&category).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Category Deleted", nil)
}
