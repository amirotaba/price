package entity

import (
	"time"
)

type Item struct {
	ID        uint      `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,default:current_timestamp"`
	Name      string
	Cost      float64
	Price     float64
}

type ItemResponse struct {
	ID        uint
	Name      string
	Cost      float64
	Price     float64
	Recipes   []RecipeResponse
	UpdatedAt time.Time
}

type CreateItemRequest struct {
	Name    string `json:"name"`
	Cost    float64
	Price   float64
	Recipes []uint `json:"recipes"`
}

type ReadItemRequest struct {
	Usual *UsualRequest
}

type UpdateItemRequest struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Recipes   []uint `json:"recipes"`
	Cost      float64
	UpdatedAt time.Time
}
