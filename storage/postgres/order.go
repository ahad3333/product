package postgres

import (
	"add/models"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO "order" (
				id,
				total_price,
				branch_id,
				client_id,
				updated_at
			)
			VALUES ($1, $2, $3, $4, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.TotalPrice,
		req.BranchId,
		req.ClientId,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *orderRepo) CreateBucket(ctx context.Context, req *models.CreateBucket) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO bucket (
				id,
				branch_id,
				client_id,
				product_id,
				updated_at
			)
			VALUES ($1, $2, $3, $4, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.BranchId,
		req.ClientId,
		req.ProductId,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *orderRepo) GetBucketByClientID(ctx context.Context, req *models.GetBucketByClientID) (*models.GetBucketByClientResponse, error) {

	var (
		resp models.GetBucketByClientResponse
		created_at sql.NullString
	)

	query := `
		SELECT
			SUM(p.price)
		FROM bucket AS b
		LEFT JOIN product AS p ON p.id = b.product_id
		WHERE b.client_id = $1 AND b.branch_id = $2
	`

	err := r.db.QueryRow(ctx, query,
		req.ClientId,
		req.BranchId,
	).Scan(
		&resp.TotalPrice,
		&created_at,
	)
	resp.CreatedAt=created_at.String
	if err != nil {
		return nil, err
	}

	return &resp, nil
}



