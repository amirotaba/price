package entity

import (
	"gopkg.in/Iwark/spreadsheet.v2"
	"time"
)

type Material struct {
	ID         uint      `bun:"id,pk,autoincrement"`
	CreatedAt  time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt  time.Time `bun:",nullzero,default:current_timestamp"`
	Name       string    `validate:"required"`
	Unit       float64   `validate:"required"`
	Price      float64   `validate:"required"`
	RealPrice  float64
	Efficiency float64
}

type MaterialResponse struct {
	ID         uint
	Name       string
	Unit       float64
	Price      float64
	RealPrice  float64
	Efficiency float64
	UpdatedAt  time.Time
}

type CreateMaterialRequest struct {
	Name       string
	Unit       float64
	Price      float64
	RealPrice  float64
	Efficiency float64
}

type ReadMaterialsRequest struct {
	Usual *UsualRequest
}

type UpdateMaterialRequest struct {
	Sheet *spreadsheet.Sheet
}
