package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        string  `gorm:"primarykey;index:idx_user_id;unique;size:50;not null;"`
	Username      string  `gorm:"index:idx_username;unique;size:30;not null;"`
	Password      string  `gorm:"size:30;not null;"`
	Mobile        string  `gorm:"index:idx_mobile;unique;size:14;not null;"`
	Email         string  `gorm:"index:idx_email;unique;size:50;"`
	EmailVerified bool    `gorm:"not null;default:false;"`
	RoleID        uint    `gorm:"not null;"`
	Role          Role    `gorm:"foreignKey:RoleID"`
	Groups        []Group `gorm:"many2many:user_group;"`
}

type Role struct {
	gorm.Model
	Role        string       `gorm:"index:idx_role;size:100;unique;not null;default:'普通用户'"`
	Users       []User       `gorm:"foreignKey:RoleID"`
	Permissions []Permission `gorm:"many2many:role_permission"`
}

type Group struct {
	gorm.Model
	Group         string `gorm:"index:idx_group;size:100;"`
	ParentGroupID uint   `gorm:"not null;"`
	ParentGroup   *Group `gorm:"foreignKey:ParentGroupID"`
	Users         []User `gorm:"many2many:user_group;"`
}

type Permission struct {
	gorm.Model
	Permission string `gorm:"index:idx_permission;size:100"`
	Roles      []Role `gorm:"many2many:role_permission"`
}
