package postgres

import (
	"add/models"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (r *branchRepo) Create(ctx context.Context, req *models.CreateBranch) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO branch(id, name, updated_at)
			VALUES ($1, $2, now())
	`

	_, err := r.db.Exec(ctx, query, id, req.Name)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *branchRepo) GetByID(ctx context.Context, req *models.BranchPrimeryKey) (*models.Branch, error) {
	var (
		branch    = &models.Branch{}
		createdAt sql.NullString
		updatedAt sql.NullString
	)
	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM branch
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&branch.Id,
		&branch.Name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return &models.Branch{}, err
	}

	branch.CreatedAt = createdAt.String
	branch.UpdatedAt = updatedAt.String

	return branch, nil
}
