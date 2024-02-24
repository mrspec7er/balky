package model

import (
	"time"

	"github.com/mrspec7er/balky/app/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"unique; index,priority:1; type:varchar(128)"`
	UID       string         `json:"uid" gorm:"index,priority:2; type:varchar(64)"`
	Name      string         `json:"name" gorm:"index,priority:3; type:varchar(256)"`
	Password  string         `json:"password" gorm:"type:varchar(256)"`
	Role      string         `json:"role" gorm:"type:varchar(32)"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Applications []*Application `json:"applications"`
}

func (u *User) store() *gorm.DB {
	return utils.DB
}

func (u *User) Create() error {
	err := u.store().Create(&u).Error
	return err
}

func (u *User) FindMany() ([]User, error) {
	users := []User{}
	err := u.store().Find(&users).Error
	return users, err
}