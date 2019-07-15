package structs

type Preference struct {
	Type    string   `json:"type"`
	Shelves []uint64 `json:"shelves,omitempty" validate:"required"`
	Content string   `json:"-"`
}
