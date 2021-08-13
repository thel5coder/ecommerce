package view_models

import (
	"github.com/ecommerce-service/user-service/domain/models"
	"time"
)

type UserVm struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      RoleVm `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserVm(model *models.User) UserVm {
	return UserVm{
		ID:        model.Id(),
		Email:     model.Email(),
		FirstName: model.FirstName(),
		LastName:  model.LastName(),
		Role:      NewRoleVm(model.Role),
		CreatedAt: model.CreatedAt().Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt().Format(time.RFC3339),
	}
}
