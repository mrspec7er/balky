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

	ReferenceID *uint    `json:"referenceId"`
	Reference   *Comment `json:"reference"`

	UserEmail string `json:"userEmail" gorm:"type:varchar(64)"`
	User      *User  `json:"user" gorm:"foreignKey:UserEmail"`

	ApplicationNumber string       `json:"applicationNumber"`
	Application       *Application `json:"application" gorm:"foreignKey:ApplicationNumber"`
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

func (a *Comment) FindMany(appNumber string) ([]*Comment, error) {
	comments := []*Comment{}
	err := a.store().Where("application_number = ?", appNumber).Preload("Application").Preload("Reference").Preload("User").Find(&comments).Error
	return comments, err
}
