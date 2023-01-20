package storage

import (
	"context"
	"add/models"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
	Category() CategoryRepoI
	User() UserRepoI
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

type UserRepoI interface {
	Insert(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimeryKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(ctx context.Context, praduct *models.UpdateUser) error
	Delete(ctx context.Context, req *models.UserPrimeryKey) error 
}