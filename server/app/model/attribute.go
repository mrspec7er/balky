package model

import (
	"time"

	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type Attribute struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Label      string         `json:"label" gorm:"type:varchar(64)"`
	Type       string         `json:"type" gorm:"type:varchar(32)"`
	IsRequired bool           `json:"isRequired"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	MasterReportID uint          `json:"masterReportId"`
	MasterReport   *MasterReport `json:"masterReport"`

	Contents []*Content `json:"contents"`
}

func (a *Attribute) store() *gorm.DB {
	return utility.DB
}

func (a *Attribute) Create(req []*Attribute) error {
	return a.store().Create(&req).Error
}

func (a *Attribute) FindMany() ([]Attribute, error) {
	attributes := []Attribute{}
	err := a.store().Preload("MasterReport").Find(&attributes).Error
	return attributes, err
}

func (a *Attribute) Delete() error {
	return a.store().Delete(&a).Error
}
