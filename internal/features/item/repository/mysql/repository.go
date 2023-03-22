package itemRepo

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

func NewMysqlRepository(db *bun.DB) domain.ItemRepository {
	return &mysqlRepository{
		Conn: db,
	}
}

func (m *mysqlRepository) CreateItem(ctx context.Context, request entity.CreateItemRequest) (entity.Item, error) {
	var model entity.Item
	err := copier.Copy(&model, &request)
	if err != nil {
		return entity.Item{}, err
	}

	if _, err := m.Conn.NewInsert().Model(&model).Exec(ctx); err != nil {
		return entity.Item{}, err
	}
	return model, nil
}

func (m *mysqlRepository) ReadItem(ctx context.Context, request *domain.ReadRequest) ([]entity.Item, error) {
	var out []entity.Item
	db := m.Conn.NewSelect().Model(&out)

	for _, filter := range request.Filters {
		filter.Sub = "item"
		db.Where(filter.Field, filter.Value)
	}

	err := db.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (m *mysqlRepository) UpdateItem(ctx context.Context, request entity.UpdateItemRequest) error {
	var model entity.Item
	err := copier.Copy(&model, &request)
	if err != nil {
		return err
	}

	if _, err := m.Conn.NewUpdate().Model(&model).Where("id = ?", request.ID).Exec(ctx); err != nil {
		return err
	}

	return nil
}
