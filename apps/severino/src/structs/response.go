package structs

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"error,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}
