package Structs

type FriendRequestRequest struct {
	Requestor	string `json:"requestor" validate:"required,email,max=255"`
	To			string `json:"to" validate:"required,email,max=255"`
}

type FriendBlockRequest struct {
	Requestor	string `json:"requestor" validate:"required,email,max=255"`
	Block		string `json:"block" validate:"required,email,max=255"`
}

type FriendListRequest struct {
	Email	string	`json:"email" validate:"required,email,max=255"`
}

type FriendBetweenRequest struct {
	Friends	[]string	`json:"friends" validate:"gt=0,dive,required,email"`
}