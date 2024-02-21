package model

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	Number		string		   `json:"number" gorm:"primaryKey"`
	Status	 	string 		   `json:"status" gorm:"type:varchar(32)"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt	gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	ReportMasterID uint
	ReportMaster *ReportMaster `json:"reportMaster"`

	Contents []*Content `json:"content"`

	UserID uint
	User *User `json:"user"`

	PicID uint
	Pic *User `json:"pic" gorm:"foreignKey:PicID"`
}