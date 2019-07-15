package structs

import (
	"encoding/json"
	"time"
)

type Album struct {
	ID          uint64     `db:"id" json:"id"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	Image       string     `db:"image" json:"image"`
	IsPublished bool       `db:"is_published" json:"is_published"`
	PublishedAt *time.Time `db:"published_at" json:"published_at"`
	Teachers    string     `db:"teachers" json:"teachers"`
	Shelves     []Shelf    `json:"shelves"`
	Sections    []Section  `json:"sections"`
}

type Shelf struct {
	ID        uint64     `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Title     string     `db:"title"`
	Albums    []Album    `json:"albums"`
	Total     int32      `db:"total"`
}

type Track struct {
	ID          uint64     `db:"id" json:"id"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	Teachers    string     `db:"teachers" json:"teachers"`
	Duration    int64      `db:"duration" json:"duration"`
	Media       string     `db:"media" json:"media"`
	Albums      []Album    `json:"albums"`
	Subject     Subject    `json:"subjects"`
}

type Section struct {
	ID          uint64    `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	AlbumID     uint64    `db:"album_id" json:"-"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Position    int       `db:"position" json:"position"`
	Tracks      []Track   `json:"tracks"`
}

type Preference struct {
	Type    string `db:"type"`
	Content string `db:"content"`
}

type Home struct {
	Shelves []uint64 `json:"shelves"`
}

type account struct {
	ID        string `struct:"id" json:"id"`
	Type      string `struct:"type" json:"type" validate:"required,eq=email|eq=phone"`
	Email     string `struct:"email" json:"email" validate:"omitempty,email,contains=@"`
	Phone     string `struct:"phone" json:"phone" validate:"omitempty,len=11,numeric"`
	Code      string `struct:"code" json:"code" validate:"omitempty,len=6,numeric"`
	Exists    bool
	CreatedAt string `struct:"created_at" json:"created_at"`
	UserID    string `json:"user_id" struct:"user_id"`
}

type Subject struct {
	ID        uint64     `db:"id"`
	Title     string     `db:"title"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func ParseAlbums(i interface{}) []Album {
	var albums []Album
	json.Unmarshal([]byte(i.(string)), &albums)
	return albums
}

func ParseAlbum(i interface{}) Album {
	var album Album
	json.Unmarshal([]byte(i.(string)), &album)
	return album
}

func ParseTrack(i interface{}) Track {
	var track Track
	json.Unmarshal([]byte(i.(string)), &track)
	return track
}

func ParseSections(i interface{}) []Section {
	var sections []Section
	json.Unmarshal([]byte(i.(string)), &sections)
	return sections
}

func ParseTracks(i interface{}) []Track {
	var tracks []Track
	json.Unmarshal([]byte(i.(string)), &tracks)
	return tracks
}

func ParseShelves(i interface{}) []Shelf {
	var shelves []Shelf
	json.Unmarshal([]byte(i.(string)), &shelves)
	return shelves
}

func ParseShelf(i interface{}) Shelf {
	var shelf Shelf
	json.Unmarshal([]byte(i.(string)), &shelf)
	return shelf
}

func ParseSubjects(i interface{}) []Subject {
	var subjects []Subject
	json.Unmarshal([]byte(i.(string)), &subjects)
	return subjects
}

func ParseSubject(i interface{}) Subject {
	var subject Subject
	json.Unmarshal([]byte(i.(string)), &subject)
	return subject
}
