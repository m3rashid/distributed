package utils

type PaginationRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PaginationResponse[T any] struct {
	Data     []T `json:"data"`
	Total    int `json:"total"`
	Page     int `json:"page"`
	Limit    int `json:"limit"`
	NextPage int `json:"next_page"`
	PrevPage int `json:"prev_page"`
}
