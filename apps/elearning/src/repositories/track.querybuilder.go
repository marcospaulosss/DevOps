package repositories

import (
	"fmt"

	"backend/apps/elearning/src/structs"
)

type TrackQueryBuilder struct {
	filtering FilteringQueryBuilder
}

func NewTrackQueryBuilder() TrackQueryBuilder {
	return TrackQueryBuilder{
		filtering: NewFilteringQueryBuilder(),
	}
}

func (this TrackQueryBuilder) Create() string {
	return `INSERT INTO tracks (title, description, duration, teachers, media)
    VALUES (:title, :description, :duration, :teachers, :media)
    RETURNING id`
}

func (this TrackQueryBuilder) Update() string {
	return `UPDATE tracks
    SET title = :title, description = :description, media = :media, teachers = :teachers, duration = :duration, updated_at = NOW()
    WHERE id = :id RETURNING id`
}

func (this TrackQueryBuilder) ReadOne() string {
	return `SELECT tr.* FROM tracks tr WHERE tr.id = $1`
}

func (this TrackQueryBuilder) ReadAll(search structs.Search) string {
	query := `SELECT tr.* FROM tracks tr`
	search.TablePrefix = "tr."
	return this.filtering.QueryFilter(query, search)
}

func (this TrackQueryBuilder) Delete() string {
	return `DELETE FROM tracks WHERE id = $1 RETURNING id`
}

func (this TrackQueryBuilder) Total(search structs.Search) string {
	query := `SELECT COUNT(tr.id) FROM tracks tr`
	var where string
	search.TablePrefix = "tr."
	if search.Where() != "" {
		where = fmt.Sprintf("WHERE %s", search.Where())
	}
	query = fmt.Sprintf("%s %s", query, where)
	return query
}

func (this TrackQueryBuilder) AssociateSubjectsAndTracks() string {
	return `INSERT INTO subjects_tracks (subject_id, track_id)
    VALUES ($1, $2)`
}
