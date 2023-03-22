package itemRecipeRepo

import (
	"github.com/jinzhu/copier"
	"github.com/uptrace/bun"
	"golang.org/x/net/context"
	"price/internal/domain"
	"price/internal/entity"
)

type mysqlRepository struct {
	Conn *bun.DB
}

func NewMysqlRepository(db *bun.DB) domain.ItemRecipeRepository {
	return &mysqlRepository{
		Conn: db,
	}
}

func (m *mysqlRepository) CreateItemRecipe(ctx context.Context, request *entity.CreateItemRecipeRequest) error {
	var model entity.ItemRecipe
	err := copier.Copy(&model, &request)
	if err != nil {
		return err
	}
	if _, err := m.Conn.NewInsert().Model(&model).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) ReadItemRecipe(ctx context.Context, request *domain.ReadRequest) ([]entity.ItemRecipe, error) {
	var out []entity.ItemRecipe
	db := m.Conn.NewSelect().Model(&out)

	for _, filter := range request.Filters {
		filter.Sub = "item_recipe"
		db.Where(filter.Field, filter.Value)
	}

	err := db.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (m *mysqlRepository) UpdateItemRecipe(ctx context.Context, request *entity.UpdateItemRecipeRequest) error {
	var model entity.ItemRecipe
	err := copier.Copy(&model, &request)
	if err != nil {
		return err
	}

	if _, err := m.Conn.NewUpdate().Model(&model).Where("id = ?", request.ID).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (m *mysqlRepository) DeleteItemRecipe(ctx context.Context, request *entity.DeleteItemRecipeRequest) error {
	var model entity.ItemRecipe
	err := copier.Copy(&model, &request)
	if err != nil {
		return err
	}

	if _, err := m.Conn.NewDelete().Model(&model).Where("id = ?", request.ID).Exec(ctx); err != nil {
		return err
	}

	return nil
}
