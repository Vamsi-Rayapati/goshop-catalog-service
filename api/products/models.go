package products

type GetProductsRequest struct {
	PageNo   int `form:"page_no" validate:"required,gte=1"`
	PageSize int `form:"page_size" validate:"required,gte=1"`
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  uint    `json:"category_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int64             `json:"total"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}
