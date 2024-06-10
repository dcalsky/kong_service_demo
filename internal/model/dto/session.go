package dto

type AccountForDetail struct {
	Id        uint
	Email     string
	NickName  string
	CreatedAt int64
	UpdatedAt int64
}

type LoginRequest struct {
	Email    string `json:",required"`
	Password string `json:",required"`
}

type LoginResponse struct {
	Token   string
	Account AccountForDetail
}

type RegisterRequest struct {
	Email    string `json:",required"`
	NickName string `json:",required"`
	Password string `json:",required"`
}

type RegisterResponse struct {
	Token   string
	Account AccountForDetail
}
