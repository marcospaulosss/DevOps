package repositories

import (
	"backend/apps/ecommerce/src/structs"
	"fmt"
)

type ProductQueryBuilder struct{}

func NewProductQueryBuilder() ProductQueryBuilder {
	return ProductQueryBuilder{}
}

func (this ProductQueryBuilder) Create() string {
	return `INSERT INTO products (
        started_selling_at,
        finished_selling_at,
        usage_expires_at,
        name,
        description,
        stock,
        image,
        is_published,
        company,
        product_type,
        items,
        payments_types
    )
    VALUES (
        :started_selling_at,
        :finished_selling_at,
        :usage_expires_at,
        :name,
        :description,
        :stock,
        :image,
        :is_published,
        :company,
        :product_type,
        CAST(:items AS jsonb),
        CAST(:payments_types AS jsonb)
    )
    RETURNING *`
}

// VALUES (:name, :description, :stock, CAST(:tags AS jsonb))

func (this ProductQueryBuilder) Update() string {
	return `WITH CTE AS (
    SELECT row_to_json(t) FROM (
        SELECT updated_at, started_selling_at, finished_selling_at, usage_expires_at, name, description, stock, sku, company, image, product_type, payments_types, items FROM products WHERE id = :id) t
    )
    UPDATE products
    SET name = :name, description = :description, stock = :stock,
        history = coalesce(history, CAST('[]' AS jsonb)) || CAST((SELECT * FROM CTE) AS jsonb)
    WHERE id = :id
    RETURNING *`
}

func (this ProductQueryBuilder) ReadOne() string {
	return `SELECT * FROM products WHERE id=$1`
}

func (this ProductQueryBuilder) ReadAll(search structs.Search) string {
	query := `SELECT * FROM products`
	pagination := search.Pagination
	if pagination.Order != "" && pagination.SortBy != "" {
		query = fmt.Sprintf(`%s ORDER BY %s %s`, query, pagination.Order, pagination.SortBy)
	}
	if pagination.PerPage > 0 && pagination.Page > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, pagination.PerPage, pagination.Offset())
	}
	return query
}

func (this ProductQueryBuilder) Total() string {
	return `SELECT COUNT(id) FROM products LIMIT 1`
}
