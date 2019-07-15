package structs

type Meta struct {
	Page       int32  `json:"page" struct:"page"`
	PerPage    int32  `json:"per_page" struct:"per_page"`
	TotalItems int32  `json:"total_items" struct:"total_items"`
	Order      string `json:"order" struct:"order"`
	Sort       string `json:"sort" struct:"sort"`
}
