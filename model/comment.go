package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserID  int    `gorm:"index"`
	Content string `gorm:"type:text;not null"`

	// one-to-one
	//User User
	// polymorphic
	CommentableID   int
	CommentableType string
}
