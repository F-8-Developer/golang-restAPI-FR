package Structs

import (
	"golang-restAPI-FR/Core/Models"
)

type CategoriesResponse struct {
	Categories		[]Models.Category	`json:"categories"`
	ResponseCode	int					`json:"responseCode"`
	ResponseMsg		string				`json:"responseMsg"`
}