package controllers

import (
	"be_deteksi_kalimat/database"
	"be_deteksi_kalimat/helpers"
	"be_deteksi_kalimat/models"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var LoginReq models.LoginReq
	var user models.User

	// get request
	if err := json.NewDecoder(r.Body).Decode(&LoginReq); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// cek email
	if err := database.DB.First(&user, "email = ?", LoginReq.Email).Error; err != nil {
		helpers.Response(w, 401, "Wrong email", nil)
		return
	}

	// cek password
	if err := helpers.VerifyPassword(user.Password, LoginReq.Password); err != nil {
		helpers.Response(w, 401, "Wrong password", nil)
		return
	}

	// create token
	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	LoginRes := models.LoginRes{
		Token: token,
	}
	helpers.Response(w, 200, "Login successfully", LoginRes)
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
