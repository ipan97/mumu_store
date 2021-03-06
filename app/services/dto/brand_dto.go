package dto

import (
	"time"
)

type BrandDto struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_atl"`
}
