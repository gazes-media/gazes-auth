package payloads

type UserRegister struct {
	Email    string `json:"email" validate:"regexp=^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"`
	Password string `json:"password" validate:"min=8"`
	Username string `json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z0-9_]+$"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"regexp=^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"`
	Password string `json:"password" validate:"min=8"`
}
