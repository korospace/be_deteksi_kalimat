package models

import "time"

type UserAccess struct {
	ID           int        `gorm:"primaryKey" json:"id"`
	Name         string     `gorm:"type:varchar(255);" json:"name"`
	Description  string     `gorm:"type:text;" json:"description"`
	UseridCreate int        `gorm:"default:null;comment:'kolom untuk user id yang membuat data'" json:"userid_create"`
	DateCreate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal pembuatan'" json:"date_create"`
	UseridUpdate int        `gorm:"default:null;comment:'kolom untuk user id yang mengupdate data'" json:"userid_update"`
	DateUpdate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal update'" json:"date_update"`
	NA           string     `gorm:"type:enum('Y', 'N');default:'N';comment:'kolom untuk soft delete'" json:"NA"`
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
