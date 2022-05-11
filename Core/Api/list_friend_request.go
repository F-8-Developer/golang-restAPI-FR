package Api

import (
	"golang-restAPI-FR/Core/Structs"
	"golang-restAPI-FR/Core/Models"
)

func ListFriendRequest(frl_req Structs.FriendListRequest) interface{} {
	error_response.Success = false
	var user Models.User

	// find user email
	if err := Models.FindUser(&user, frl_req.Email); err != nil {
		error_response.Msg = "Email user not found"
		return error_response
	}
	// ----------

	// get list of friend request
	list, err := Models.FriendRequestListQuery(&user)
	if err != nil {
		error_response.Msg = "Invalid friend request list"
		return error_response
	}

	return list
}

func ListFriend(frl_req Structs.FriendListRequest) interface{} {
	error_response.Success = false
	var user Models.User

	// find user email
	if err := Models.FindUser(&user, frl_req.Email); err != nil {
		error_response.Msg = "Email user not found"
		return error_response
	}
	// ----------

	// get list of friend request
	list, err := Models.FriendListQuery(&user)
	if err != nil {
		error_response.Msg = "Invalid friend list"
		return error_response
	}

	return list
}

func ListFriendBetween(frb_req Structs.FriendBetweenRequest) interface{} {
	error_response.Success = false
	var user1, user2 Models.User

	// find user email
	if err := Models.FindUser(&user1, frb_req.Friends[0]); err != nil {
		error_response.Msg = "Email " + frb_req.Friends[0] + " user not found"
		return error_response
	}

	if err := Models.FindUser(&user2, frb_req.Friends[1]); err != nil {
		error_response.Msg = "Email " + frb_req.Friends[1] + " user not found"
		return error_response
	}
	// ----------

	// get list of friend request
	list, err := Models.FriendListBetweenQuery(user1, user2)
	if err != nil {
		error_response.Msg = "Invalid friend list between"
		return error_response
	}

	return list
}