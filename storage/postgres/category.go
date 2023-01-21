package postgres

import (
	"add/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}

}

func (r *CategoryRepo) Insert(ctx context.Context, req *models.CreateCategory) (string, error) {
	var (
		id = uuid.New().String()
	)
	if req.CategoryId == "" {
		fmt.Println(req.CategoryId)

		query := `
			INSERT INTO category (
				id,
				name,
				photo,
				branch_id,
				updated_at
			) VALUES ($1, $2, $3, $4, now())
		`

		_, err := r.db.Exec(ctx, query,
			id,
			req.Name,
			req.Photo,
			req.BranchId,
		)

		if err != nil {
			return "", err
		}
		return id, nil
	}

	if req.CategoryId != "" {

		query := `
			INSERT INTO category (
				id,
				name,
				photo,
				type_id,
				branch_id,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, now())
		`

		_, err := r.db.Exec(ctx, query,
			id,
			req.Name,
			req.Photo,
			req.CategoryId,
			req.BranchId,
		)

		if err != nil {
			return "", err
		}
		return id, nil

	}
	return "", nil

}

func (r *CategoryRepo) GetByID(ctx context.Context, req *models.CategoryPrimeryKey) (*models.Category, error) {
	query := `
	select 
	c.id,
	c.name,
	c.photo,
	c.created_at,
	c.updated_at,
    (
				SELECT
					ARRAY_AGG(id)
				FROM category AS ca 
				WHERE ca.type_id = $1
			) AS category_ids
from category as c
where c.id = $1
`

	var (
		id        sql.NullString
		name      sql.NullString
		photo     sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
		categorys []string
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&name,
			&photo,
			&createdAt,
			&updatedAt,
			(*pq.StringArray)(&categorys),
		)
	if err != nil {
		return nil, err
	}

	category := &models.Category{
		Id:        id.String,
		Name:      name.String,
		Photo:     photo.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	if len(categorys) > 0 {
		typeQuery := `
			SELECT
				id,
				name,
				photo,
				type_id,
				created_at,
				updated_at
			FROM
				category
			WHERE id IN (`

		for _, catego := range categorys {
			typeQuery += fmt.Sprintf("'%s',", catego)
		}
		typeQuery = typeQuery[:len(typeQuery)-1]
		typeQuery += ")"
		rows, err := r.db.Query(ctx, typeQuery)
		if err != nil {
			return nil, err
		}

		defer rows.Close()
		for rows.Next() {

			var (
				id        sql.NullString
				name      sql.NullString
				photo     sql.NullString
				type_id   sql.NullString
				createdAt sql.NullString
				updatedAt sql.NullString
			)

			err = rows.Scan(
				&id,
				&name,
				&photo,
				&type_id,
				&createdAt,
				&updatedAt,
			)
			types := &models.CategoryPrduct{
				Id:         id.String,
				Name:       name.String,
				Photo:      photo.String,
				CategoryId: type_id.String,
				CreatedAt:  createdAt.String,
				UpdatedAt:  updatedAt.String,
			}
			category.Type = append(category.Type, types)
		}
	}
	return category, nil
}

func (r *CategoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

	var (
		resp   models.GetListCategoryResponse
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
			photo,
			type_id,
			created_at,
			updated_at
		FROM category
		where  type_id is null
	`
	if search != "" {
		search = fmt.Sprintf("and name like  '%s%s' ", req.Search, f)
		query += search
	}
	if req.Offset > 0 {
		offset = fmt.Sprintf(" \n OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	query += offset + limit
	rows, err := r.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return &models.GetListCategoryResponse{}, err
	}
	var (
		id        sql.NullString
		name      sql.NullString
		photo     sql.NullString
		type_id   sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&resp.Count,
			&id,
			&name,
			&photo,
			&type_id,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return &models.GetListCategoryResponse{}, err
		}
		category := models.CreateCategory{
			Id:         id.String,
			Name:       name.String,
			Photo:      photo.String,
			CategoryId: type_id.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		}
		resp.Categories = append(resp.Categories, &category)

	}

	return &resp, nil
}

func (r *CategoryRepo) Update(ctx context.Context, Category *models.UpdateCategory) error {

	if Category.CategoryId != "" {
		query := `
		UPDATE 
		Category
		SET 
		name = $2,
		photo = $3,
		category_id =$4
		WHERE id = $1
	`

		_, err := r.db.Exec(ctx, query,
			Category.Id,
			Category.Name,
			Category.Photo,
			Category.CategoryId,
		)
		if err != nil {
			return err
		}
	} else {
		query := `
		UPDATE 
		Category
		SET 
		name = $2,
		photo = $3
		WHERE id = $1
	`

		_, err := r.db.Exec(ctx, query,
			Category.Id,
			Category.Name,
			Category.Photo,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *CategoryRepo) Delete(ctx context.Context, req *models.CategoryPrimeryKey) error {
	_, err := r.db.Exec(ctx, "DELETE FROM product WHERE category_id  = $1 ", req.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, "DELETE FROM category WHERE id = $1 ", req.Id)

	if err != nil {
		return err
	}

	return nil
}
