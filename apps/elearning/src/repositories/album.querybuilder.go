package repositories

import (
	"encoding/json"
	"fmt"
	"strings"

	"backend/apps/elearning/src/structs"
)

type AlbumQueryBuilder struct{}

func NewAlbumQueryBuilder() AlbumQueryBuilder {
	return AlbumQueryBuilder{}
}

func (this AlbumQueryBuilder) Create() string {
	return `INSERT INTO albums (title, description, image, is_published)
    VALUES (:title, :description, :image, :is_published)
    RETURNING id`
}

func (this AlbumQueryBuilder) Update() string {
	return `UPDATE albums SET title = :title, description = :description, image = :image, updated_at = NOW() WHERE id = :id RETURNING id`
}

func (this AlbumQueryBuilder) CreateSection() string {
	return `INSERT INTO sections (album_id, title, description, position) VALUES ($1, $2, $3, $4) RETURNING id`
}

func (this AlbumQueryBuilder) AssociateTracksToSection(section structs.Section) string {
	var values []string
	for index, track := range section.Tracks {
		values = append(values, fmt.Sprintf("(:section_id, %d, %d)", track.ID, index))
	}
	return fmt.Sprintf(`INSERT INTO sections_tracks (section_id, track_id, track_position) VALUES %s`, strings.Join(values, ","))
}

func (this AlbumQueryBuilder) DeleteSectionsByAlbumID() string {
	return `DELETE FROM sections WHERE album_id = :album_id`
}

func (this AlbumQueryBuilder) ReadAll(search structs.Search) string {
	query := `SELECT a.* FROM albums a`

	where := "WHERE 1=1"
	search.TablePrefix = "a."

	if search.Extra != "" {
		var extra map[string]interface{}
		json.Unmarshal([]byte(search.Extra), &extra)
		for key, value := range extra {
			where = fmt.Sprintf(`%s AND %s%s = '%s'`, where, search.TablePrefix, key, value)
		}
	}

	if search.Where() != "" {
		where = fmt.Sprintf("%s AND %s", where, search.Where())
	}

	query = fmt.Sprintf("%s %s GROUP BY a.id", query, where)

	pagination := search.Pagination
	if pagination.Order != "" && pagination.SortBy != "" {
		query = fmt.Sprintf(`%s ORDER BY a.%s %s`, query, pagination.Order, pagination.SortBy)
	}
	if pagination.PerPage > 0 && pagination.Page > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, pagination.PerPage, pagination.Offset())
	}
	return query
}

func (this AlbumQueryBuilder) GetSectionsByAlbumID() string {
	return `SELECT * FROM sections WHERE album_id = $1 ORDER BY position ASC`
}

func (this AlbumQueryBuilder) GetTracksBySectionID() string {
	return `SELECT tr.* FROM tracks tr, sections_tracks st WHERE tr.id = st.track_id AND st.section_id = $1 ORDER BY st.track_position ASC`
}

func (this AlbumQueryBuilder) FindAlbumsByShelfID() string {
	return `SELECT a.* FROM albums a, albums_shelves ash WHERE ash.album_id = a.id AND ash.shelf_id = $1 ORDER BY ash.album_position ASC`
}

func (this AlbumQueryBuilder) FindPublishedAlbumsByShelfID() string {
	return `SELECT a.* FROM albums a, albums_shelves ash WHERE ash.album_id = a.id AND a.is_published = true AND ash.shelf_id = $1 ORDER BY ash.album_position ASC`
}

func (this AlbumQueryBuilder) FindAlbumsByTrackID() string {
	return `SELECT a.*
    FROM albums a
        LEFT JOIN sections s ON a.id = s.album_id           
        LEFT JOIN sections_tracks st ON st.section_id = s.id
        LEFT JOIN tracks tr on st.track_id = tr.id
    WHERE tr.id = $1
    GROUP BY a.id, tr.id`
}

func (this AlbumQueryBuilder) GetTeachersByAlbumID() string {
	return `SELECT string_agg(DISTINCT tr.teachers, ',') as teachers
    FROM tracks tr, sections s, sections_tracks st, albums a
    WHERE a.id = $1 AND a.id = s.album_id AND s.id = st.section_id AND st.track_id = tr.id LIMIT 1`
}

func (this AlbumQueryBuilder) ReadOne() string {
	return `SELECT
        CAST(to_jsonb(a.*) AS TEXT) as album,
        CAST(jsonb_agg(DISTINCT ta.*) AS TEXT) AS shelves
    FROM albums a
        LEFT JOIN albums_shelves at ON a.id = at.album_id
        LEFT JOIN shelves ta ON at.shelf_id = ta.id
    WHERE a.id = $1
    GROUP BY a.id`
}

func (this AlbumQueryBuilder) Delete() string {
	return `DELETE FROM albums WHERE id = $1 RETURNING id`
}

func (this AlbumQueryBuilder) Total(search structs.Search) string {
	query := `SELECT COUNT(id) FROM albums`
	where := "WHERE 1=1"
	if search.Extra != "" {
		var extra map[string]interface{}
		json.Unmarshal([]byte(search.Extra), &extra)
		for key, value := range extra {
			where = fmt.Sprintf(`%s AND %s%s = '%s'`, where, search.TablePrefix, key, value)
		}
	}
	if search.Where() != "" {
		where = fmt.Sprintf("%s AND %s", where, search.Where())
	}
	query = fmt.Sprintf("%s %s", query, where)
	return query
}

func (this AlbumQueryBuilder) Publish() string {
	return `UPDATE albums SET is_published = true, published_at = NOW() WHERE id = $1`
}

func (this AlbumQueryBuilder) Unpublish() string {
	return `UPDATE albums SET is_published = false WHERE id = $1`
}
