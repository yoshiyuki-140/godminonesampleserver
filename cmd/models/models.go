package models

import (
	"gorm.io/gorm"
)

// HasMany
type User struct {
	gorm.Model
	Name      string `gorm:"size:255" json:"name"`
	Password  string `gorm:"size:255" json:"password"`
	SessionId int
	Tasks     []Task
}

type Task struct {
	gorm.Model
	Task        string `gorm:"type:text" json:"task"`
	UserID      int
	IsCompleted bool `gorm:"default:false" json:"is_completed"`
}
