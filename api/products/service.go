package products

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/smartbot/catalog/database"
	"github.com/smartbot/catalog/pkg/dbclient"
	"github.com/smartbot/catalog/pkg/errors"
	"github.com/smartbot/catalog/pkg/utils"
	"gorm.io/gorm"
)

type ProductsService struct {
}

func (us *ProductsService) AddProduct(req CreateProductRequest) (*ProductResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	item := database.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
	}
	result := db.Create(&item)

	if result.Error != nil {
		mysqlErr, ok := result.Error.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return nil, errors.ConfilctError("Productname already exists")
		}
		return nil, errors.InternalServerError("Failed to create product")
	}

	return &ProductResponse{
		ID:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		Stock:       item.Stock,
		Price:       item.Price,
		CreatedAt:   item.CreatedAt.String(),
		UpdatedAt:   item.UpdatedAt.String(),
	}, nil

}

func (us *ProductsService) GetProducts(request GetProductsRequest) (*ProductsResponse, *errors.ApiError) {
	db := dbclient.GetCient()
	var items []database.Product
	var total int64

	db.Model(&database.Product{}).Count(&total)
	result := db.Order("created_at").Offset(request.PageSize * (request.PageNo - 1)).Limit(request.PageSize).Find(&items)

	if result.Error != nil {
		return nil, errors.InternalServerError("Failed to get products")
	}

	itemList := utils.Map(items, func(item database.Product) ProductResponse {
		return ProductResponse{
			ID:          item.ID.String(),
			Name:        item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			CreatedAt:   item.CreatedAt.String(),
			UpdatedAt:   item.UpdatedAt.String(),
		}
	})

	return &ProductsResponse{
		Products: itemList,
		Total:    total,
	}, nil

}

func (us *ProductsService) GetProduct(id string) (*ProductResponse, *errors.ApiError) {
	db := dbclient.GetCient()
	var item database.Product
	result := db.Where("id = ?", id).First(&item)

	if result.Error != nil {
		log.Println("GetProduct: %+v", result.Error)
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NotFoundError("Product not found")
		}

		return nil, errors.InternalServerError("Failed to get product")
	}

	return &ProductResponse{
		ID:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		Stock:       item.Stock,
		Price:       item.Price,
		CreatedAt:   item.CreatedAt.String(),
		UpdatedAt:   item.UpdatedAt.String(),
	}, nil
}

func (us *ProductsService) UpdateProduct(id string, request CreateProductRequest) (*ProductResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	var item database.Product

	data, err := utils.StructToMap(request)
	if err != nil {
		return nil, errors.InternalServerError("Failed to read payload")
	}

	log.Println("UpdateProduct: %+v", data)
	result := db.Model(database.Product{}).Where("id = ?", id).Updates(data).First(&item)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NotFoundError("Product not found")
		}

		return nil, errors.InternalServerError("Failed to update product")
	}

	return &ProductResponse{
		ID:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		Stock:       item.Stock,
		Price:       item.Price,
		CreatedAt:   item.CreatedAt.String(),
		UpdatedAt:   item.UpdatedAt.String(),
	}, nil

}

func (us *ProductsService) DeleteProduct(productId string) *errors.ApiError {
	db := dbclient.GetCient()
	result := db.Where("id = ?", productId).Delete(&database.Product{})
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.NotFoundError("Product not found")
		}
		return errors.InternalServerError("Failed to get products")
	}

	return nil
}
