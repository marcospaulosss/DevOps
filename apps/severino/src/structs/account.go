package structs

type Account struct {
	ID        string `json:"id"`
	Type      string `json:"type" validate:"required,eq=email|eq=phone"`
	Email     string `json:"email" validate:"omitempty,email,contains=@"`
	Phone     string `json:"phone" validate:"omitempty,len=11,numeric"`
	Code      string `json:"code" validate:"omitempty,len=6,numeric"`
	Exists    bool   `validate:"isdefault=false"`
	CreatedAt string `json:"created_at"`
	UserID    string `json:"user_id"`
}
