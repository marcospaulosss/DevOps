package testutil

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"backend/apps/severino/src/structs"
)

func DoGET(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

func DoDELETE(url string) *http.Request {
	req, _ := http.NewRequest("DELETE", url, nil)
	return req
}

func DoPOST(url, jsonReq string) *http.Request {
	return makeRequest("POST", url, jsonReq)
}

func DoPUT(url, jsonReq string) *http.Request {
	return makeRequest("PUT", url, jsonReq)
}

func makeRequest(method, url, jsonReq string) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(jsonReq)))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func ToJSON(m interface{}) string {
	v, err := json.Marshal(m)
	if err != nil {
		log.Println("Failed to marshal in testutil.")
		return ""
	}
	return string(v)
}

func DecodePagination(m interface{}) structs.Pagination {
	str, _ := json.Marshal(m)
	var item structs.Pagination
	json.Unmarshal([]byte(str), &item)
	return item
}

func DecodeAlbum(m interface{}) structs.Album {
	str, _ := json.Marshal(m)
	var item structs.Album
	json.Unmarshal([]byte(str), &item)
	return item
}

func DecodeAccount(m interface{}) structs.Account {
	str, _ := json.Marshal(m)
	var item structs.Account
	json.Unmarshal([]byte(str), &item)
	return item
}

func DecodeAlbums(m interface{}) []structs.Album {
	str, _ := json.Marshal(m)
	var items []structs.Album
	json.Unmarshal([]byte(str), &items)
	return items
}

func DecodeTrack(m interface{}) structs.Track {
	str, _ := json.Marshal(m)
	var item structs.Track
	json.Unmarshal([]byte(str), &item)
	return item
}

func DecodeTracks(m interface{}) []structs.Track {
	str, _ := json.Marshal(m)
	var items []structs.Track
	json.Unmarshal([]byte(str), &items)
	return items
}

func DecodeShelf(m interface{}) structs.Shelf {
	str, _ := json.Marshal(m)
	var item structs.Shelf
	json.Unmarshal([]byte(str), &item)
	return item
}

func DecodeShelves(m interface{}) []structs.Shelf {
	str, _ := json.Marshal(m)
	var items []structs.Shelf
	json.Unmarshal([]byte(str), &items)
	return items
}

func DecodeUser(m interface{}) structs.User {
	str, _ := json.Marshal(m)
	var item structs.User
	json.Unmarshal([]byte(str), &item)
	return item
}

func DecodeUsers(m interface{}) []structs.User {
	str, _ := json.Marshal(m)
	var items []structs.User
	json.Unmarshal([]byte(str), &items)
	return items
}

func DecodeSubject(m interface{}) structs.Subject {
	str, _ := json.Marshal(m)
	var item structs.Subject
	json.Unmarshal([]byte(str), &item)
	return item
}

func DecodeSubjects(m interface{}) []structs.Subject {
	str, _ := json.Marshal(m)
	var items []structs.Subject
	json.Unmarshal([]byte(str), &items)
	return items
}
