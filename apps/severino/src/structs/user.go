package structs

type User struct {
	ID        string `struct:"id" json:"id"`
	Name      string `struct:"name" json:"name" validate:"required,max=255"`
	Email     string `struct:"email" json:"email" validate:"required,max=255,email"`
	EmailCode string `struct:"email_code,omitempty" json:"email_code,omitempty" validate:"required,len=6,numeric"`
	Phone     string `struct:"phone" json:"phone" validate:"required,len=11,numeric"`
	PhoneCode string `struct:"phone_code,omitempty" json:"phone_code,omitempty" validate:"required,len=6,numeric"`
	Active    bool   `struct:"active" json:"active,omitempty"`
	CreatedAt string `struct:"created_at" json:"created_at"`
	UpdatedAt string `struct:"updated_at,omitempty" json:"updated_at,omitempty"`
	Token     string `struct:"token,omitempty" json:"token,omitempty"`
}

type UserResult struct {
	Data interface{}
	Err  error
}

type UserResponse struct {
	Data  User   `json:"data"`
	Error string `json:"error"`
	Meta  Meta   `json:"meta"`
}

type UsersResponse struct {
	Data  []User `json:"data"`
	Error string `json:"error"`
	Meta  Meta   `json:"meta"`
}
