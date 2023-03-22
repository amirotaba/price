package domain

import (
	"golang.org/x/net/context"
	"price/internal/entity"
)

type ItemUsecase interface {
	CreateItem(ctx context.Context, request entity.CreateItemRequest) error
	ReadItem(ctx context.Context, request entity.ReadItemRequest) ([]entity.ItemResponse, error)
	UpdateItem(ctx context.Context, request entity.UpdateItemRequest) error
}

type ItemRepository interface {
	CreateItem(ctx context.Context, request entity.CreateItemRequest) (entity.Item, error)
	ReadItem(ctx context.Context, request *ReadRequest) ([]entity.Item, error)
	UpdateItem(ctx context.Context, request entity.UpdateItemRequest) error
}
