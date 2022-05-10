package Models

import (
	"golang-restAPI-FR/Database"
	_ "github.com/go-sql-driver/mysql"
)

type Friend struct {
	ID				uint	`json:"-";gorm:"primaryKey"`
	UserID			uint	`json:"-";gorm:"column:user_id"`
	FriendRequestID	uint	`json:"-";gorm:"column:friend_request_id"`
	User			User	`gorm:"foreignKey:UserID"`
	Friend			User	`gorm:"foreignKey:friendID"`
	Status			string	`gorm:"column:status"`
}

func (frd *Friend) TableName() string {
	return "friends"
}

// Insert friend request which will be saved in database returning with error info
// err := Models.CreateFriend(&frd);
func CreateFriend(frd *Friend, usr *User, friend_request *User) error {
	err := Database.Mysql.FirstOrCreate(&frd, Friend{UserID: usr.ID, FriendRequestID: friend_request.ID}).Error
	return err
}