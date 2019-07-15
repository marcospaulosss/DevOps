package structs

type Shelf struct {
	ID        uint64  `json:"id"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
	Title     string  `json:"title,omitempty" validate:"required,min=1,max=250"`
	Albums    []Album `json:"albums,omitempty"`
}
