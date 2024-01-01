package models

import "time"

type Category struct {
	ID           int        `gorm:"primaryKey"`
	Name         string     `gorm:"type:varchar(255);"`
	Description  string     `gorm:"type:text;"`
	UseridCreate int        `gorm:"default:null;comment:'kolom untuk user id yang membuat data'"`
	DateCreate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal pembuatan'"`
	UseridUpdate int        `gorm:"default:null;comment:'kolom untuk user id yang mengupdate data'"`
	DateUpdate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal update'"`
	NA           string     `gorm:"type:enum('Y', 'N');default:'N';comment:'kolom untuk soft delete'"`
}

type CategoryCreateReq struct {
	Name        string
	Description string
}

type CategoryUpdateReq struct {
	ID          int
	Name        string
	Description string
}
