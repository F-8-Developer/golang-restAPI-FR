package vaonline

import (
	"intrajasa-merchant-api-gateway/core/structs"
)

// Register register one new user in db, return a boolean value to make know success or not.
func GetVaPaymentStatus(gvp_req structs.GetVaPaymentStatusRequest) (gvp_res structs.GetVaPaymentStatusResponse) {
	gvp_res.MerchantId = gvp_req.MerchantId
	gvp_res.MerchantRefCode = gvp_req.MerchantRefCode
	gvp_res.VaNumber = gvp_req.VaNumber
	gvp_res.TotalAmount = gvp_req.TotalAmount
	gvp_res.PaymentStatus = 1
	gvp_res.ResponseCode = 200
	gvp_res.ResponseMsg = "Success generate VA Number"

	return gvp_res
}