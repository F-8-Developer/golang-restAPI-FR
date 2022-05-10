package Api

import (
	"golang-restAPI-FR/Core/Structs"
	"golang-restAPI-FR/Core/Models"
)

var (
	success_response Structs.SuccessResponse
	error_response Structs.ErrorResponse
)

func FriendRequest(frq_req Structs.FriendRequestRequest) interface{} {
	success_response.Status = true
	error_response.Status = false
	var friend_requestor Models.User
	var user Models.User
	var friend Models.Friend
	friend.Status = "request"

	// store user email requestor and receiver
	friend_requestor.Email = frq_req.Requestor
	if err := Models.CreateUser(&friend_requestor); err != nil {
		error_response.Msg = err.Error()
		return error_response
	}
	user.Email = frq_req.To
	if err := Models.CreateUser(&user); err != nil {
		error_response.Msg = err.Error()
		return error_response
	}
	// ----------

	if err := Models.CreateFriend(&friend, &user, &friend_requestor); err != nil {
		error_response.Msg = err.Error()
		return error_response
	}

	return success_response
}
