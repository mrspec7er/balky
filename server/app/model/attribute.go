package model

import (
	"time"

	"github.com/mrspec7er/balky/app/utils"
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
	return utils.DB
}

func (a *Attribute) Create() error {
	err := a.store().Create(&a).Error
	return err
}

func (a *Attribute) FindMany() ([]Attribute, error) {
	attributes := []Attribute{}
	err := a.store().Find(&attributes).Error
	return attributes, err
}

func (a *Attribute) Delete() error {
	err := a.store().Delete(&a).Error
	return err
}
