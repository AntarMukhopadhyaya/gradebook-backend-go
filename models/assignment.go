package models

import (
	"gorm.io/gorm"
	"time"
)

type Assignment struct {
	gorm.Model
	Topic             string `json:"topic" gorm:"not null"`
	UserID            uint
	DueDate           time.Time `json:"due_date" gorm:"not null"`
	AssignedClass     string    `json:"assigned_class" gorm:"not null"`
	NumberOfQuestions int       `json:"number_of_questions" gorm:"not null"`
	Questions         []Question
}
