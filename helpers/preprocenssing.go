package helpers

import (
	"be_deteksi_kalimat/models"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/RadhiFadlillah/go-sastrawi"
	"github.com/tealeg/xlsx"
	"gorm.io/gorm"
)

func ReadExcelFile(file io.Reader) (*xlsx.Sheet, error) {
	tempFile, err := os.CreateTemp("", "uploaded-*.xlsx")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}

	xlFile, err := xlsx.OpenFile(tempFile.Name())
	if err != nil {
		return nil, err
	}

	if len(xlFile.Sheets) == 0 {
		return nil, fmt.Errorf("header harus di baris pertama")
	}

	sheet := xlFile.Sheets[0]

	return sheet, nil
}

func CleanText(text string) (string, error) {
	// Menghapus karakter aneh dan tanda baca
	re := regexp.MustCompile(`[^\w\s]`)
	cleanedText := re.ReplaceAllString(text, "")

	// Menghapus whitespace berlebihan dan mengonversi ke huruf kecil
	cleanedText = strings.ToLower(strings.TrimSpace(cleanedText))

	// Menyingkirkan karakter non-alfanumerik
	var builder strings.Builder
	for _, r := range cleanedText {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			builder.WriteRune(r)
		}
	}

	if builder.Len() == 0 {
		return "", fmt.Errorf(text)
	}

	return builder.String(), nil
}

func StopwordText(text string) (string, error) {
	stopwords := sastrawi.DefaultStopword()
	dictionary := sastrawi.DefaultDictionary()
	stemmer := sastrawi.NewStemmer(dictionary)
	sentence := text

	var cleanedWords []string
	for _, word := range sastrawi.Tokenize(sentence) {
		if stopwords.Contains(word) {
			continue
		}
		cleanedWords = append(cleanedWords, stemmer.Stem(word))
	}

	if len(cleanedWords) == 0 {
		return "", fmt.Errorf("tidak ada kata setelah penghapusan stopword")
	}

	cleanedText := strings.Join(cleanedWords, " ")
	return cleanedText, nil
}

func StemmingText(text string) (string, error) {
	dictionary := sastrawi.DefaultDictionary()
	stemmer := sastrawi.NewStemmer(dictionary)
	sentence := text

	var stemmedWords []string
	for _, word := range sastrawi.Tokenize(sentence) {
		stemmedWords = append(stemmedWords, stemmer.Stem(word))
	}

	if len(stemmedWords) == 0 {
		return "", fmt.Errorf("tidak ada kata setelah stemming")
	}

	stemmedText := strings.Join(stemmedWords, " ")
	return stemmedText, nil
}

func TokenizeText(text string) (string, error) {
	tokens := strings.Fields(text)
	if len(tokens) == 0 {
		return "", fmt.Errorf("tidak ada token yang dihasilkan")
	}
	tokenized := strings.Join(tokens, ", ")
	return tokenized, nil
}

func PreProcess(rawText string) (models.PreProcessedText, error) {
	var processedText models.PreProcessedText

	// cleaning raw text
	cleanText, err := CleanText(rawText)
	if err != nil {
		return processedText, err
	}
	processedText.Clean = cleanText

	// stopword text
	stopwordText, err := StopwordText(cleanText)
	if err != nil {
		return processedText, err
	}
	processedText.Stopword = stopwordText

	// stemming text
	stemmingText, err := StemmingText(stopwordText)
	if err != nil {
		return processedText, err
	}
	processedText.Stemming = stemmingText

	// tokenize text
	tokenizeText, err := TokenizeText(stemmingText)
	if err != nil {
		return processedText, err
	}
	processedText.Tokenization = tokenizeText

	return processedText, nil
}

func PredictText(DB *gorm.DB, rawText string) (models.SingleTrainingRes, error) {
	// Preprocess the text
	preProcessedText, err := PreProcess(rawText)
	if err != nil {
		return models.SingleTrainingRes{}, err
	}

	// Fetch datasets from the database
	var datasets []models.Dataset
	if err := DB.Where("na = ?", "N").Order("id desc").Find(&datasets).Error; err != nil {
		return models.SingleTrainingRes{}, err
	}

	// Process datasets to build the dictionary
	dictionary := make([]models.Dictionary, len(datasets))
	for i, r := range datasets {
		dictionary[i].Tokenization = r.Tokenization

		var category models.Category
		if err := DB.Where("id = ?", r.CategoryID).First(&category).Error; err == nil {
			dictionary[i].CategoryName = category.Name
		}
	}

	// Calculate category and word counts
	categoryCounts := make(map[string]int)
	wordCategoryCounts := make(map[string]map[string]int)
	for _, item := range dictionary {
		category := item.CategoryName
		if _, exists := categoryCounts[category]; !exists {
			categoryCounts[category] = 0
		}
		categoryCounts[category]++

		if _, exists := wordCategoryCounts[category]; !exists {
			wordCategoryCounts[category] = make(map[string]int)
		}

		words := strings.Split(item.Tokenization, ", ")
		for _, word := range words {
			if _, exists := wordCategoryCounts[category][word]; !exists {
				wordCategoryCounts[category][word] = 0
			}
			wordCategoryCounts[category][word]++
		}
	}

	// Calculate word weights for the input text
	arrSplitedText := make(map[string]map[string]int)
	splitedText := strings.Split(preProcessedText.Tokenization, ", ")

	for _, word := range splitedText {
		arrSplitedText[word] = make(map[string]int)
	}

	for indexCategory := range categoryCounts {
		for _, word := range splitedText {
			if _, exists := arrSplitedText[word][indexCategory]; !exists {
				arrSplitedText[word][indexCategory] = 0
			}
			arrSplitedText[word][indexCategory] += wordCategoryCounts[indexCategory][word]
		}
	}

	// Determine the best scoring category
	bestCategory := ""
	bestScore := -math.MaxFloat64

	for indexCategory, numOfCategory := range categoryCounts {
		score := math.Log(float64(numOfCategory) / float64(len(dictionary)))

		for _, word := range splitedText {
			count := wordCategoryCounts[indexCategory][word]

			score = score + math.Log((float64(count)+1)/(float64(len(wordCategoryCounts[indexCategory]))+float64(len(splitedText))))
		}

		if score > bestScore {
			bestCategory = indexCategory
			bestScore = score
		}
	}

	// Create the result
	result := models.SingleTrainingRes{
		BestCategory: bestCategory,
		BestScore:    bestScore,
		Bobot:        arrSplitedText,
	}

	return result, nil
}

func CalculateConfusionMatrix(predictions []models.Prediction) map[string]int {
	confusionMatrix := map[string]int{
		"TP": 0,
		"TN": 0,
		"FP": 0,
		"FN": 0,
	}

	for _, pred := range predictions {
		if pred.ExpectedCategory == "bullying" && pred.PredictedCategory == "bullying" {
			confusionMatrix["TP"]++
		} else if pred.ExpectedCategory == "netral" && pred.PredictedCategory == "netral" {
			confusionMatrix["TN"]++
		} else if pred.ExpectedCategory == "netral" && pred.PredictedCategory == "bullying" {
			confusionMatrix["FP"]++
		} else if pred.ExpectedCategory == "bullying" && pred.PredictedCategory == "netral" {
			confusionMatrix["FN"]++
		}
	}

	return confusionMatrix
}

func CalculateAccuracy(confusionMatrix map[string]int) float64 {
	TP := float64(confusionMatrix["TP"])
	TN := float64(confusionMatrix["TN"])
	FP := float64(confusionMatrix["FP"])
	FN := float64(confusionMatrix["FN"])

	accuracy := (TP + TN) / (TP + TN + FP + FN)
	return accuracy
}
