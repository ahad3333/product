package models

type CreateOrder struct {
	TotalPrice int64  `json:"total_price"`
	BranchId   string `json:"branch_id"`
	ClientId   string `json:"client_id"`
}

type OrderPrimeryKey struct {
	Id string `json:"id"`
}

type Order struct {
	Id         string `json:"id"`
	TotalPrice int64  `json:"total_price"`
	BranchId   string `json:"branch_id"`
	ClientId   string `json:"client_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateBucket struct {
	BranchId  string `json:"branch_id"`
	ClientId  string `json:"client_id"`
	ProductId string `json:"product_id"`
}

type BucketPrimeryKey struct {
	Id string `json:"id"`
}

type Bucket struct {
	Id        string   `json:"id"`
	BranchId  string   `json:"branch_id"`
	ClientId  string   `json:"client_id"`
	ProductId string   `json:"product_id"`
	Product   *Product `json:"product"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type GetBucketByClientID struct {
	ClientId string `json:"client_id"`
	BranchId string `json:"branch_id"`
}

type GetBucketByClientResponse struct {
	TotalPrice int64 `json:"total_price"`
	CreatedAt string   `json:"created_at"`
}
