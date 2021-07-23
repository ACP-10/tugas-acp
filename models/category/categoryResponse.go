package category

import "tugas-acp/models/base"

type CategoryResponse struct {
	base.BaseResponse
	Data []Category `json:"data"`
}