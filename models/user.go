package models

import (
	"golang.org/x/crypto/bcrypt"
	"myapp/database"
)

// User defines the user in db
type User struct {
	ID        int    `gorm:"primary_key;AUTO_INCREMENT"`
	FirstName string `gorm:"type:VARCHAR(100);NOT NULL"`
	LastName  string `gorm:"type:VARCHAR(100);NOT NULL"`
	Email     string `gorm:"type:VARCHAR(255);NOT NULL;UNIQUE"`
	Password  string `gorm:"type:CHAR(60);NOT NULL"`
	Posts     []Post
}

// FindUserByEmail searches for the user in the database by email
func (user *User) FindUserByEmail(email string) error {
	result := database.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// CreateUserRecord creates a controllers_post record in the database
func (user *User) CreateUserRecord() error {
	result := database.DB.Create(user)
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
