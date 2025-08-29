package types

import "time"

type PostModel struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Description string `json:"description"`
}

type PostDetailModel struct {
	Image string `json:"image"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}