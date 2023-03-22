package recipeRepo

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

func NewMysqlRepository(db *bun.DB) domain.RecipeRepository {
	return &mysqlRepository{
		Conn: db,
	}
}

func (m *mysqlRepository) CreateRecipe(ctx context.Context, request entity.CreateRecipeRequest) error {
	var model entity.Recipe
	err := copier.Copy(&model, &request)
	if err != nil {
		return err
	}
	if _, err := m.Conn.NewInsert().Model(&model).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (m *mysqlRepository) ReadRecipe(ctx context.Context, request *domain.ReadRequest) ([]entity.Recipe, error) {
	var out []entity.Recipe
	db := m.Conn.NewSelect().Model(&out)

	for _, filter := range request.Filters {
		filter.Sub = "recipe"
		db.Where(filter.Field, filter.Value)
	}

	err := db.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (m *mysqlRepository) UpdateRecipe(ctx context.Context, request entity.UpdateRecipeRequest) error {
	var model entity.Recipe
	err := copier.Copy(&model, &request)
	if err != nil {
		return err
	}

	if _, err := m.Conn.NewUpdate().Model(&model).Where("id = ?", request.ID).Exec(ctx); err != nil {
		return err
	}

	return nil
}
