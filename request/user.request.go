package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Role     string `json:"role"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Role  string `json:"role"`
}

type UserResponse struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
