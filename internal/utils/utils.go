package utils

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"log"
	"os"
	"price/internal/domain"
	"price/internal/entity"
	"strconv"
)

func GetFromSheet() (*spreadsheet.Sheet, error) {
	//get spreedsheet id and sheet id from env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	spreedSheetID := os.Getenv("SpreedSheetID")
	sheetIDStr := os.Getenv("SheetID")
	sheetID, err := strconv.Atoi(sheetIDStr)

	data, err := os.ReadFile("client_secret.json")
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return nil, err
	}
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(spreedSheetID)
	if err != nil {
		return nil, err
	}
	sheet, err := spreadsheet.SheetByIndex(uint(sheetID))
	if err != nil {
		return nil, err
	}
	if len(sheet.Columns) != 4 {
		return nil, errors.New("sheet is in wrong format")
	}
	return sheet, nil
}

func AddFilter(field string, value interface{}, tp string) domain.Filter {
	var filter domain.Filter

	switch tp {
	case "like":
		filter = domain.Filter{
			Field: field + " LIKE ?",
			Value: "%" + fmt.Sprint(value) + "%",
		}
	case "equal":
		filter = domain.Filter{
			Field: field + " = ?",
			Value: value,
		}
	}
	return filter
}

func ConvertMaterial(org []spreadsheet.Cell) (*entity.CreateMaterialRequest, error) {
	unit, err := strconv.ParseFloat(org[1].Value, 64)
	if err != nil {
		return &entity.CreateMaterialRequest{}, err
	}

	price, err := strconv.ParseFloat(org[2].Value, 64)
	if err != nil {
		return &entity.CreateMaterialRequest{}, err
	}

	if org[3].Value == "" {
		org[3].Value = "1"
	}
	efficiency, err := strconv.ParseFloat(org[3].Value, 64)
	if err != nil {
		return &entity.CreateMaterialRequest{}, err
	}

	converted := &entity.CreateMaterialRequest{
		Name:       org[0].Value,
		Unit:       unit,
		Price:      price,
		RealPrice:  price / efficiency,
		Efficiency: efficiency,
	}
	return converted, nil
}

func ConvertRequest(org *entity.UsualRequest) (*domain.ReadRequest, error) {
	var err error
	var filters []domain.Filter
	var pageInfo domain.PageInt
	if org.ID.Str != "" {
		id, err := strconv.Atoi(org.ID.Str)
		if err != nil {
			return &domain.ReadRequest{}, err
		}
		filter := AddFilter("id", id, "equal")
		filters = append(filters, filter)
	}
	if org.Name != "" {
		filter := AddFilter("name", org.Name, "equal")
		filters = append(filters, filter)
	}

	pageInfo.PerPage, err = strconv.Atoi(org.PageRequest.PerPage)
	if err != nil {
		return &domain.ReadRequest{}, err
	}
	pageInfo.PageNumber, err = strconv.Atoi(org.PageRequest.PageNumber)
	if err != nil {
		return &domain.ReadRequest{}, err
	}

	response := &domain.ReadRequest{
		Filters:  filters,
		PageInfo: pageInfo,
	}

	return response, nil
}

func SliceSum(slice []float64) float64 {
	var response float64
	for _, cost := range slice {
		response += cost
	}
	return response
}

func CalculateCost(material []entity.Material, recipe []entity.Recipe) float64 {
	materialCostPerUnit := material[0].RealPrice / material[0].Unit
	response := materialCostPerUnit * recipe[0].Volume
	return response
}
