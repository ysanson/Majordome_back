package entity

import (
	//"uuid"
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Name      	string  `json:"name" db:"name" gorm:"varchar(255); column:name"`
	Logo 	  	string  `json:"logo" db:"logo" gorm:"varchar(255); column:logo"`
	Description string `json:"description" db:"description" gorm:"varchar(255); column:description"`	
}

func (Item) TableName() string {
	return "item"
}

func (s *Item) BeforeCreate(scope *gorm.Scope) {
	// uuid := uuid.NewV4()
	// scope.SetColumn("ID", uuid.String())
}

type APIItemListAll struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Logo       string  `json:"logo"`
	Description string `json:"description"`
	CreatedAt  string  `json:"created_at"`
}