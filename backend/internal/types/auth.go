package types

type AuthRequst struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	Expired int `json:"expired"`
	User UserResponse `json:"user"`
}
