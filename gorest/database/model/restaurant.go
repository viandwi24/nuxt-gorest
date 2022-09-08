package model

import (
	"time"

	"gorm.io/gorm"
)

type Restaurant struct {
	RestaurantID uint64         `gorm:"primaryKey" json:"RestaurantID,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	UpdatedAt    time.Time      `json:"updatedAt,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"Name,omitempty"`
	Description  string         `json:"Description,omitempty"`
	Address      string         `json:"Address,omitempty"`
	Rating       uint64         `json:"Rating,omitempty"`
	avearageTime uint64         `json:"avearageTime,omitempty"`
}
