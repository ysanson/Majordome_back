package entity

import (
	"github.com/jinzhu/gorm"
)

type Realisation struct {
	gorm.Model
	ItemId uint32 `json:"item_id" db:"item_id" gorm:"column:item_id"`
	Item   Item   `gorm:"foreignKey:ItemId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Realisation) TableName() string {
	return "realisation"
}

func (s *Realisation) BeforeCreate(scope *gorm.Scope) {
	// uuid := uuid.NewV4()
	// scope.SetColumn("ID", uuid.String())
}

type APIRealisationListAll struct {
	ID        uint   `json:"id"`
	ItemId    uint   `json:"item_id"`
	Item      Item   `json:"item"`
	CreatedAt string `json:"created_at"`
}
