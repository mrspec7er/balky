package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Message   string         `json:"message" gorm:"type:text"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	UserEmail string `json:"userEmail"`
	User      *User  `json:"user" gorm:"foreignKey:UserEmail"`

	ReferenceID uint     `json:"referenceId"`
	Reference   *Comment `json:"reference"`
}
