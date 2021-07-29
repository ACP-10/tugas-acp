package payment

import "tugas-acp/models/base"

type PaymentResponse struct {
	base.BaseResponse
	Data []Payment `json:"data"`
}