package structs

import (
	"fmt"
	"math"
	"strings"

	"backend/libs/filtering"
)

type Search struct {
	Pagination  Pagination
	Raw         string
	TablePrefix string
}

type Pagination struct {
	Order   string
	SortBy  string
	PerPage int32
	Page    int32
	Total   int32
}

type Result struct {
	Items      interface{}
	Pagination Pagination
}

func (this Search) Where() string {
	s := filtering.Search{Raw: this.Raw}
	x := s.Where(this.TablePrefix)
	return x
}

func (this Search) String() string {
	p := this.Pagination
	raw := strings.Replace(this.Raw, "[", " ", -1)
	raw = strings.Replace(raw, "]", " ", -1)
	return fmt.Sprintf("raw:%s page:%d per_page:%d order:%s sort_by:%s", raw, p.Page, p.PerPage, p.Order, p.SortBy)
}

func (p Pagination) Offset() int {
	page := int(math.Max(1, float64(p.Page)))
	return (page - 1) * int(p.PerPage)
}
