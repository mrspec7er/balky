package model

import (
	"time"

	"github.com/mrspec7er/balky/app/utils"
	"gorm.io/gorm"
)

type ReportMaster struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(64)"`
	IsActive  bool           `json:"isActive"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Applications []*Application `json:"applications"`
	Attributes   []*Attribute   `json:"attributes"`
}

func (m *ReportMaster) store() *gorm.DB {
	return utils.DB
}

func (m *ReportMaster) Create() error {
	err := m.store().Create(&m).Error
	return err
}

func (m *ReportMaster) FindMany() ([]ReportMaster, error) {
	masters := []ReportMaster{}
	err := m.store().Find(&masters).Error
	return masters, err
}

func (m *ReportMaster) Delete() error {
	err := m.store().Delete(&m).Error
	return err
}
