package models_user

import (
	"golang.org/x/crypto/bcrypt"
	"myapp/database"
	"myapp/models/models_post"
)

// User defines the user in db
type User struct {
	ID        int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	FirstName string `gorm:"type:VARCHAR(100);NOT NULL" json:"first_name" validate:"required"`
	LastName  string `gorm:"type:VARCHAR(100);NOT NULL" json:"last_name" validate:"required"`
	Email     string `gorm:"type:VARCHAR(255);NOT NULL;UNIQUE" json:"email" validate:"required,email"`
	Password  string `gorm:"type:CHAR(60);NOT NULL" json:"password" validate:"required"`
	Posts     []models_post.Post
}

// CreateUserRecord creates a controllers_post record in the database
func (user *User) CreateUserRecord() error {
	result := database.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// HashPassword encrypts controllers_post password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks controllers_post password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
