package user

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GeneralIdInput struct {
	ID int `uri:"id" binding:"required"`
}

type ChangePasswordInput struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
