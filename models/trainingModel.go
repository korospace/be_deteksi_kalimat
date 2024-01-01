package models

type SingleTrainingReq struct {
	RawText string
}

type SingleTrainingRes struct {
	BestCategory string
	BestScore    float64
	Bobot        map[string]map[string]int
}

type BulkTrainingReq struct {
	RawText     string
	Expectation string
}

type BulkTrainingRes struct {
	ConfusionMatrix map[string]int
	Accuracy        float64
	Predictions     []Prediction
}

type Dictionary struct {
	Tokenization string
	CategoryName string
}

type Prediction struct {
	RawText           string
	ExpectedCategory  string
	PredictedCategory string
	Detail            SingleTrainingRes
}
