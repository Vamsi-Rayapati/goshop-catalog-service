package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apicommon "github.com/smartbot/catalog/pkg/api_common"
	"github.com/smartbot/catalog/pkg/validator"
)

type ProductsController struct {
	service ProductsService
}

func (pc *ProductsController) GetCurrentProduct(c *gin.Context) {

	value, _ := c.Get("product_id")

	res, err := pc.service.GetProduct(value.(string))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (pc *ProductsController) GetProducts(c *gin.Context) {
	var request GetProductsRequest

	err := validator.ValidateQueryParams(c, &request)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	products, err := pc.service.GetProducts(request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, products)

}

func (pc *ProductsController) GetProduct(c *gin.Context) {
	productId := c.Param("id")
	err := validator.ValidateUUID(productId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := pc.service.GetProduct(productId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (pc *ProductsController) PostProduct(c *gin.Context) {
	var postRequest CreateProductRequest
	err := validator.ValidateBody(c, &postRequest)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := pc.service.AddProduct(postRequest)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, res)

}

func (pc *ProductsController) DeleteProduct(c *gin.Context) {
	productId := c.Param("id")
	err := validator.ValidateUUID(productId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	err = pc.service.DeleteProduct(productId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, &apicommon.ApiResponse{
		Code:    http.StatusAccepted,
		Message: "Product deleted successfully",
	})

}

func (pc *ProductsController) UpdateProduct(c *gin.Context) {
	var request CreateProductRequest
	productId := c.Param("id")
	err := validator.ValidateUUID(productId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	err = validator.ValidateBody(c, &request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := pc.service.UpdateProduct(productId, request)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, res)

}

func (pc *ProductsController) PostImages(c *gin.Context) {
	var req PostImagesRequest

	productId := c.Param("id")
	err := validator.ValidateUUID(productId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	err = validator.ValidateBody(c, &req)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := pc.service.PostImages(productId, req)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, res)

}

func (pc *ProductsController) GetImages(c *gin.Context) {

	productId := c.Param("id")
	err := validator.ValidateUUID(productId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	res, err := pc.service.GetImages(productId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
