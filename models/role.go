package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255;;" json:"description"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type Permission struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255;;" json:"description"`
}