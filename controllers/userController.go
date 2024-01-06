package controllers

import (
	"be_deteksi_kalimat/database"
	"be_deteksi_kalimat/helpers"
	"be_deteksi_kalimat/models"
	"encoding/json"
	"net/http"
	"time"
)

func ListingUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := database.DB.Where("na = ?", "N").Order("id desc").Find(&users).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, "Failed to fetch user", nil)
		return
	}

	helpers.Response(w, 200, "User List", users)
}

func Me(w http.ResponseWriter, r *http.Request) {
	// get token info
	tokeninfo := r.Context().Value("tokeninfo").(*helpers.TokenInfo)

	var user models.User
	if err := database.DB.First(&user, "id = ?", tokeninfo.ID).Error; err != nil {
		helpers.Response(w, 500, "failed", err)
		return
	}

	me := models.Me{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		UserAccessID: user.UserAccessID,
	}

	helpers.Response(w, 200, "Me", me)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.UserCreateReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// make create
	currentTime := time.Now()
	user := models.User{
		Name:         request.Name,
		Email:        request.Email,
		Password:     helpers.HashPassword(request.Password),
		UserAccessID: request.UserAccessID,
		UseridCreate: userinfo.ID,
		DateCreate:   &currentTime,
	}

	// create row
	if err := database.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "User Created", user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.UserUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// cek user id
	var user models.User
	if err := database.DB.Where("id = ? AND NA = ?", request.ID, "N").First(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make update
	currentTime := time.Now()
	user.Name = request.Name
	user.Email = request.Email
	user.Password = helpers.HashPassword(request.Password)
	user.UserAccessID = request.UserAccessID
	user.UseridUpdate = userinfo.ID
	user.DateUpdate = &currentTime

	// update row
	if err := database.DB.Save(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "User Updated", user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	userID := r.URL.Query().Get("id")

	// cek user access id
	var user models.User
	if err := database.DB.Where("id = ? AND NA = ?", userID, "N").First(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make delete
	currentTime := time.Now()
	user.NA = "Y"
	user.UseridUpdate = userinfo.ID
	user.DateUpdate = &currentTime

	// delete row
	if err := database.DB.Save(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "User Deleted", nil)
}