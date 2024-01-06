package controllers

import (
	"be_deteksi_kalimat/database"
	"be_deteksi_kalimat/helpers"
	"be_deteksi_kalimat/models"
	"encoding/json"
	"net/http"
	"time"
)

func ListingAccess(w http.ResponseWriter, r *http.Request) {
	var user_accesses []models.UserAccess

	if err := database.DB.Where("na = ?", "N").Order("id desc").Find(&user_accesses).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, "Failed to fetch user accesses", nil)
		return
	}

	helpers.Response(w, 200, "User List", user_accesses)
}

func CreateAccess(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.UserAccessCreateReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// make create
	currentTime := time.Now()
	user_access := models.UserAccess{
		Name:         request.Name,
		Description:  request.Description,
		UseridCreate: userinfo.ID,
		DateCreate:   &currentTime,
	}

	// create row
	if err := database.DB.Create(&user_access).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "User Access Created", user_access)
}

func UpdateAccess(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.UserAccessUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// cek user access id
	var user_access models.UserAccess
	if err := database.DB.Where("id = ? AND NA = ?", request.ID, "N").First(&user_access).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make update
	currentTime := time.Now()
	user_access.Name = request.Name
	user_access.Description = request.Description
	user_access.UseridUpdate = userinfo.ID
	user_access.DateUpdate = &currentTime

	// update row
	if err := database.DB.Save(&user_access).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "User Access Updated", user_access)
}

func DeleteAccess(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	userAccessID := r.URL.Query().Get("id")

	// cek user access id
	var user_access models.UserAccess
	if err := database.DB.Where("id = ? AND NA = ?", userAccessID, "N").First(&user_access).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make delete
	currentTime := time.Now()
	user_access.NA = "Y"
	user_access.UseridUpdate = userinfo.ID
	user_access.DateUpdate = &currentTime

	// delete row
	if err := database.DB.Save(&user_access).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "User Access Deleted", nil)
}
