package authrequest

import (
	"github.com/go-playground/validator/v10"
)

// RegisterInput 定义注册请求所需的输入
type RegisterInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// Validate 实现自定义的验证方法
func (r *RegisterInput) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
