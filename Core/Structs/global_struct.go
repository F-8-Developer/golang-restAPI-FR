package Structs

type SuccessResponse struct {
	Status	bool `json:"status"`
}

type ErrorResponse struct {
	Status	bool 	`json:"status"`
	Msg		string `json:"msg"`
}