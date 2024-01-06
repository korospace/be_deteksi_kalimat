package models

import "time"

type UserAccess struct {
	ID           int        `gorm:"primaryKey"`
	Name         string     `gorm:"type:varchar(255);"`
	Description  string     `gorm:"type:text;"`
	UseridCreate int        `gorm:"default:null;comment:'kolom untuk user id yang membuat data'"`
	DateCreate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal pembuatan'"`
	UseridUpdate int        `gorm:"default:null;comment:'kolom untuk user id yang mengupdate data'"`
	DateUpdate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal update'"`
	NA           string     `gorm:"type:enum('Y', 'N');default:'N';comment:'kolom untuk soft delete'"`
}

type UserAccessCreateReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserAccessUpdateReq struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
