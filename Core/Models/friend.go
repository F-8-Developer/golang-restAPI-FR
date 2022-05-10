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

// Find friend request and return error info.
// err := Models.FindFriend(&frd, &usr, &friend_request, status);
func FindFriendRequest(frd *Friend, usr *User, friend_request *User, status string) error {
	err := Database.Mysql.
		Where("user_id = ?", usr.ID).
		Where("friend_request_id = ?", friend_request.ID).
		Where("status = ?", status).
		First(frd).Error
	return err
}

// Insert friend request which will be saved in database returning with error info
// err := Models.CreateFriend(&frd, &usr, &friend_request);
func CreateFriendRequest(frd *Friend, usr *User, friend_request *User) error {
	err := Database.Mysql.FirstOrCreate(&frd, Friend{UserID: usr.ID, FriendRequestID: friend_request.ID}).Error
	return err
}

// Update friend request status and return error info.
// err := Models.FindFriend(&frd, &usr, &friend_request, status);
func UpdateFriendRequestStatus(frd *Friend, usr *User, friend_request *User, status string) error {
	err := Database.Mysql.Model(&frd).
		Where("user_id = ?", usr.ID).
		Where("friend_request_id = ?", friend_request.ID).
		Update("status", status).Error
	return err
}