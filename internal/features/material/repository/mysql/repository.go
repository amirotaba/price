package materialRepo

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

func NewMysqlRepository(db *bun.DB) domain.MaterialRepository {
	return &mysqlRepository{
		Conn: db,
	}
}

func (m *mysqlRepository) CreateMaterial(ctx context.Context, material *entity.CreateMaterialRequest) error {
	var model entity.Material
	err := copier.Copy(&model, &material)
	if err != nil {
		return err
	}
	if _, err := m.Conn.NewInsert().Model(&model).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepository) ReadMaterial(ctx context.Context, request *domain.ReadRequest) ([]entity.Material, error) {
	var out []entity.Material
	db := m.Conn.NewSelect().Model(&out)

	for _, filter := range request.Filters {
		filter.Sub = "material"
		db.Where(filter.Field, filter.Value)
	}

	err := db.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (m *mysqlRepository) UpdateMaterial(ctx context.Context, material entity.Material) error {
	var model entity.Material
	err := copier.Copy(&model, &material)
	if err != nil {
		return err
	}

	if _, err := m.Conn.NewUpdate().Model(&model).Where("id = ?", material.ID).Exec(ctx); err != nil {
		return err
	}

	return nil
}
