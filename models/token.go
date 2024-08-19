package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"foreignKey:user_id"`
	Token  string `json:"token"`
}
