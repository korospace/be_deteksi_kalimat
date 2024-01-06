package controllers

import (
	"be_deteksi_kalimat/database"
	"be_deteksi_kalimat/helpers"
	"be_deteksi_kalimat/models"
	"encoding/json"
	"net/http"
	"time"
)

func ListingDataset(w http.ResponseWriter, r *http.Request) {
	var datasets []models.Dataset
	var datasetListRes []models.DatasetListRes

	if err := database.DB.Where("na = ?", "N").Order("id desc").Find(&datasets).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, "Failed to fetch dataset", err.Error())
		return
	}

	datasetListRes = make([]models.DatasetListRes, len(datasets))

	for i, r := range datasets {
		datasetListRes[i].ID = r.ID
		datasetListRes[i].Raw = r.Raw
		datasetListRes[i].Clean = r.Clean
		datasetListRes[i].Stopword = r.Stopword
		datasetListRes[i].Stemming = r.Stemming
		datasetListRes[i].Tokenization = r.Tokenization
		datasetListRes[i].CategoryID = r.CategoryID
		datasetListRes[i].Verify = r.Verify
		datasetListRes[i].UseridVerify = r.UseridVerify
		datasetListRes[i].CategoryName = "-"

		var category models.Category
		if err := database.DB.Where("id = ?", r.CategoryID).First(&category).Error; err == nil {
			datasetListRes[i].CategoryName = category.Name
		}
	}

	helpers.Response(w, 200, "Dataset List", datasetListRes)
}

func CreateDataset(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.DatasetCreateReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// labeling
	var categoryID int
	var category models.Category
	if err := database.DB.Where("id = ? OR Name = ?", request.Category, request.Category).First(&category).Error; err == nil {
		categoryID = category.ID
	}

	// Pre Processing text
	preProcessedText, err := helpers.PreProcess(request.RawText)
	if err != nil {
		helpers.Response(w, 500, "Gagal pre process teks", nil)
		return
	}

	// make dataset
	currentTime := time.Now()
	dataset := models.Dataset{
		CategoryID:   categoryID,
		Raw:          request.RawText,
		Clean:        preProcessedText.Clean,
		Stopword:     preProcessedText.Stopword,
		Stemming:     preProcessedText.Stemming,
		Tokenization: preProcessedText.Tokenization,
		UseridCreate: userinfo.ID,
		DateCreate:   &currentTime,
	}

	// create row
	if err := database.DB.Create(&dataset).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Dataset Created", dataset)
}

func ImportDataset(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// Menerima file upload
	file, _, err := r.FormFile("file_dataset")
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	defer file.Close()

	// Membaca data excel
	sheet, err := helpers.ReadExcelFile(file)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// Masukin data excel ke object
	var newDatasets []models.Dataset
	var isFirstRow = true
	for _, row := range sheet.Rows {
		if isFirstRow {
			isFirstRow = false
			continue // Skip the first row
		}
		if len(row.Cells) < 2 {
			continue // Skip rows that don't have both columns populated
		}

		RawText := row.Cells[0].String()
		Category := row.Cells[1].String()

		// labeling
		var categoryID int
		var category models.Category
		if err := database.DB.Where("id = ? OR Name = ?", Category, Category).First(&category).Error; err == nil {
			categoryID = category.ID
		}

		// Pre Processing text
		preProcessedText, err := helpers.PreProcess(RawText)
		if err != nil {
			helpers.Response(w, 500, err.Error(), nil)
			return
		}

		// make dataset
		currentTime := time.Now()
		dataset := models.Dataset{
			CategoryID:   categoryID,
			Raw:          RawText,
			Clean:        preProcessedText.Clean,
			Stopword:     preProcessedText.Stopword,
			Stemming:     preProcessedText.Stemming,
			Tokenization: preProcessedText.Tokenization,
			UseridCreate: userinfo.ID,
			DateCreate:   &currentTime,
		}

		// create row
		if err := database.DB.Create(&dataset).Error; err != nil {
			helpers.Response(w, 500, err.Error(), nil)
			return
		}

		newDatasets = append(newDatasets, dataset)
	}

	helpers.Response(w, 200, "New Datasets", newDatasets)
}

func VerifyDataset(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.VerifyDatasetReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// cek dataset id
	var dataset models.Dataset
	if err := database.DB.Where("id = ? AND NA = ?", request.ID, "N").First(&dataset).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make update
	currentTime := time.Now()
	dataset.Verify = request.Verify
	dataset.UseridVerify = userinfo.ID
	dataset.DateUpdate = &currentTime

	// update row
	if err := database.DB.Save(&dataset).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Dataset Verified", nil)
}

func UpdateDataset(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	var request models.DatasetUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// cek dataset id
	var dataset models.Dataset
	if err := database.DB.Where("id = ? AND NA = ?", request.ID, "N").First(&dataset).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make update
	currentTime := time.Now()
	dataset.CategoryID = request.CategoryID
	dataset.Raw = request.Raw
	dataset.Clean = request.Clean
	dataset.Stopword = request.Stopword
	dataset.Stemming = request.Stemming
	dataset.Tokenization = request.Tokenization
	dataset.UseridUpdate = userinfo.ID
	dataset.DateUpdate = &currentTime

	// update row
	if err := database.DB.Save(&dataset).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Dataset Updated", dataset)
}

func DeleteDataset(w http.ResponseWriter, r *http.Request) {
	// get token info
	userinfo := r.Context().Value("userinfo").(models.User)

	// get request
	datasetID := r.URL.Query().Get("id")

	// cek category id
	var dataset models.Dataset
	if err := database.DB.Where("id = ? AND NA = ?", datasetID, "N").First(&dataset).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// make delete
	currentTime := time.Now()
	dataset.NA = "Y"
	dataset.UseridUpdate = userinfo.ID
	dataset.DateUpdate = &currentTime

	// delete row
	if err := database.DB.Save(&dataset).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Dataset Deleted", nil)
}
