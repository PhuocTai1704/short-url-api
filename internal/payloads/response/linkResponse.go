package response

import "short-url-api/internal/payloads/dto"

type LinkResponse struct {
	Data  []dto.LinkDTO `json:"data"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
	IsLastPage bool  `json:"isLastPage"`
}