package categories

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/smartbot/catalog/database"
	"github.com/smartbot/catalog/pkg/client"
	"github.com/smartbot/catalog/pkg/dbclient"
	"github.com/smartbot/catalog/pkg/errors"
	"github.com/smartbot/catalog/pkg/utils"
	"gorm.io/gorm"
)

type CategoriesService struct {
}

var ctx = context.Background()

func (cs *CategoriesService) GetCategories(request CategoriesRequest) (*CategoriesResponse, *errors.ApiError) {

	redisClient := client.GetRedisClient()
	cached, err := redisClient.Get(ctx, "categories").Result()

	if err == nil {
		log.Println("Got Cache", cached)
		var resp []CategoryResponse
		if unmarshalErr := json.Unmarshal([]byte(cached), &resp); unmarshalErr == nil {
			return &CategoriesResponse{
				Categories: resp,
				Total:      int64(len(resp)),
			}, nil
		}
	}

	db := dbclient.GetCient()
	var categories []database.Category
	var total int64
	db.Model(&database.Category{}).Count(&total)
	result := db.Order("created_at").Offset(request.PageSize * (request.PageNo - 1)).Limit(request.PageSize).Find(&categories)

	if result.Error != nil {
		return nil, errors.InternalServerError("Failed to get categories")
	}

	catList := utils.Map(categories, func(user database.Category) CategoryResponse {
		return CategoryResponse{
			ID:   user.ID,
			Name: user.Name,
		}
	})

	serialized, _ := json.Marshal(catList)
	redisClient.SetEx(ctx, "categories", serialized, 10*time.Minute)

	return &CategoriesResponse{
		Categories: catList,
		Total:      total,
	}, nil

}

func (cs *CategoriesService) CreateCategory(req CreateCategoryRequest) (*CategoryResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	category := database.Category{
		Name: req.Name,
	}
	result := db.Create(&category)

	if result.Error != nil {
		return nil, errors.InternalServerError("Failed to create category")
	}

	return &CategoryResponse{
		Name: category.Name,
		ID:   category.ID,
	}, nil

}

func (us *CategoriesService) DeleteCategory(Id string) *errors.ApiError {
	db := dbclient.GetCient()
	result := db.Where("id = ?", Id).Delete(&database.Category{})
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.NotFoundError("Category not found")
		}
		return errors.InternalServerError("Failed to get category")
	}

	return nil
}
