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

	MasterReportID uint
	MasterReport   *MasterReport `json:"masterReport"`

	Contents []*Content `json:"contents"`

	UserID uint
	User   *User `json:"user"`
}

func (a *Application) store() *gorm.DB {
	return utility.DB
}

func (a *Application) Create() error {
	err := a.store().Create(&a).Error
	return err
}

func (a *Application) FindMany() ([]*Application, error) {
	applications := []*Application{}
	err := a.store().Find(&applications).Error
	return applications, err
}

func (a *Application) FindOne() (*Application, error) {
	err := a.store().Where("number = ?", a.Number).First(&a).Error
	return a, err
}

func (a *Application) Delete() error {
	err := a.store().Delete(&a).Error
	return err
}
