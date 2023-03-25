package dto

type CreateActivity struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateActivity struct {
	Title string `json:"title" validate:"required"`
}
