package model

import (
	"time"

	"gorm.io/gorm"
)

type ReportMaster struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(64)"`
	IsActive  bool           `json:"isActive"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Applications []*Application `json:"applications"`
	Attributes   []*Attribute   `json:"attributes"`
}
