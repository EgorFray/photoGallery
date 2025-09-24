package types

type UserModel struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
}

type UserRequest struct {
	Name string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}