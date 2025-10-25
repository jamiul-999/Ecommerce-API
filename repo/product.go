package repo

import (
	"database/sql"

	"ecommerce-api/domain"
	"ecommerce-api/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (
			title,
			description,
			price,
			img_url
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil

}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product

	query := `
		SELECT
			id,
			title,
			description,
			price,
			img_url
		from products
		where id = $1
	`
	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &prd, nil

}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {

	offset := ((page - 1) * limit) + 1
	var prdList []*domain.Product

	query := `
		SELECT
			id,
			title,
			description,
			price,
			img_url
		from products
		LIMIT $1
		OFFSET $2;
		`

	err := r.db.Select(&prdList, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return prdList, nil
}

func (r *productRepo) Count() (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM products
	`
	var count int
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET title=$1, description=$2, price=$3, img_url=$4
		WHERE id=$5
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL, p.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &p, nil
}
