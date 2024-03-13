package model

import (
	"time"

	"github.com/lib/pq"
	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type Reaction struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	LikesBy   pq.StringArray `json:"likesBy" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	ApplicationNumber string       `json:"applicationNumber"`
	Application       *Application `json:"application" gorm:"foreignKey:ApplicationNumber"`
}

func (a *Reaction) store() *gorm.DB {
	return utility.DB
}

func (a *Reaction) Create(reaction *Reaction) error {
	return a.store().Save(&reaction).Error
}

func (a *Reaction) FindMany(appNumber string) ([]*Reaction, error) {
	reactions := []*Reaction{}
	err := a.store().Where("application_number = ?", appNumber).Preload("Application").Preload("Attribute").Find(&reactions).Error
	return reactions, err
}

func (a *Reaction) FindOne() (*Reaction, error) {
	err := a.store().Where("application_number = ?", a.ApplicationNumber).First(&a).Error
	return a, err
}
