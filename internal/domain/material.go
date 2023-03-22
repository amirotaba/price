package domain

import (
	"golang.org/x/net/context"
	"price/internal/entity"
)

type MaterialUsecase interface {
	ReadMaterials(ctx context.Context, request entity.ReadMaterialsRequest) ([]entity.MaterialResponse, error)
	UpdateMaterials(ctx context.Context, material entity.UpdateMaterialRequest) error
}

type MaterialRepository interface {
	CreateMaterial(ctx context.Context, material *entity.CreateMaterialRequest) error
	ReadMaterial(ctx context.Context, request *ReadRequest) ([]entity.Material, error)
	UpdateMaterial(ctx context.Context, material entity.Material) error
}
