package model

import (
	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type Content struct {
	ApplicationNumber string       `json:"applicationNumber" gorm:"primaryKey"`
	Application       *Application `json:"application" gorm:"foreignKey:ApplicationNumber"`

	AttributeID uint       `json:"attributeId" gorm:"primaryKey;autoIncrement:false"`
	Attribute   *Attribute `json:"attribute"`

	Value string `json:"value" gorm:"type:text"`
}

func (a *Content) store() *gorm.DB {
	return utility.DB
}

func (a *Content) Create(contents []*Content) error {
	return a.store().Create(&contents).Error
}

func (a *Content) FindMany(appId string) ([]*Content, error) {
	contents := []*Content{}
	err := a.store().Where("application_number = ?", appId).Preload("Application").Preload("Attribute").Find(&contents).Error
	return contents, err
}

func (a *Content) Delete() error {
	return a.store().Delete(&a).Error
}
