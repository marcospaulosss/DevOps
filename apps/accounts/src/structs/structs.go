package structs

type Account struct {
	ID        string `json:"id" struct:"id"`
	Email     string `json:"email" struct:"email"`
	EmailCode string `json:"email_code" struct:"email_code" db:"email_code"`
	Phone     string `json:"phone" struct:"phone"`
	PhoneCode string `json:"phone_code" struct:"phone_code" db:"phone_code"`
	CreatedAt string `json:"created_at" struct:"created_at" db:"created_at"`
	Exists    bool   `json:"exists" struct:"exists"`
	Type      string
	UserID    string `json:"user_id" struct:"user_id" db:"user_id"`
}

type User struct {
	ID        string  `db:"id",json:"id",struct:"id"`
	Name      string  `db:"name",json:"name",struct:"name,omitempty"`
	Email     string  `db:"email",json:"email",struct:"email,omitempty"`
	Phone     string  `db:"phone",json:"phone",struct:"phone,omitempty"`
	Active    bool    `db:"active",json:"active",struct:"active"`
	CreatedAt string  `db:"created_at",json:"created_at",struct:"created_at"`
	UpdatedAt *string `db:"updated_at",json:"updated_at",struct:"updated_at"`
	Total     int32   `db:"total",json:"total",struct:"total"`
}
