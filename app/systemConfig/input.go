package systemConfig

type CreateSystemConfigInput struct {
	Code  string `json:"code" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type GeneralIdInput struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateSystemConfigInput struct {
	Code  string `json:"code" binding:"required"`
	Value string `json:"value" binding:"required"`
}
