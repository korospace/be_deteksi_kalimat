package controllers

import (
	"be_deteksi_kalimat/database"
	"be_deteksi_kalimat/helpers"
	"be_deteksi_kalimat/models"
	"encoding/json"
	"net/http"
)

func SingleTraining(w http.ResponseWriter, r *http.Request) {
	// get request
	var request models.SingleTrainingReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// text will predict
	result, err := helpers.PredictText(database.DB, request.RawText)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Single Training", result)
}

func BulkTraining(w http.ResponseWriter, r *http.Request) {
	// Menerima file upload
	file, _, err := r.FormFile("file_bulktraining")
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

	// Bulk Training
	var isFirstRow = true
	var predictions []models.Prediction

	for _, row := range sheet.Rows {
		if isFirstRow {
			isFirstRow = false
			continue // Skip the first row
		}
		if len(row.Cells) < 2 {
			continue // Skip rows that don't have both columns populated
		}

		rawText := row.Cells[0].String()
		expectedCategory := row.Cells[1].String()

		resultPredict, err := helpers.PredictText(database.DB, rawText)
		if err != nil {
			helpers.Response(w, 500, err.Error(), nil)
			return
		}

		predictions = append(predictions, models.Prediction{
			RawText:           rawText,
			ExpectedCategory:  expectedCategory,
			PredictedCategory: resultPredict.BestCategory,
			Detail:            resultPredict,
		})
	}

	// Hitung confusion matrix
	confusionMatrix := helpers.CalculateConfusionMatrix(predictions)

	// Hitung akurasi
	accuracy := helpers.CalculateAccuracy(confusionMatrix)

	// Result
	bulkTrainingRes := models.BulkTrainingRes{
		ConfusionMatrix: confusionMatrix,
		Accuracy:        accuracy,
		Predictions:     predictions,
	}

	helpers.Response(w, 200, "Bulk Training", bulkTrainingRes)
}
