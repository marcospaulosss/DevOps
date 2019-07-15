package structs

import (
	"fmt"
	"math"
)

type Search struct {
	Pagination Pagination
	Term       string
}

func (this Search) String() string {
	return fmt.Sprintf("pagination: %#v, term: %s", this.Pagination, this.Term)
}

type Pagination struct {
	Order   string
	SortBy  string
	PerPage int32
	Page    int32
	Total   int32
}

func (p Pagination) Offset() int {
	page := int(math.Max(1, float64(p.Page)))
	return (page - 1) * int(p.PerPage)
}

type Result struct {
	Items      interface{}
	Pagination Pagination
}
