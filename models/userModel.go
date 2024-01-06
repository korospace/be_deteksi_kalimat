package models

import "time"

type User struct {
	ID           int        `gorm:"primaryKey"`
	Name         string     `gorm:"type:varchar(255);"`
	Email        string     `gorm:"type:varchar(255);"`
	Password     string     `gorm:"type:text;"`
	UserAccessID int        `gorm:"default:null;comment:'kolom untuk menyimpan id pada tabel user access'"`
	UseridCreate int        `gorm:"default:null;comment:'kolom untuk user id yang membuat data'"`
	DateCreate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal pembuatan'"`
	UseridUpdate int        `gorm:"default:null;comment:'kolom untuk user id yang mengupdate data'"`
	DateUpdate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal update'"`
	NA           string     `gorm:"type:enum('Y', 'N');default:'N';comment:'kolom untuk soft delete'"`
}

type Me struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"type:varchar(255);"`
	Email        string `gorm:"type:varchar(255);"`
	UserAccessID int    `gorm:"default:null;comment:'kolom untuk menyimpan id pada tabel user access'"`
}
