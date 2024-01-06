package models

import "time"

type Dataset struct {
	ID           int        `gorm:"primaryKey"`
	Raw          string     `gorm:"type:text;default:null;comment:'kolom untuk menyimpan kalimat raw/asli'"`
	Clean        string     `gorm:"type:text;default:null;comment:'kolom untuk menyimpan kalimat hasil dari proses pembersihan'"`
	Stopword     string     `gorm:"type:text;default:null;comment:'kolom untuk menyimpan kalimat hasil dari proses stopword'"`
	Stemming     string     `gorm:"type:text;default:null;comment:'kolom untuk menyimpan kalimat hasil dari proses stemming'"`
	Tokenization string     `gorm:"type:text;default:null;comment:'kolom untuk menyimpan kalimat hasil dari proses tokenisasi'"`
	CategoryID   int        `gorm:"default:null;comment:'kolom untuk menyimpan id pada tabel category'"`
	Verify       string     `gorm:"type:enum('Y', 'N');default:'N';comment:'kolom untuk status verifikasi oleh pakar'"`
	UseridVerify int        `gorm:"default:null;comment:'kolom untuk user id yang melakukan verifikasi'"`
	UseridCreate int        `gorm:"default:null;comment:'kolom untuk user id yang membuat data'"`
	DateCreate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal pembuatan'"`
	UseridUpdate int        `gorm:"default:null;comment:'kolom untuk user id yang mengupdate data'"`
	DateUpdate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal update'"`
	NA           string     `gorm:"type:enum('Y', 'N');default:'N';comment:'kolom untuk soft delete'"`
}

type PreProcessedText struct {
	Clean        string `json:"clean"`
	Stopword     string `json:"stopword"`
	Stemming     string `json:"stemming"`
	Tokenization string `json:"tokenization"`
}

type DatasetListRes struct {
	ID           int    `json:"id"`
	Raw          string `json:"raw"`
	Clean        string `json:"clean"`
	Stopword     string `json:"stopword"`
	Stemming     string `json:"stemming"`
	Tokenization string `json:"tokenization"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Verify       string `json:"verify"`
	UseridVerify int    `json:"userid_verify"`
}

type DatasetCreateReq struct {
	RawText  string `json:"raw_text"`
	Category string `json:"category"`
}

type DatasetUpdateReq struct {
	ID           int    `json:"id"`
	Raw          string `json:"raw"`
	Clean        string `json:"clean"`
	Stopword     string `json:"stopword"`
	Stemming     string `json:"stemming"`
	Tokenization string `json:"tokenization"`
	CategoryID   int    `json:"category_id"`
}

type VerifyDatasetReq struct {
	ID     int    `json:"id"`
	Verify string `json:"verify"`
}
