package models

type ProductResponse struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Price    float64       `json:"price"`
	Image    ProductImages `json:"image"`
}

func MapProductToResponse(product Product) ProductResponse {
	return ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
		Image:    product.Images,
	}
}

// MapProductsToResponse maps a slice of Product models to a slice of ProductResponse
func MapProductsToResponse(products []Product) []ProductResponse {
	responses := make([]ProductResponse, len(products))
	for i, product := range products {
		responses[i] = MapProductToResponse(product)
	}
	return responses
}
