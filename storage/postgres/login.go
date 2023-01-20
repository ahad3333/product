package postgres

import (
	"add/models"
	// "database/sql"
	"context"
	// "fmt"

	// "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type LoginRepo struct {
	db *pgxpool.Pool
}

func NewLoginRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *LoginRepo) Insert(ctx context.Context, req *models.Login) (models.Login, error) {
 res:=models.Login{
	 Login: req.Login,
	Password: req.Password,
 }
	return res, nil
}