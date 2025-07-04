package products

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	productsService := ProductsService{}
	productsController := ProductsController{service: productsService}

	group.GET("/products", productsController.GetProducts)
	group.POST("/products", productsController.PostProduct)
	group.GET("/products/:id", productsController.GetProduct)
	group.PATCH("/products/:id", productsController.UpdateProduct)
	group.DELETE("/products/:id", productsController.DeleteProduct)
}
