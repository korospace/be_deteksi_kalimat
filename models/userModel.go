package models

import "time"

type User struct {
	ID           int        `gorm:"primaryKey" json:"userId"`
	Name         string     `gorm:"type:varchar(255);" json:"name"`
	Email        string     `gorm:"type:varchar(255);" json:"email"`
	Password     string     `gorm:"type:text;" json:"password"`
	UserAccessID int        `gorm:"default:null;comment:'kolom untuk menyimpan id pada tabel user access'" json:"user_access_id"`
	UseridCreate int        `gorm:"default:null;comment:'kolom untuk user id yang membuat data'" json:"userid_create"`
	DateCreate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal pembuatan'" json:"date_create"`
	UseridUpdate int        `gorm:"default:null;comment:'kolom untuk user id yang mengupdate data'" json:"userid_update"`
	DateUpdate   *time.Time `gorm:"type:datetime;default:null;comment:'kolom untuk tanggal update'" json:"date_update"`
	NA           string     `gorm:"type:enum('Y', 'N');default:'N';comment:'kolom untuk soft delete'" json:"NA"`
}

type UserCreateReq struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	UserAccessID int    `json:"user_access_id"`
}

type UserUpdateReq struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	UserAccessID int    `json:"user_access_id"`
}

type Me struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	UserAccessID   int    `json:"user_access_id"`
	UserAccessName string `json:"user_access_name"`
}
