package Structs

type FriendRequestRequest struct {
	Requestor	string `json:"requestor" validate:"required,email,max=255"`
	To			string `json:"to" validate:"required,email,max=255"`
}