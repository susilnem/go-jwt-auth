package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`

	Roles []Role `gorm:"many2many:user_roles;" json:"roles"`
	Permissions []Permission `gorm:"many2many:user_permissions;" json:"permissions"`
}
