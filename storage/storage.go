package storage

import (
	"add/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Branch() BranchRepoI
	User() UserRepoI
	Client() ClientRepoI
	Order() OrderRepoI
	Product() ProductRepoI
	Category() CategoryRepoI
}

type BranchRepoI interface {
	Create(ctx context.Context, req *models.CreateBranch) (string, error)
	GetByID(ctx context.Context, req *models.BranchPrimeryKey) (*models.Branch, error)
}

type UserRepoI interface {
	Insert(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimeryKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(ctx context.Context, praduct *models.UpdateUser) error
	Delete(ctx context.Context, req *models.UserPrimeryKey) error
	CheckUser(ctx context.Context, req *models.Login) (bool, error)
}

type ClientRepoI interface {
	Create(ctx context.Context, req *models.CreateClient) (string, error)
	GetByID(ctx context.Context, req *models.ClientPrimeryKey) (*models.Client, error)
}

type OrderRepoI interface {
	Create(ctx context.Context, req *models.CreateOrder) (string, error)
	CreateBucket(ctx context.Context, req *models.CreateBucket) (string, error)
	GetBucketByClientID(ctx context.Context, req *models.GetBucketByClientID) (*models.GetBucketByClientResponse, error)
	// Update(ctx context.Context, order *models.Order) error
}

type ProductRepoI interface {
	Insert(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimeryKey) (*models.Product, error)
	GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, praduct *models.UpdateProduct) error
	Delete(ctx context.Context, req *models.ProductPrimeryKey) error
}

type CategoryRepoI interface {
	Insert(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimeryKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(ctx context.Context, category *models.UpdateCategory) error
	Delete(ctx context.Context, req *models.CategoryPrimeryKey) error
}


