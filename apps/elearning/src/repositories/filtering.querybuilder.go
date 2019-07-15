package repositories

import (
	"fmt"

	"backend/apps/elearning/src/structs"
)

type FilteringQueryBuilder struct{}

func NewFilteringQueryBuilder() FilteringQueryBuilder {
	return FilteringQueryBuilder{}
}

func (s FilteringQueryBuilder) QueryFilter(query string, search structs.Search) string {
	var where string
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
