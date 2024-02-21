package models

import (
	"time"
)

type Titan struct {
	ID        uint      `json:"id" gorm:"column:Id;type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Name      string    `json:"name" gorm:"column:Name"`
	Age       int       `json:"age" gorm:"column:Age"`
	Power     string    `json:"power" gorm:"column:Power"`
	CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:UpdatedAt"`

	CreatedByID uint `json:"created_by_id" gorm:"column:CreatedByID;type:INT(10) UNSIGNED"`
	User        User `json:"user" gorm:"foreignKey:CreatedByID;references:ID"`
}

func (Titan) TableName() string {
	return "Titan"
}
