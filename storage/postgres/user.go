package postgres

import (
	"add/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Insert(ctx context.Context, req *models.CreateUser) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO "user" (
		id,
		name,
		login,
		password,
		branch_id,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, now())
`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Login,
		req.Password,
		req.BranchId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepo) GetByID(ctx context.Context, req *models.UserPrimeryKey) (*models.User, error) {

	query := `select 
			id,
			name,
			login,
			password,
			created_at,
			updated_at
	from "user" 
	where id = $1
	`

	var (
		id        sql.NullString
		name      sql.NullString
		login     sql.NullString
		password  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&name,
			&login,
			&password,
			&createdAt,
			&updatedAt,
		)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Id:        id.String,
		Name:      name.String,
		Login:     login.String,
		Password:  password.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return user, nil
}

func (r *UserRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	var (
		resp   models.GetListUserResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
		f      = "%"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			login,
			password,
			created_at,
			updated_at 
		FROM "user"
	`
	if search != "" {
		search = fmt.Sprintf("where name like  '%s%s' ", req.Search, f)
		query += search
	}
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit
	rows, err := r.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return &models.GetListUserResponse{}, err
	}
	var (
		id        sql.NullString
		name      sql.NullString
		login     sql.NullString
		password  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	for rows.Next() {

		err = rows.Scan(
			&resp.Count,
			&id,
			&name,
			&login,
			&password,
			&createdAt,
			&updatedAt,
		)
		user := models.User{
			Id:        id.String,
			Name:      name.String,
			Login:     login.String,
			Password:  password.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		}
		if err != nil {
			return &models.GetListUserResponse{}, err
		}

		resp.Users = append(resp.Users, &user)

	}
	return &resp, nil
}

func (r *UserRepo) Update(ctx context.Context, user *models.UpdateUser) error {
	query := `
		UPDATE 
			"user"
		SET 
			name = $2,
			login =$3,
			password = $4,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		user.Id,
		user.Name,
		user.Login,
		user.Password,
		// user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Delete(ctx context.Context, req *models.UserPrimeryKey) error {

	_, err := r.db.Exec(ctx, `DELETE FROM "user" WHERE id = $1`, req.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) CheckUser(ctx context.Context, req *models.Login) (bool, error) {

	var count int32

	query := `
		SELECT count(*) FROM "user"
		WHERE login = $1 AND password = $2
	`

	err := r.db.QueryRow(ctx, query, req.Login, req.Password).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
