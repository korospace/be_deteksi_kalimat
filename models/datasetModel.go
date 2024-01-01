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
	Clean        string
	Stopword     string
	Stemming     string
	Tokenization string
}

type DatasetListRes struct {
	ID           int
	Raw          string
	Clean        string
	Stopword     string
	Stemming     string
	Tokenization string
	CategoryID   int
	CategoryName string
	Verify       string
	UseridVerify int
}

type DatasetCreateReq struct {
	RawText  string
	Category string
}

type DatasetUpdateReq struct {
	ID           int
	Raw          string
	Clean        string
	Stopword     string
	Stemming     string
	Tokenization string
	CategoryID   int
}

type VerifyDatasetReq struct {
	ID     int
	Verify string
}
