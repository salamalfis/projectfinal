package model

import ()

// Photo ...

type Photo struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}


