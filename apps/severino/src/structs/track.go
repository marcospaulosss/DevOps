package structs

type Track struct {
	ID          uint64  `json:"id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Title       string  `json:"title" validate:"required,min=5,max=250"`
	Description string  `json:"description"`
	Teachers    string  `json:"teachers"`
	Duration    int64   `json:"duration" validate:"numeric"`
	Media       string  `json:"media" validate:"required,min=36"`
	Albums      []Album `json:"albums"`
	Subject     uint64  `json:"subject" validate:"required,numeric"`
}
