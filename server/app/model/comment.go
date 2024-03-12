package model

import (
	"time"

	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Message   string         `json:"message" gorm:"type:text"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	UserEmail string   `json:"userEmail"`
	User      *Comment `json:"user" gorm:"foreignKey:UserEmail"`

	ReferenceID uint     `json:"referenceId"`
	Reference   *Comment `json:"reference"`
}

func (c *Comment) store() *gorm.DB {
	return utility.DB
}

func (c *Comment) Create() error {
	return c.store().Create(&c).Error
}

func (c *Comment) Delete() error {
	return c.store().Delete(&c).Error
}
