package Models

import (
	"ToDoApp/Enums"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title  string
	Body   string
	Name   string
	Status Enums.Status `gorm:"type:varchar(20);default: 'To Do';"`
	Time   time.Time
}
