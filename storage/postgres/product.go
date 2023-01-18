package postgres

import (
	"add/models"
	"database/sql"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Insert(ctx context.Context, req *models.CreateProduct) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
	INSERT INTO product (
		id,
		name,
		price,
		description,
		photo,
		category_id,
		updated_at
	) VALUES ($1, $2, $3, $4,$5,$6, now())
`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Price,
		req.Description,
		req.Photo,
		req.CategoryId,
	)

	if err != nil {
		return "", err
	}
	

	return id, nil
}

func (r *ProductRepo) GetByID(ctx context.Context, req *models.ProductPrimeryKey) (*models.Product, error) {

	query := `select 
			id,
			name,
			price,
			description,
			photo,
			category_id,
			updated_at
	from product 
	where id = $1
	`

	queryCategory := `select 
			c.id,
			c.name,
			c.photo,
	from product as p
	join category as c on c.id = p.category_id
	where p.id = $1
	`

	var (
		id          sql.NullString
		name        sql.NullString
		price       sql.NullFloat64
		description sql.NullString
		photo    	sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).
		Scan(
			&id,
			&name,
			&price,
			&description,
			&photo,
			&createdAt,
			&updatedAt,
		)
	if err != nil {
		return nil, err
	}

	praduct := &models.Product{
		Id:          id.String,
		Name:        name.String,
		Price:       price.Float64,
		Description: description.String,
		Photo:        photo.String,	
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}
	var categorys models.CategoryPrduct
	err = r.db.QueryRow(ctx, queryCategory, req.Id).Scan(
		&categorys.Id,
		&categorys.Name,
		&categorys.Photo,
	)



praduct.CategoryId = categorys		

	return praduct, nil
}

func (r *ProductRepo)GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error) {

	var (
		resp   models.GetListProductResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
		f ="%"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			price,
			description,
			photo,
			created_at,
			updated_at 
		FROM product
	`
	if search !="" {
		search = fmt.Sprintf("where name like  '%s%s' ", req.Search,f)
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
		return &models.GetListProductResponse{}, err
	}
	var (
		id          sql.NullString
		name        sql.NullString
		price       sql.NullFloat64
		description sql.NullString
		photo       sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	for rows.Next() {


		err = rows.Scan(
			&resp.Count,
			&id,
			&name,
			&price,
			&description,
			&photo,
			&createdAt,
			&updatedAt,
		)
		praduct := models.Productlist{
			Id:          id.String,
			Name:        name.String,
			Price:       price.Float64,
			Description: description.String,
			Photo:       photo.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}
		if err != nil {
			return &models.GetListProductResponse{}, err
		}
		
		resp.Products = append(resp.Products, &praduct)


	}
	return &resp, nil
}

func (r *ProductRepo)Update(ctx context.Context, praduct *models.UpdateProduct) error {
	query := `
		UPDATE 
			product
		SET 
			name = $2,
			price = $3,
			description = $4,
			photo = $5,
			category_id = $6
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx,query,
		praduct.Id,
		praduct.Name,
		praduct.Price,
		praduct.Description,
		praduct.Photo,
		praduct.CategoryId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepo)Delete(ctx context.Context, req *models.ProductPrimeryKey) error {
	
	_, err := r.db.Exec(ctx, "DELETE FROM product WHERE id = $1", req.Id)

	if err != nil {
		return err
	}

	return nil
}
