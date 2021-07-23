package customer

import "tugas-acp/models/base"

type CustomerResponse struct {
	base.BaseResponse
	Data []Customer `json:"data"`
}