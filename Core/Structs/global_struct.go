package Structs

type SuccessResponse struct {
	Success	bool `json:"success"`
}

type ErrorResponse struct {
	Success	bool 	`json:"success"`
	Msg		string	`json:"msg"`
}