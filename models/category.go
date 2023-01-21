package models

type Category struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	Photo     string            `json:"photo"`
	Type      []*CategoryPrduct `json:"type"`
	BranchId  string            `json:"branch_id"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}
type CategoryPrduct struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	CategoryId string `json:"type_id"`
	BranchId   string `json:"branch_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
type CategoryPrimeryKey struct {
	Id string `json:"id"`
}

type CreateCategory struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	CategoryId string `json:"type_id"`
	BranchId   string `json:"branch_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateCategory struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	CategoryId string `json:"category_id"`
	BranchId   string `json:"branch_id"`
}
type UpdateCategoryPut struct {
	Name       string `json:"name"`
	Photo      string `json:"photo"`
	CategoryId string `json:"category_id"`
	BranchId   string `json:"branch_id"`
}
type GetListCategoryRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListCategoryResponse struct {
	Count      int64             `json:"count"`
	Categories []*CreateCategory `json:"categories"`
}

type Empty struct{}
