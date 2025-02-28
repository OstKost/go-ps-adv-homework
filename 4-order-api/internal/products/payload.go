package products

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=50"`
	Description string   `json:"description" validate:"max=500"`
	Images      []string `json:"images"`
}

type CreateProductResponse struct{}

type GetProductRequest struct {
	ProductId string `json:"productId" validate:"required"`
}

type GetProductResponse struct {
	Product Product
}

type GetProductsRequest struct {
	Products []string `json:"products" validate:"required"`
}

type GetProductsResponse struct {
	Products []Product
}

type UpdateProductRequest struct {
	Name        string   `json:"name" validate:"min=3,max=50"`
	Description string   `json:"description" validate:"max=500"`
	Images      []string `json:"images"`
}

type UpdateProductResponse struct {
	Product Product
}

type DeleteProductRequest struct {
	ProductId string `json:"productId" validate:"required"`
}

type DeleteProductResponse struct {
	Message string `json:"message"`
}
