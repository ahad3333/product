package postgres

import (
	"add/models"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type clientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *clientRepo {
	return &clientRepo{
		db: db,
	}
}

func (r *clientRepo) Create(ctx context.Context, req *models.CreateClient) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO client(id, name, updated_at)
			VALUES ($1, $2, now())
	`

	_, err := r.db.Exec(ctx, query, id, req.Name)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *clientRepo) GetByID(ctx context.Context, req *models.ClientPrimeryKey) (*models.Client, error) {
	var (
		Client    = &models.Client{}
		createdAt sql.NullString
		updatedAt sql.NullString
	)
	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM client
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&Client.Id,
		&Client.Name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return &models.Client{}, err
	}

	Client.CreatedAt = createdAt.String
	Client.UpdatedAt = updatedAt.String

	return Client, nil
}
