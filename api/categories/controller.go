package categories

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	apicommon "github.com/smartbot/catalog/pkg/api_common"
	"github.com/smartbot/catalog/pkg/validator"
)

type CategoriesController struct {
	service CategoriesService
}

func (cc CategoriesController) GetCategories(c *gin.Context) {
	var request CategoriesRequest

	log.Printf("Got request GetCategories2")
	err := validator.ValidateQueryParams(c, &request)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := cc.service.GetCategories(request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (cc CategoriesController) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	err := validator.ValidateBody(c, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := cc.service.CreateCategory(req)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, res)

}

func (cc CategoriesController) DeleteCategory(c *gin.Context) {
	Id := c.Param("id")

	err := cc.service.DeleteCategory(Id)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, &apicommon.ApiResponse{
		Code:    http.StatusAccepted,
		Message: "Category deleted successfully",
	})

}
