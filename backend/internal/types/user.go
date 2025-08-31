package types

type UserModel struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	Token string `json:"string"`
}

type UserRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Avatar string `json:"avatar"`
}