package myValidator

import (
	"regexp"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func PhoneValidator(phone string) (err error) {
	//utilGin := response.Gin{Ctx: this}
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	r := reg.MatchString(phone)
	if r {
		return nil
	} else {
		return &errorString{"输入的手机号格式有误"}
	}
}
func PasswordValidator(password string) (err error) {
	regular1 := "^[a-zA-Z][a-zA-Z0-9_]{4,15}$"
	//regular1 := "^[a-zA-Z]\\w{5,17}$";
	reg := regexp.MustCompile(regular1)
	r := reg.MatchString(password)
	if r {
		return nil
	} else {
		return &errorString{"输入的密码号格式有误"}
	}
}
