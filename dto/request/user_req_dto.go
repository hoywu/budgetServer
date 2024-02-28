package request

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
