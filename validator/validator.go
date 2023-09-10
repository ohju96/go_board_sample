package validator

import (
	"fmt"
	custom "ginSample/handler/err"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidate() {
	validate = validator.New()
}

// 유효성 검사 메서드
func Validator(req interface{}) *custom.BadRequestRes {

	if err := validate.Struct(req); err != nil {
		return custom.NewBadRequestRes(custom.ERR_VALIDATOR)
	}

	return nil
}

// 요청 객체 바인딩 메서드
func Binder(g *gin.Context, req interface{}) *custom.BadRequestRes {

	if err := g.ShouldBind(req); err != nil {
		return custom.NewBadRequestRes(custom.ERR_BINDING)
	}

	return nil
}

// 유효성 검사 및 요청 객체 바인딩 메서드
func BinderAndValidator(g *gin.Context, req interface{}) *custom.BadRequestRes {

	if err := g.ShouldBind(req); err != nil {
		fmt.Println("#validator# 바인딩 에러 : ", err.Error())
		return custom.NewBadRequestRes(custom.ERR_BINDING)
	}

	if err := validate.Struct(req); err != nil {
		fmt.Println("#validator# 유효성 검사 에러 : ", err.Error())
		return custom.NewBadRequestRes(custom.ERR_VALIDATOR)
	}

	return nil
}
