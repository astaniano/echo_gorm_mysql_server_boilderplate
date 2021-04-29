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

func (post *Post) GetAllPosts() ([]Post, error) {
	var postsFromDB []Post
	result := database.DB.Find(&postsFromDB)
	if result.Error != nil {
		return postsFromDB, result.Error
	}

	return postsFromDB, nil
}

func (post *Post) GetPostByID(postId int) (*Post, error) {
	foundPost := new(Post)
	result := database.DB.First(foundPost, postId)
	if result.Error != nil {
		return foundPost, result.Error
	}

	return foundPost, nil
}

// CreatePostRecord creates a post record in the database
func (post *Post) CreatePostRecord() error {
	result := database.DB.Create(post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdatePostRecord updates a post record in the database
func (post *Post) UpdatePostRecord() error {
	result := database.DB.Save(post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeletePostByID updates a post record in the database
func (post *Post) DeletePost() error {
	result := database.DB.Unscoped().Delete(post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
