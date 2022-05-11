package Api

import (
	"golang-restAPI-FR/Core/Structs"
	"golang-restAPI-FR/Core/Models"
)

func UpdateFriendRequest(frq_req Structs.FriendRequestRequest, find_status string, status string) interface{} {
	success_response.Success = true
	error_response.Success = false
	var friend_requestor Models.User
	var user Models.User
	var friend Models.Friend
	friend.Status = status

	// find user email requestor and receiver
	if err := Models.FindUser(&friend_requestor, frq_req.Requestor); err != nil {
		error_response.Msg = "Email Requestor not found"
		return error_response
	}

	if err := Models.FindUser(&user, frq_req.To); err != nil {
		error_response.Msg = "Email To user not found"
		return error_response
	}
	// ----------

	// create friend request
	friend.UserID = user.ID
	friend.FriendRequestID = friend_requestor.ID

	if err := Models.FindFriendRequest(&friend, []string{find_status}); err != nil {
		error_response.Msg = "Friend " + find_status + " not found"
		return error_response
	}

	if err := Models.UpdateFriendRequest(&friend, status); err != nil {
		error_response.Msg = err.Error()
		return error_response
	}

	return success_response
}
