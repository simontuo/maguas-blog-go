package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	UserID    int    `gorm:"index"`
	Type      string `gorm:"type:varchar(100)"`
	Title     string `gorm:"not null;unique"`
	Content   string `gorm:"type:text;not null"`
	ReadCount int
	IsPublic  bool `gorm:"type:tinyint(1);default:0;not null"`

	// one-to-one
	//User User
	// polymorphic
	Comments []Comment `gorm:"polymorphic:Commentable"`
}
