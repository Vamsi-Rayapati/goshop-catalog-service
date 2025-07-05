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
	Category    string  `json:"category"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int64             `json:"total"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	CategoryID  uint    `json:"category_id" validate:"required"`
}

type ImageRequest struct {
	IsPrimary    bool   `json:"is_primary"`
	ImageUrl     string `json:"image_url" validate:"required"`
	DisplayOrder int    `json:"display_order"  validate:"required,gte=1"`
}
type PostImagesRequest struct {
	Images []ImageRequest `json:"images"`
}

type ImageResponse struct {
	ID           string `json:"id"`
	IsPrimary    bool   `json:"is_primary"`
	ImageUrl     string `json:"image_url"`
	DisplayOrder int    `json:"display_order"`
}
type PostImagesResponse struct {
	Images []ImageResponse `json:"images"`
}
