package product

import "tugas-acp/models/base"

type ProductResponse struct {
	base.BaseResponse
	Data []Product `json:"data"`
}