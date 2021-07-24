package cartitem

import "tugas-acp/models/base"

type CartItemResponse struct {
	base.BaseResponse
	Data []CartItem `json:"data"`
}