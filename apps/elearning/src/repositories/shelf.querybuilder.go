package repositories

import (
	"fmt"
	"strings"

	"backend/apps/elearning/src/structs"
)

type ShelfQueryBuilder struct{}

func NewShelfQueryBuilder() ShelfQueryBuilder {
	return ShelfQueryBuilder{}
}

func (this ShelfQueryBuilder) Create() string {
	return `INSERT INTO shelves (title)
    VALUES (:title)
    RETURNING id`
}

func (this ShelfQueryBuilder) Update() string {
	return `UPDATE shelves SET title = :title, updated_at = NOW() WHERE id = :id RETURNING id`
}

func (this ShelfQueryBuilder) ReadOne() string {
	return `SELECT sh.* FROM shelves sh WHERE sh.id = $1 LIMIT 1`
}

func (this ShelfQueryBuilder) ReadAll(search structs.Search) string {
	query := `SELECT sh.* FROM shelves sh`
	var where string
	search.TablePrefix = "sh."
	if search.Where() != "" {
		where = fmt.Sprintf("WHERE %s", search.Where())
	}
	query = fmt.Sprintf("%s %s", query, where)

	pagination := search.Pagination
	if pagination.Order != "" && pagination.SortBy != "" {
		query = fmt.Sprintf(`%s ORDER BY %s %s`, query, pagination.Order, pagination.SortBy)
	}
	if pagination.PerPage > 0 && pagination.Page > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, pagination.PerPage, pagination.Offset())
	}
	return query
}

func (this ShelfQueryBuilder) AssociateShelfAndAlbums(shelf structs.Shelf) string {
	var values []string
	for index, item := range shelf.Albums {
		values = append(values, fmt.Sprintf("($1, %d, %d)", item.ID, index))
	}
	return fmt.Sprintf(`INSERT INTO albums_shelves (shelf_id, album_id, album_position) VALUES %s`, strings.Join(values, ","))
}

func (this ShelfQueryBuilder) DisassociateAlbumsFromShelf() string {
	return `DELETE FROM albums_shelves WHERE shelf_id = $1`
}

func (this ShelfQueryBuilder) Delete() string {
	return `DELETE FROM shelves WHERE id = $1 RETURNING id`
}

func (this ShelfQueryBuilder) FindShelvesByAlbumID() string {
	return `SELECT s.* FROM shelves s, albums_shelves ash WHERE s.id = ash.shelf_id AND ash.album_id = $1`
}

func (this ShelfQueryBuilder) FindHomeShelves(search structs.Search) string {
	query := `select
		       s.id, s.title, s.created_at, s.updated_at,
		       (select count(distinct s.id) from shelves s
			inner join albums_shelves ash on s.id = ash.shelf_id where s.id IN(%s)) as total
		from shelves s
		inner join albums_shelves ash on s.id = ash.shelf_id
        inner join albums a on a.is_published = true and ash.album_id = a.id
		where s.id IN(%s)
		group by s.id
		order by array_position(array[%s], s.id)`
	query = fmt.Sprintf(query, search.Raw, search.Raw, search.Raw)

	pagination := search.Pagination
	if pagination.PerPage > 0 && pagination.Page > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, pagination.PerPage, pagination.Offset())
	}
	return query
}

func (this ShelfQueryBuilder) Total(search structs.Search) string {
	query := `SELECT COUNT(sh.id) FROM shelves sh`
	var where string
	if search.Where() != "" {
		where = fmt.Sprintf("WHERE %s", search.Where())
	}
	query = fmt.Sprintf("%s %s", query, where)
	return query
}
