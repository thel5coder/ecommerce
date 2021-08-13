package requests

type UserEditRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password"`
	RoleId    int64  `json:"role_id" validate:"required"`
}
