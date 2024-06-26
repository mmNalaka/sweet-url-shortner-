package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	Id        uint   `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Color     string `gorm:"not null;default:'#8B93FF'"`
	Links     []Link `gorm:"many2many:link_tag;"` // many to many
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Tag) TableName() string {
	return "tag"
}
