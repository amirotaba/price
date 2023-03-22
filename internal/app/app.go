package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"log"
	"os"
	"price/internal/app/http/echo"
	"price/internal/domain"
	"price/internal/features/item/repository/mysql"
	"price/internal/features/item/usecase"
	itemRecipeRepo "price/internal/features/itemRecipe/repository/mysql"
	materialRepo "price/internal/features/material/repository/mysql"
	materialUsecase "price/internal/features/material/usecase"
	"price/internal/features/recipe/repository/mysql"
	"price/internal/features/recipe/usecase"
)

func Connection() *bun.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("User")
	dbPassword := os.Getenv("Password")
	dbName := os.Getenv("Name")
	dsn := dbUser + ":" + dbPassword + "@/" + dbName

	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	return db
}

func NewRepository() domain.Repositories {
	d := Connection()
	repos := domain.Repositories{
		Item:       itemRepo.NewMysqlRepository(d),
		Recipe:     recipeRepo.NewMysqlRepository(d),
		Material:   materialRepo.NewMysqlRepository(d),
		ItemRecipe: itemRecipeRepo.NewMysqlRepository(d),
	}
	return repos
}

func NewUseCase() domain.Usecases {
	r := NewRepository()
	usecases := domain.Usecases{
		Item:     itemUsecase.NewUseCase(r),
		Recipe:   recipeUsecase.NewUseCase(r),
		Material: materialUsecase.NewUseCase(r),
	}
	return usecases
}

func StartServer() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	u := NewUseCase()

	services := domain.Services{
		Item:     u.Item,
		Recipe:   u.Recipe,
		Material: u.Material,

		Port: ":" + os.Getenv("Port"),
	}

	echo.New(services)
}
