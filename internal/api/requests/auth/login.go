package authrequest

import (
	"github.com/go-playground/validator/v10"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Validate 实现自定义的验证方法
func (r *LoginInput) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
