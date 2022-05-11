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
// err := Models.FindFriend(&frd, []string{"pending"});
func FindFriendRequest(frd *Friend, status []string) error {
	err := Database.Mysql.
		Where("user_id = ?", frd.UserID).
		Where("friend_request_id = ?", frd.FriendRequestID).
		Where("status IN (?)", status).
		First(frd).Error
	return err
}

// Insert friend request which will be saved in database returning with error info
// err := Models.CreateFriend(&frd);
func CreateFriendRequest(frd *Friend) error {
	err := Database.Mysql.Create(&frd).Error
	return err
}

// Update friend request status and return error info.
// err := Models.FindFriend(&frd, status);
func UpdateFriendRequest(frd *Friend, status string) error {
	Database.Mysql.First(&frd)
	frd.Status = status
	err := Database.Mysql.Save(&frd).Error
	return err
}