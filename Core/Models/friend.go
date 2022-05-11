package Models

import (
	// "fmt"
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

type APIFriendRequest struct {
	Requests	[]FriendRequestList	`json:"requests"`
}

type FriendRequestList struct {
	Requestor	string	`json:"requestor"`
	Status		string	`json:"status"`
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
		First(&frd).Error
	return err
}

// List of friend and return error info.
// list, err := Models.FriendRequestListQuery();
func FriendRequestListQuery(usr *User) (response APIFriendRequest, err error) {
	var requestor, status string
	list := []FriendRequestList{}

	rows, err := Database.Mysql.Model(&Friend{}).
		Joins("JOIN users as friend_request on friend_request.id = friends.friend_request_id").
		Where("user_id = ?", usr.ID).
		Where("status IN (?)", []string{"pending", "accepted", "rejected"}).
		Select("friend_request.email, friends.status").Rows()
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
        rows.Scan(&requestor, &status)
		list = append(list, FriendRequestList{Requestor:requestor,Status:status})
    }
	response.Requests = list
	return
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