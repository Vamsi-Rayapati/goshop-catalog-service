package categories

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type CategoriesRequest struct {
	PageNo   int `form:"page_no" validate:"required,gte=1"`
	PageSize int `form:"page_size" validate:"required,gte=1"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
	Total      int64              `json:"total"`
}
