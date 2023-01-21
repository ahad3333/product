package models

type ProductPrimeryKey struct {
	Id string `json:"id"`
}
type Product struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Price       float64        `json:"price"`
	Description string         `json:"description"`
	Photo       string         `json:"photo"`
	CategoryId  CategoryPrduct `json:"category_id"`
	BranchId    string         `json:"branch_id"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

type Productlist struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Photo       string  `json:"photo"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CreateProduct struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Photo       string  `json:"photo"`
	BranchId    string  `json:"branch_id"`
	CategoryId  string  `json:"category_id"`
}

type GetListProductRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int64          `json:"count"`
	Products []*Productlist `json:"books"`
}
type UpdateProduct struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Photo       string  `json:"photo"`
	CategoryId  string  `json:"category_id"`
	BranchId    string  `json:"branch_id"`
}
type UpdateProductPut struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Photo       string  `json:"photo"`
	CategoryId  string  `json:"category_id"`
}
