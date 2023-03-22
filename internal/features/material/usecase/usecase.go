package materialUsecase

import (
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
	"gopkg.in/go-playground/validator.v9"
	"price/internal/domain"
	"price/internal/entity"
	"price/internal/utils"
	"time"
)

type usecase struct {
	Material domain.MaterialRepository
}

func NewUseCase(r domain.Repositories) domain.MaterialUsecase {
	return &usecase{
		Material: r.Material,
	}
}

func (u *usecase) ReadMaterials(ctx context.Context, request entity.ReadMaterialsRequest) ([]entity.MaterialResponse, error) {
	var response entity.MaterialResponse
	var responses []entity.MaterialResponse
	convRequest, err := utils.ConvertRequest(request.Usual)
	if err != nil {
		return nil, err
	}

	materials, err := u.Material.ReadMaterial(ctx, convRequest)

	for _, material := range materials {
		if err := copier.Copy(&response, &material); err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (u *usecase) UpdateMaterials(ctx context.Context, request entity.UpdateMaterialRequest) error {
	var filters []domain.Filter
	for _, i := range request.Sheet.Rows {
		//check for not duplicating
		filters = nil
		filter := utils.AddFilter("name", i[0].Value, "equal")
		filters = append(filters, filter)
		readRequest := &domain.ReadRequest{
			Filters: filters,
		}
		exist, err := u.Material.ReadMaterial(ctx, readRequest)
		if err != nil {
			return err
		}
		convMat, err := utils.ConvertMaterial(i)
		if err != nil {
			return err
		}
		//validate not null field
		v := validator.New()
		if err := v.Struct(convMat); err != nil {
			return err
		}
		switch exist {
		case nil:
			if err := u.Material.CreateMaterial(ctx, convMat); err != nil {
				return err
			}
		default:
			if convMat.Price != exist[0].Price || convMat.Unit != exist[0].Unit {
				if err := copier.Copy(&exist[0], &convMat); err != nil {
					return err
				}
				exist[0].UpdatedAt = time.Now()
				if err := u.Material.UpdateMaterial(ctx, exist[0]); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
