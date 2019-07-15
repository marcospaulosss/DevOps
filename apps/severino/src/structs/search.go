package structs

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

type Pagination struct {
	PerPage int    `json:"per_page"`
	Page    int    `json:"page"`
	Total   int    `json:"total"`
	Order   string `json:"order,omitempty"`
	SortBy  string `json:"sort,omitempty"`
}

type Result struct {
	Items      interface{}
	Pagination Pagination
}

type Search struct {
	Pagination Pagination `json:"pagination,omitempty"`
	Raw        string     `json:"-"`
	Extra      string     `json:"-"`
}

func NewSearch(params url.Values) Search {
	p := Pagination{
		PerPage: getIntValueOrDefault(params, "per_page", 0),
		SortBy:  "desc",
		Order:   "created_at",
	}
	sort := getStrValueOrDefault(params, "sort", "-created_at")
	if strings.HasPrefix(sort, "+") {
		p.SortBy = "asc"
	}
	if strings.HasPrefix(sort, "-") || strings.HasPrefix(sort, "+") {
		p.Order = sort[1:]
	}

	page := getIntValueOrDefault(params, "page", 0)
	if page > 0 {
		p.Page = int(math.Max(1, float64(page)))
	}

	s := Search{Pagination: p}
	raw := getStrValueOrDefault(params, "search", "")
	if strings.ContainsAny(raw, "( & )") {
		s.Raw = raw
	}
	return s
}

func (this Search) String() string {
	return fmt.Sprintf("search: %#v, pagination: %#v", this.Raw, this.Pagination)
}

func getStrValueOrDefault(params url.Values, key, value string) string {
	v := params.Get(key)
	if v == "" {
		v = value
	}
	return v
}

func getIntValueOrDefault(params url.Values, key string, value int) int {
	v := getStrValueOrDefault(params, key, "")
	if v == "" {
		return value
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return value
	}
	return i
}
