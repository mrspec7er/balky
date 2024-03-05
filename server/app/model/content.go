package model

import (
	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type Content struct {
	ApplicationNumber string       `gorm:"primaryKey"`
	Application       *Application `json:"application" gorm:"foreignKey:ApplicationNumber"`

	AttributeID uint       `gorm:"primaryKey;autoIncrement:false"`
	Attribute   *Attribute `json:"attribute"`

	Value string `json:"value" gorm:"type:text"`
}

func (a *Content) store() *gorm.DB {
	return utility.DB
}

func (a *Content) Create() error {
	err := a.store().Create(&a).Error
	return err
}

func (a *Content) FindMany() ([]*Content, error) {
	contents := []*Content{}
	err := a.store().Find(&contents).Error
	return contents, err
}

func (a *Content) Delete() error {
	err := a.store().Delete(&a).Error
	return err
}
