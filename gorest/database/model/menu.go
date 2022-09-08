package model

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	MenuID       uint64         `gorm:"primaryKey" json:"MenuID,omitempty"`
	RestaurantID uint64         `json:"-"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	UpdatedAt    time.Time      `json:"updatedAt,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"Name,omitempty"`
	Description  string         `json:"Description,omitempty"`
	Image        string         `json:"Image,omitempty"`
	Price        float64        `json:"Price,omitempty"`
	Category     string         `json:"Category,omitempty"`
}
