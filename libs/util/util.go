package util

import (
	"encoding/json"
	"log"
	"time"
)

func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func PtrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func StrToPtr(s string) *string {
	return &s
}

func StrToTime(str string) time.Time {
	if str == "" {
		return time.Now()
	}
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		log.Println("Failed to parse date", str, " Error:", err)
		return time.Now()
	}
	return t
}

func StrToArr(str string) []string {
	var v []string
	err := json.Unmarshal([]byte(str), &v)
	if err != nil {
		return []string{}
	}
	return v
}

func ArrToStr(a []string) string {
	s, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(s)
}

func StrToJSONArr(s string) []map[string]interface{} {
	var m []map[string]interface{}
	json.Unmarshal([]byte(s), &m)
	return m
}

func JSONArrToStr(m []map[string]interface{}) string {
	v, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(v)
}
