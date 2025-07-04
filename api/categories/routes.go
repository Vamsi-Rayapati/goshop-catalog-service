package categories

import "github.com/gin-gonic/gin"

func RegisterRoutes(group *gin.RouterGroup) {
	service := CategoriesService{}
	controller := CategoriesController{service: service}
	group.GET("/categories", controller.GetCategories)
	group.POST("/categories", controller.CreateCategory)
	group.DELETE("/categories/:id", controller.DeleteCategory)
}
