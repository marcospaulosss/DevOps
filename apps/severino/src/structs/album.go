package structs

type Album struct {
	ID          uint64    `json:"id"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at,omitempty"`
	Title       string    `json:"title" validate:"required,min=1,max=250"`
	Description string    `json:"description"`
	Image       string    `json:"image" validate:"required,min=36"`
	IsPublished bool      `json:"is_published"`
	PublishedAt string    `json:"published_at,omitempty"`
	Teachers    []string  `json:"teachers,omitempty"`
	Shelves     []Shelf   `json:"shelves,omitempty"`
	Sections    []Section `json:"sections,omitempty"`
}
