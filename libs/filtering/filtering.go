package filtering

import (
	"fmt"
	"regexp"
	"strings"

	log "backend/libs/logger"
)

var operators = map[string]string{
	"eq":        "=",
	"ne":        "<>",
	"gt":        ">",
	"lt":        "<",
	"gte":       ">=",
	"lte":       "<=",
	"contains":  "LIKE",
	"icontains": "ILIKE",
}

type Search struct {
	Raw string
}

func (s Search) Where(prefix string) string {
	if s.Raw == "" || strings.ContainsAny(s.Raw, "=") || !strings.ContainsAny(s.Raw, "( & )") {
		log.Info("Parametros de busca em branco ou invalidos. Vou pular o filtro.", s.Raw)
		return ""
	}
	var clause string
	r, _ := regexp.Compile(`\(\w+[^(,|\+)]+`)
	first := r.FindString(s.Raw)
	clause = toSQL(prefix, first)

	r, _ = regexp.Compile(`,([^(,\+)]+)`)
	all := r.FindAllStringSubmatch(s.Raw, -1)
	for _, item := range all {
		clause = fmt.Sprintf("%s OR %s", clause, toSQL(prefix, item[1]))
	}

	r, _ = regexp.Compile(`\+([^(,\+)]+)`)
	all = r.FindAllStringSubmatch(s.Raw, -1)
	for _, item := range all {
		clause = fmt.Sprintf("%s AND %s", clause, toSQL(prefix, item[1]))
	}
	return clause
}

func toSQL(prefix, raw string) string {
	r, _ := regexp.Compile(`\w+`)
	field := r.FindString(raw)
	if field == "" {
		return ""
	}

	var operator, value string
	r, _ = regexp.Compile(`\[([^]]+)\]`)
	groups := r.FindStringSubmatch(raw)
	if len(groups) == 0 {
		return ""
	}
	operator = groups[1]
	op := operators[operator]
	if op == "" {
		op = "="
	}

	r, _ = regexp.Compile(`\'(.*?)\'`)
	groups = r.FindStringSubmatch(raw)
	if len(groups) == 0 {
		return ""
	}
	value = groups[1]
	if value == "" {
		return ""
	}

	if strings.ContainsAny(op, "LIKE") && !strings.ContainsAny(value, "%") {
		value = "%" + value + "%"
	}

	str := fmt.Sprintf(`%s%s %s `, prefix, field, op)
	return str + "'" + value + "'"
}
