package itemUsecase

import (
	"errors"
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
	"price/internal/domain"
	"price/internal/entity"
	"price/internal/utils"
	"strconv"
	"time"
)

type usecase struct {
	Item       domain.ItemRepository
	Recipe     domain.RecipeRepository
	Material   domain.MaterialRepository
	ItemRecipe domain.ItemRecipeRepository
}

func NewUseCase(r domain.Repositories) domain.ItemUsecase {
	return &usecase{
		Item:       r.Item,
		Recipe:     r.Recipe,
		Material:   r.Material,
		ItemRecipe: r.ItemRecipe,
	}
}

func (u *usecase) CreateItem(ctx context.Context, request entity.CreateItemRequest) error {
	var filters []domain.Filter
	var materialCost float64
	var materialCosts []float64
	//check if not duplicate
	filters = nil
	filter := utils.AddFilter("name", request.Name, "equal")
	readRequest := &domain.ReadRequest{
		Filters: append(filters, filter),
	}
	item, err := u.Item.ReadItem(ctx, readRequest)
	if err != nil {
		return err
	}
	if item != nil {
		return errors.New("this item exists")
	}
	for _, recipe := range request.Recipes {
		//check recipe existence
		filters = nil
		filter = utils.AddFilter("id", uint(recipe), "equal")
		readRequest.Filters = append(filters, filter)
		recipeResponse, _ := u.Recipe.ReadRecipe(ctx, readRequest)
		if recipeResponse == nil {
			recipeID := strconv.Itoa(int(recipeResponse[0].ID))
			return errors.New("recipe id " + recipeID + "not found!")
		}

		//calculate recipe cost
		filters = nil
		filter = utils.AddFilter("id", recipeResponse[0].MaterialID, "equal")
		readRequest.Filters = append(filters, filter)
		material, err := u.Material.ReadMaterial(ctx, readRequest)
		if err != nil {
			return err
		}
		materialCost = utils.CalculateCost(material, recipeResponse)
		materialCosts = append(materialCosts, materialCost)
	}

	//calculate overall cost
	request.Cost = utils.SliceSum(materialCosts)

	response, err := u.Item.CreateItem(ctx, request)
	if err != nil {
		return err
	}
	//create related item-recipe
	for _, recipe := range request.Recipes {
		itemRecipeRequest := &entity.CreateItemRecipeRequest{
			ItemID:   response.ID,
			RecipeID: recipe,
		}
		if err := u.ItemRecipe.CreateItemRecipe(ctx, itemRecipeRequest); err != nil {
			return err
		}
	}
	return nil
}

func (u *usecase) ReadItem(ctx context.Context, request entity.ReadItemRequest) ([]entity.ItemResponse, error) {
	var response entity.ItemResponse
	var responses []entity.ItemResponse
	convRequest, err := utils.ConvertRequest(request.Usual)
	if err != nil {
		return nil, err
	}

	recipes, err := u.Recipe.ReadRecipe(ctx, convRequest)

	for _, recipe := range recipes {
		if err := copier.Copy(&response, &recipe); err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (u *usecase) UpdateItem(ctx context.Context, request entity.UpdateItemRequest) error {
	var filter domain.Filter
	var filters []domain.Filter
	var materialCost float64
	var materialCosts []float64
	for _, recipe := range request.Recipes {
		//check recipe existence
		filter = utils.AddFilter("id", uint(recipe), "equal")
		readRequest := &domain.ReadRequest{
			Filters: append(filters, filter),
		}
		recipeResponse, _ := u.Recipe.ReadRecipe(ctx, readRequest)
		if recipeResponse == nil {
			recipeID := strconv.Itoa(int(recipeResponse[0].ID))
			return errors.New("recipe id " + recipeID + "not found!")
		}

		//calculate recipe cost
		filters = nil
		filter = utils.AddFilter("id", recipeResponse[0].MaterialID, "equal")
		readRequest = &domain.ReadRequest{
			Filters: append(filters, filter),
		}
		material, err := u.Material.ReadMaterial(ctx, readRequest)
		if err != nil {
			return err
		}
		materialCost = utils.CalculateCost(material, recipeResponse)
		materialCosts = append(materialCosts, materialCost)
	}

	//calculate overall cost
	request.Cost = utils.SliceSum(materialCosts)

	//update related item-recipe
	filters = nil
	filter = utils.AddFilter("item_id", request.ID, "equal")
	itemRecipeReadRequest := &domain.ReadRequest{
		Filters: append(filters, filter),
	}
	//read old recipes for item and update them
	oldItemRecipes, err := u.ItemRecipe.ReadItemRecipe(ctx, itemRecipeReadRequest)
	if err != nil {
		return err
	}

	//replace old recipes with new recipes
	if len(oldItemRecipes) == len(request.Recipes) {
		for key := 0; key < len(request.Recipes); key++ {
			itemRecUpReq := &entity.UpdateItemRecipeRequest{
				UpdatedAt: time.Now(),
				ID:        oldItemRecipes[key].ID,
				RecipeID:  request.Recipes[key],
			}
			if err := u.ItemRecipe.UpdateItemRecipe(ctx, itemRecUpReq); err != nil {
				return err
			}
		}
	} else if len(oldItemRecipes) > len(request.Recipes) {
		//have to delete some item-recipe
		for key := 0; key < len(request.Recipes); key++ {
			itemRecUpReq := &entity.UpdateItemRecipeRequest{
				UpdatedAt: time.Now(),
				ID:        oldItemRecipes[key].ID,
				RecipeID:  request.Recipes[key],
			}
			if err := u.ItemRecipe.UpdateItemRecipe(ctx, itemRecUpReq); err != nil {
				return err
			}
		}
		for key := len(request.Recipes); key < len(oldItemRecipes); key++ {
			itemRecDeleteReq := &entity.DeleteItemRecipeRequest{
				ID:        oldItemRecipes[key].ID,
				DeletedAt: time.Now(),
			}
			if err := u.ItemRecipe.DeleteItemRecipe(ctx, itemRecDeleteReq); err != nil {
			}
			return err
		}
	} else {
		//if len(oldItemRecipes) < len(request.Recipes)
		//so have to create some item-recipe
		for key := 0; key < len(oldItemRecipes); key++ {
			itemRecUpReq := &entity.UpdateItemRecipeRequest{
				UpdatedAt: time.Now(),
				ID:        oldItemRecipes[key].ID,
				ItemID:    oldItemRecipes[key].ItemID,
				RecipeID:  request.Recipes[key],
			}
			if err := u.ItemRecipe.UpdateItemRecipe(ctx, itemRecUpReq); err != nil {
				return err
			}
		}
		for key := len(oldItemRecipes); key < len(request.Recipes); key++ {
			itemRecCreateReq := &entity.CreateItemRecipeRequest{
				ItemID:   request.ID,
				RecipeID: request.Recipes[key],
			}
			if err := u.ItemRecipe.CreateItemRecipe(ctx, itemRecCreateReq); err != nil {
				return err
			}
		}
	}

	if err := u.Item.UpdateItem(ctx, request); err != nil {
		return err
	}
	return nil
}
