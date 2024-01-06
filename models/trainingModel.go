package models

type SingleTrainingReq struct {
	RawText string `json:"raw_text"`
}

type SingleTrainingRes struct {
	BestCategory string                    `json:"best_category"`
	BestScore    float64                   `json:"best_score"`
	Bobot        map[string]map[string]int `json:"bobot"`
}

type BulkTrainingReq struct {
	RawText     string `json:"raw_text"`
	Expectation string `json:"expectation"`
}

type BulkTrainingRes struct {
	ConfusionMatrix map[string]int `json:"confusion_matrix"`
	Accuracy        float64        `json:"accuracy"`
	Predictions     []Prediction   `json:"predictions"`
}

type Dictionary struct {
	Tokenization string `json:"tokenization"`
	CategoryName string `json:"category_name"`
}

type Prediction struct {
	RawText           string            `json:"raw_text"`
	ExpectedCategory  string            `json:"expected_category"`
	PredictedCategory string            `json:"predicted_category"`
	Detail            SingleTrainingRes `json:"detail"`
}
