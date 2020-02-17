package dto

type CategoryDto struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Cover  string `json:"cover"`
	Slug   string `json:"slug"`
	Status int32  `json:"status"`
}
