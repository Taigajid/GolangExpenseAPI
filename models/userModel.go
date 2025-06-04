package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"->"`
}

// BeforeSave is a GORM hook that gets called before saving the user to the database
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Only hash the password if it's not already hashed
	if len(u.Password) > 0 && len(u.Password) < 60 {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}

	return nil
}
