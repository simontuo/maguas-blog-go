package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null;unique"` // string默认长度为150
	Phone    string `gorm:"not null;unique"`
	Email    string `gorm:"unique"`
	Avatar   string
	Password string

	// One-To-Many
	Articles []Article
	// many-to-many
	Likes []Article `gorm:"many2many:article_user;"`
}

func (u *User) Verify() (err error) {
	return nil
}
