package cart

import "tugas-acp/models/base"

type CartResponse struct {
	base.BaseResponse
	Data []Cart `json:"data"`
}