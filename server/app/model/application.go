package model

import (
	"time"

	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type Application struct {
	Number    string         `json:"number" gorm:"primaryKey"`
	Status    string         `json:"status" gorm:"type:varchar(32)"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	MasterReportID uint          `json:"masterReportId"`
	MasterReport   *MasterReport `json:"masterReport"`

	Contents []*Content `json:"contents"`

	UserEmail string `json:"userEmail"`
	User      *User  `json:"user" gorm:"foreignKey:UserEmail"`

	Reaction *Reaction `json:"reaction"`
}

func (a *Application) store() *gorm.DB {
	return utility.DB
}

func (a *Application) Create() error {
	return a.store().Create(&a).Error

}

func (a *Application) FindMany() ([]*Application, error) {
	applications := []*Application{}
	err := a.store().Preload("Contents").Preload("User").Preload("Reaction").Preload("MasterReport").Find(&applications).Error
	return applications, err
}

func (a *Application) FindOne() (*Application, error) {
	err := a.store().Where("number = ?", a.Number).Preload("Contents").Preload("User").Preload("MasterReport").First(&a).Error
	return a, err
}

func (a *Application) Delete() error {
	return a.store().Delete(&a).Error
}
