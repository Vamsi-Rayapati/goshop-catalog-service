package products

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/smartbot/catalog/database"
	"github.com/smartbot/catalog/pkg/dbclient"
	"github.com/smartbot/catalog/pkg/errors"
	"github.com/smartbot/catalog/pkg/utils"
	"gorm.io/gorm"
)

type ProductsService struct {
}

func (ps *ProductsService) AddProduct(req CreateProductRequest) (*ProductResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	item := database.Product{
		ID:          uuid.New(),
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
		CategoryID:  item.CategoryID,
		CreatedAt:   item.CreatedAt.String(),
		UpdatedAt:   item.UpdatedAt.String(),
	}, nil

}

func (ps *ProductsService) GetProducts(request GetProductsRequest) (*ProductsResponse, *errors.ApiError) {
	db := dbclient.GetCient()
	var items []database.Product
	var total int64

	db.Model(&database.Product{}).Count(&total)

	query := db.Model(&database.Product{}).Preload("Category")

	if request.CategoryID != 0 {
		query = query.Where("category_id = ?", request.CategoryID)
	}

	result := query.Order("created_at").Offset(request.PageSize * (request.PageNo - 1)).Limit(request.PageSize).Find(&items)

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
			CategoryID:  item.CategoryID,
			Category:    item.Category.Name,
			CreatedAt:   item.CreatedAt.String(),
			UpdatedAt:   item.UpdatedAt.String(),
		}
	})

	return &ProductsResponse{
		Products: itemList,
		Total:    total,
	}, nil

}

func (ps *ProductsService) GetProduct(id string) (*ProductResponse, *errors.ApiError) {
	db := dbclient.GetCient()
	var item database.Product
	var images []database.ProductImages
	result := db.Where("id = ?", id).First(&item)
	result2 := db.Model(database.ProductImages{}).Where("product_id = ?", id).Find(&images)

	if result.Error != nil || result2.Error != nil {
		log.Println("GetProduct: %+v", result.Error)
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NotFoundError("Product not found")
		}

		return nil, errors.InternalServerError("Failed to get product")
	}

	finalImages := utils.Map(images, func(img database.ProductImages) ImageResponse {
		return ImageResponse{
			ID:           img.ID.String(),
			IsPrimary:    img.IsPrimary,
			ImageUrl:     img.ImageURL,
			DisplayOrder: img.DisplayOrder,
		}
	})

	return &ProductResponse{
		ID:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		Stock:       item.Stock,
		Price:       item.Price,
		Images:      finalImages,
		CategoryID:  item.CategoryID,
		CreatedAt:   item.CreatedAt.String(),
		UpdatedAt:   item.UpdatedAt.String(),
	}, nil
}

func (ps *ProductsService) UpdateProduct(id string, request CreateProductRequest) (*ProductResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	var item database.Product
	var images []database.ProductImages

	data, err := utils.StructToMap(request)
	if err != nil {
		return nil, errors.InternalServerError("Failed to read payload")
	}

	result := db.Model(database.Product{}).Where("id = ?", id).Updates(data).First(&item)
	result2 := db.Model(database.ProductImages{}).Where("product_id = ?", id).Find(&images)
	if result.Error != nil || result2.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NotFoundError("Product not found")
		}

		return nil, errors.InternalServerError("Failed to update product")
	}

	finalImages := utils.Map(images, func(img database.ProductImages) ImageResponse {
		return ImageResponse{
			ID:           img.ID.String(),
			IsPrimary:    img.IsPrimary,
			ImageUrl:     img.ImageURL,
			DisplayOrder: img.DisplayOrder,
		}
	})

	return &ProductResponse{
		ID:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		Stock:       item.Stock,
		Price:       item.Price,
		CategoryID:  item.CategoryID,
		Images:      finalImages,
		CreatedAt:   item.CreatedAt.String(),
		UpdatedAt:   item.UpdatedAt.String(),
	}, nil

}

func (ps *ProductsService) DeleteProduct(productId string) *errors.ApiError {
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

func (ps *ProductsService) PostImages(productId string, req PostImagesRequest) (*PostImagesResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	tx := db.Begin()

	deleteResult := db.Where("product_id = ?", productId).Delete(&database.ProductImages{})

	if deleteResult.Error != nil {
		tx.Rollback()
		return nil, errors.InternalServerError("Failed to save images")
	}

	var newImages []database.ProductImages

	for _, img := range req.Images {
		newImages = append(newImages, database.ProductImages{
			ProductID:    productId,
			ImageURL:     img.ImageUrl,
			IsPrimary:    img.IsPrimary,
			DisplayOrder: img.DisplayOrder,
		})
	}

	result := tx.Create(&newImages)
	if result.Error != nil {
		tx.Rollback()
		return nil, errors.InternalServerError("Failed to save images")
	}

	tx.Commit()

	finalImages := utils.Map(newImages, func(img database.ProductImages) ImageResponse {
		return ImageResponse{
			ID:           img.ID.String(),
			DisplayOrder: img.DisplayOrder,
			IsPrimary:    img.IsPrimary,
			ImageUrl:     img.ImageURL,
		}
	})

	return &PostImagesResponse{
		Images: finalImages,
	}, nil

}
