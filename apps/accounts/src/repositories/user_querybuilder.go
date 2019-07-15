package repositories

import (
	"fmt"

	"backend/apps/accounts/src/structs"
)

type UserQueryBuilder struct{}

func NewUserQueryBuilder() UserQueryBuilder {
	return UserQueryBuilder{}
}

func (this UserQueryBuilder) CreateUser() string {
	return "INSERT INTO users (name, email, phone) VALUES (:name, :email, :phone) RETURNING id;"
}

func (this UserQueryBuilder) ReadAllUsers(search structs.Search) string {
	query := "SELECT u.* FROM users AS u "
	query = this.applyReadAllFilters(query, search)
	return query
}

func (this UserQueryBuilder) CountUsers(search structs.Search) string {
	query := "SELECT count(u.id) as total FROM users AS u "
	// Disable ordenation and pagination counting users
	search.Pagination.Order = ""
	search.Pagination.SortBy = ""
	search.Pagination.Page = 0
	search.Pagination.PerPage = 0
	query = this.applyReadAllFilters(query, search)
	return query
}

func (this UserQueryBuilder) applyReadAllFilters(query string, search structs.Search) string {
	search.TablePrefix = "u."
	if search.Where() != "" {
		query += "WHERE " + search.Where()
	}

	pagination := search.Pagination
	if pagination.Order != "" && pagination.SortBy != "" {
		query = fmt.Sprintf(`%s ORDER BY %s %s`, query, pagination.Order, pagination.SortBy)
	}
	if pagination.PerPage > 0 && pagination.Page > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, pagination.PerPage, pagination.Offset())
	}

	return query
}
