package model

import (
	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type Reaction struct {
	ApplicationNumber string       `json:"applicationNumber" gorm:"primaryKey"`
	Application       *Application `json:"application" gorm:"foreignKey:ApplicationNumber"`

	Users []User `json:"users" gorm:"many2many:user_reactions;foreignKey:ApplicationNumber;References:Email;"`
}

func (a *Reaction) store() *gorm.DB {
	return utility.DB
}

func (a *Reaction) Create(reactions []*Reaction) error {
	err := a.store().Create(&reactions).Error
	return err
}

func (a *Reaction) FindMany(appId string) ([]*Reaction, error) {
	reactions := []*Reaction{}
	err := a.store().Where("application_number = ?", appId).Preload("Application").Preload("Attribute").Find(&reactions).Error
	return reactions, err
}

func (a *Reaction) Delete() error {
	err := a.store().Delete(&a).Error
	return err
}
