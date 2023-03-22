package domain

type Repositories struct {
	Item       ItemRepository
	Material   MaterialRepository
	Recipe     RecipeRepository
	ItemRecipe ItemRecipeRepository
}

type Usecases struct {
	Item     ItemUsecase
	Material MaterialUsecase
	Recipe   RecipeUsecase
}

type Services struct {
	Item     ItemUsecase
	Material MaterialUsecase
	Recipe   RecipeUsecase
	Port     string
}

type Filter struct {
	Field string
	Value interface{}
	Sub   string
}

type PageStr struct {
	PageNumber string
	PerPage    string
}

type PageInt struct {
	PageNumber int
	PerPage    int
}

type PageResponse struct {
}

type ReadRequest struct {
	Filters  []Filter
	PageInfo PageInt
}
