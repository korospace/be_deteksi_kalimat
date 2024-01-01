package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseWithData struct {
	Status  string `json:"status"`
	Massage string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Massage string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "success"

	if code >= 400 {
		status = "failed"
	}

	if payload != nil {
		response = &ResponseWithData{
			Status:  status,
			Massage: message,
			Data:    payload,
		}
	} else {
		response = &ResponseWithoutData{
			Status:  status,
			Massage: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
