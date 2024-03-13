package model

import (
	"time"

	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type Logger struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Status    int            `json:"status"`
	Author    string         `json:"author" gorm:"index,priority:1; type:varchar(64)"`
	Message   string         `json:"message" gorm:"type:text"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (l *Logger) store() *gorm.DB {
	return utility.DB
}

func (l *Logger) Create() error {
	return l.store().Create(&l).Error
}

func (l *Logger) FindMany() ([]Logger, error) {
	loggers := []Logger{}
	err := l.store().Find(&loggers).Error
	return loggers, err
}

func (l *Logger) Delete() error {
	return l.store().Delete(&l).Error
}
