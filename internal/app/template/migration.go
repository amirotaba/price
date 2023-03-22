package template

import (
	"context"
	"log"
	"price/internal/app"
	"price/internal/entity"
)

func CreateItem() {
	db := app.Connection()

	_, err := db.NewCreateTable().
		Model((*entity.Item)(nil)).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func CreateMaterial() {
	db := app.Connection()

	_, err := db.NewCreateTable().
		Model((*entity.Material)(nil)).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func CreateRecipe() {
	db := app.Connection()

	_, err := db.NewCreateTable().
		Model((*entity.Recipe)(nil)).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func CreateItemRecipe() {
	db := app.Connection()

	_, err := db.NewCreateTable().
		Model((*entity.ItemRecipe)(nil)).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func DropItem() {
	db := app.Connection()

	_, err := db.NewDropTable().
		Model((*entity.Item)(nil)).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func DropMaterial() {
	db := app.Connection()

	_, err := db.NewDropTable().
		Model((*entity.Material)(nil)).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func DropRecipe() {
	db := app.Connection()

	_, err := db.NewDropTable().
		Model((*entity.Recipe)(nil)).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func DropItemRecipe() {
	db := app.Connection()

	_, err := db.NewDropTable().
		Model((*entity.ItemRecipe)(nil)).
		Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}
