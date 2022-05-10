package Models

import (
	"golang-restAPI-FR/Database"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID				uint	`json:"-";gorm:"primaryKey"`
	Email			string	`gorm:"column:email;unique_index"`
}

func (usr *User) TableName() string {
	return "users"
}

// Find user by email address and return error info.
// err := Models.FindUser(&user, "user_email_address")
func FindUser(usr *User, email string) error {
	err := Database.Mysql.Where("email = ?", email).First(usr).Error
	return err
}

// Insert user which will be saved in database returning with error info
// if err := Models.CreateUser(&user); err != nil { ... }
func CreateUser(usr *User) error {
	err := Database.Mysql.FirstOrCreate(&usr, User{Email: usr.Email}).Error
	return err
}