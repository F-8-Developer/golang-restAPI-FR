package Structs

import (
	"golang-restAPI-FR/Core/Models"
)

type ProductDetailRequest struct {
	ProductID	uint	`json:"productID" validate:"required,uint-len-10"`
}

type ProductDetailResponse struct {
	Product			Models.Product	`json:"product"`
	ResponseCode	int				`json:"responseCode"`
	ResponseMsg		string			`json:"responseMsg"`
}

type ProductsResponse struct {
	Products		[]Models.Product	`json:"products"`
	ResponseCode	int					`json:"responseCode"`
	ResponseMsg		string				`json:"responseMsg"`
}