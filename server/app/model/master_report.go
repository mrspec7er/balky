package model

import (
	"time"

	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type MasterReport struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(64)"`
	IsActive  bool           `json:"isActive"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Applications []*Application `json:"applications"`
	Attributes   []*Attribute   `json:"attributes"`
}

func (m *MasterReport) store() *gorm.DB {
	return utility.DB
}

func (m *MasterReport) Create() error {
	err := m.store().Create(&m).Error
	return err
}

func (m *MasterReport) FindMany() ([]MasterReport, error) {
	masters := []MasterReport{}
	err := m.store().Find(&masters).Error
	return masters, err
}

func (m *MasterReport) Delete() error {
	err := m.store().Delete(&m).Error
	return err
}
