package model

import "github.com/jinzhu/gorm"

//Note with title and completed
type Note struct {
	gorm.Model
	Title     string `json:"title" gorm:"type:varchar(255);not null"`
	Completed bool   `json:"completed" gorm:"type:bool"`
	UserID    uint   `json:"userid" gorm:"type:int"`
}
