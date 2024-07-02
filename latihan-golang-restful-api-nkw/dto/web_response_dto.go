package dto

type WebResponseDto[T any] struct {
	Data    T      `json:"data"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}
