package entity

import "time"

type Recipe struct {
	ID         uint      `bun:"id,pk,autoincrement"`
	CreatedAt  time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt  time.Time `bun:",nullzero,default:current_timestamp"`
	Volume     float64
	MaterialID int
}

type RecipeResponse struct {
	ID        uint
	Volume    float64
	UpdatedAt time.Time
	Material  MaterialResponse
}

type CreateRecipeRequest struct {
	Volume     float64 `json:"volume"`
	MaterialID float64 `json:"material_id"`
}

type ReadRecipeRequest struct {
	Usual *UsualRequest
}

type UpdateRecipeRequest struct {
	ID         uint    `json:"id"`
	Volume     float64 `json:"volume"`
	MaterialID float64 `json:"material_id"`
	UpdatedAt  time.Time
}
