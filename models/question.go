package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Question     string `json:"question" gorm:"not null"`
	Answer       string `json:"answer" gorm:"not null"`
	UserID       uint
	AssignmentID uint
	Topic        string `json:"topic" gorm:"not null"`
}
