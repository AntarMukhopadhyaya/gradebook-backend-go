package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string         `json:"email" gorm:"unique"`
	Password    string         `json:"password"`
	Verified    bool           `json:"verified" gorm:"default:false"`
	Role        string         `json:"role" gorm:"not null"`
	Avatar      sql.NullString `json:"avatar" gorm:"default:null"`
	Questions   []Question     `json:"questions"`
	Assignments []Assignment   `json:"assignments"`
}
