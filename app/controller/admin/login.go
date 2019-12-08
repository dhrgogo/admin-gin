package admin

import (
	"admin-gin/app/util/error"
	"admin-gin/app/util/myValidator"
	"admin-gin/app/util/response"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	//"fmt"
	"github.com/gin-gonic/gin"
)

type Account struct {
	Phone      string `validate:"is-Phone"`
	Password   string `validate:"is-Password"`
	NumCode    string `validate:"required"`
	LetterCode string `validate:"required"`
}

//func ValidatePhone(fl govalidator.FieldLevel) bool {
//	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
//	reg := regexp.MustCompile(regular)
//	r := reg.MatchString(fl.Field().String())
//	return r
//}
//func ValidatePassword(fl govalidator.FieldLevel) bool {
//	regular := "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,17}$"
//	//regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
//
//	reg := regexp.MustCompile(regular)
//	r := reg.MatchString(fl.Field().String())
//	return r
//}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
	Age    int    `validate:"min=12,max=15"`
}

func Login(this *gin.Context) {
	utilGin := response.Gin{Ctx: this}
	//var student Address
	address := &Address{
		Street: "Eavesdown Docks",
		City:   "beijing",
		Planet: "Persphone",
		Phone:  "none",
		Age:    7,
	}
	r := myValidator.PhoneValidator("15503083260")
	fmt.Println("==============1", r)
	if r != nil {
		error.ErrorNew(r.Error(), this)
		return
	}
	r1 := myValidator.PasswordValidator("11")
	fmt.Println("==============2", r)
	if r1 != nil {
		error.ErrorNew(r.Error(), this)
		return
	}
	validate := validator.New()
	err := validate.Struct(address)
	if err != nil {
		error.ErrorNew(err.Error(), this)
		return
	}
	utilGin.Response(1000, "pong", nil)
	return
}

//func UserUpdate(this *gin.Context) {
//	utilGin := response.Gin{Ctx: this}
//	var us dbmodel.User
//	user_unique_id := this.Param("unique_id")
//	us.Name = this.PostForm("name")
//	//us.User_unique_id = this.PostForm("user_unique_id")
//	us.Phone = this.PostForm("phone")
//	//us.Password       = this.PostForm("password")
//	//us.Post           = this.PostForm("post")
//	////us.Birthday       = this.PostForm("birthday")
//	//us.RoleID         = this.PostForm("roleID")
//	//us.English_name   = this.PostForm("english_name")
//	//us.Department     = this.PostForm("department")
//	////us.Sex            = this.PostForm("sex")
//	//us.Email          = this.PostForm("email")
//	//us.Header_img     = this.PostForm("header_img")
//	////is_deleted     := this.PostForm("is_deleted")
//	//us.Last_name      = this.PostForm("last_name")
//	//us.Nickname       = this.PostForm("nickname")
//	//us.Openid         = this.PostForm("openid")
//	//fmt.Println("=========1",string(data))
//	err := dbmodel.UpdateUser(user_unique_id, &us)
//	fmt.Println(err)
//	utilGin.Response(1000, "pong", us)
//}
//func UserObj(this *gin.Context) {
//	utilGin := response.Gin{Ctx: this}
//	var us dbmodel.User
//	id := this.Param("id")
//	//id64, err := strconv.ParseInt(id, 10, 64)
//	_, err := dbmodel.GetUserById(id, &us)
//	fmt.Println(err)
//	utilGin.Response(1000, "pong", us)
//
//	//
//	//if err != nil {
//	//	utilGin.Response(500, "pong", err)
//	//	return
//	//}
//	////userList := make([]*dbmodel.Role2, 0)
//	////err = dbmodel.X.Find(&userList)
//	////fmt.Println("=============err", err)
//	////map1Byte, _ := json.Marshal(userRole1)
//	//utilGin.Response(200, "用户列表", userList)
//
//}
//func UserList(this *gin.Context) {
//	utilGin := response.Gin{Ctx: this}
//	//var ro dbmodel.Role
//	//role_unique_id := this.Param("unique_id")
//	//id64, err := strconv.ParseInt(id, 10, 64)
//	ro, err := dbmodel.GetUserListModule()
//	fmt.Println(err)
//	fmt.Println("---------------------", ro)
//	utilGin.Response(1000, "pong", ro)
//}
