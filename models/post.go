package models

import (
	"myapp/database"
)

// User defines the post in db
type Post struct {
	ID     int    `gorm:"primary_key;AUTO_INCREMENT" json:"id" xml:"id"`
	UserID int    `gorm:"type:BIGINT;NOT NULL" json:"user_id" xml:"user_id"`
	Title  string `gorm:"type:VARCHAR(300);NOT NULL" json:"title" xml:"title" validate:"required"`
	Body   string `gorm:"type:BLOB;NOT NULL" json:"body" xml:"body" validate:"required"`
}

func (post *Post) GetPostByID(postId int) (*Post, error) {
	var FoundPost *Post
	result := database.DB.First(&FoundPost, postId)
	if result.Error != nil {
		return FoundPost, result.Error
	}

	return FoundPost, nil
}

// CreatePostRecord creates a post record in the database
func (post *Post) CreatePostRecord() error {
	result := database.DB.Create(&post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdatePostRecord updates a post record in the database
func (post *Post) UpdatePostRecord() error {
	result := database.DB.Save(&post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

//// DeletePostRecord updates a post record in the database
//func (post *Post) DeletePostRecord() error {
//	result := database.DB.Save(&post)
//	if result.Error != nil {
//		return result.Error
//	}
//
//	return nil
//}
