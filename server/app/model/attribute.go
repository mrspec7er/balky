package model

import (
	"time"

	"gorm.io/gorm"
)

type Attribute struct {
	ID				uint		   `json:"id" gorm:"primaryKey"`
	Label			string		   `json:"name" gorm:"type:varchar(64)"`
	Type 			string		   `json:"type" gorm:"type:varchar(32)"`
	IsRequired 		bool 		   `json:"isRequired"`
	CreatedAt   	time.Time      `json:"createdAt"`
	UpdatedAt   	time.Time      `json:"updatedAt"`
	DeletedAt		gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	ReportMasterID 	uint
	ReportMaster 	*ReportMaster  `json:"reportMaster"`

	Contents []*Content `json:"contents"`
}