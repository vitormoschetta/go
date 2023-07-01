package requests

type CreateProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type UpdateProductRequest struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
}

type ApplyPromotionProductRequest struct {
	ProductId  string  `json:"product_id"`
	Percentage float64 `json:"percentage"`
}

type ApplyPromotionProductByCategoryRequest struct {
	CategoryId string  `json:"category_id"`
	Percentage float64 `json:"percentage"`
}
