package repositories

import (
	"fmt"

	"backend/apps/elearning/src/structs"
)

type SubjectQueryBuilder struct {
	filtering FilteringQueryBuilder
}

func NewSubjectQueryBuilder() SubjectQueryBuilder {
	return SubjectQueryBuilder{
		filtering: NewFilteringQueryBuilder(),
	}
}

func (this SubjectQueryBuilder) Create() string {
	return `INSERT INTO subjects (title) VALUES (:title) RETURNING id;`
}

func (this SubjectQueryBuilder) Update() string {
	return `UPDATE subjects SET title = :title, updated_at = NOW() WHERE id = :id RETURNING id;`
}

func (this SubjectQueryBuilder) ReadOne() string {
	return `SELECT * FROM subjects WHERE id = $1;`
}

func (this SubjectQueryBuilder) ReadAll(search structs.Search) string {
	query := `SELECT subj.* FROM subjects subj`

	search.TablePrefix = "subj."

	return this.filtering.QueryFilter(query, search)
}

func (this SubjectQueryBuilder) Total(search structs.Search) string {
	query := `SELECT COUNT(subj.id) FROM subjects subj`
	var where string
	if search.Where() != "" {
		where = fmt.Sprintf("WHERE %s", search.Where())
	}
	query = fmt.Sprintf("%s %s", query, where)
	return query
}

func (this SubjectQueryBuilder) Delete() string {
	return `DELETE FROM subjects WHERE id = $1 RETURNING id`
}
