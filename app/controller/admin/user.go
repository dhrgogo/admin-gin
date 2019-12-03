package admin

import (
	"admin-gin/app/model/dbmodel"
	"fmt"
	"github.com/gin-gonic/gin"
	//"admin-gin/app/model/dbmodel"
	"admin-gin/app/util/response"
	"gopkg.in/go-playground/validator.v9"
)

type FormData struct {
	Age      int    `validate:"min=12,max=15"`
	Password string `validate:"required"`
	Post     string `validate:"required"`
	//us.Birthday       = this.PostForm("birthday")
	RoleID       string `validate:"required"`
	English_name string `validate:"required"`
	Department   string `validate:"required"`
	//us.Sex            = this.PostForm("sex")
	Email      string `validate:"required"`
	Header_img string `validate:"required"`
	//is_deleted     := this.PostForm("is_deleted")
	Last_name string `validate:"required"`
	Nickname  string `validate:"required"`
	Openid    string `validate:"required"`
}

func UserCreat(this *gin.Context) {

	formDb := &FormData{
		Age:      13,
		Password: "11",
		Post:     "11",
		//us.Birthday       = this.PostForm("birthday")
		RoleID:       "11",
		English_name: "11",
		Department:   "11",
		//us.Sex            = this.PostForm("sex")
		Email:      "11",
		Header_img: "11",
		//is_deleted     := this.PostForm("is_deleted")
		Last_name: "11",
		Nickname:  "11",
		Openid:    "11",
	}
	validate := validator.New()
	err := validate.Struct(formDb)
	if err != nil {
		fmt.Println(err)
		return
	}

	utilGin := response.Gin{Ctx: this}
	var us dbmodel.User
	us.Name = this.PostForm("name")
	us.User_unique_id = this.PostForm("user_unique_id")
	us.Phone = this.PostForm("phone")
	//us.Password       = this.PostForm("password")
	//us.Post           = this.PostForm("post")
	////us.Birthday       = this.PostForm("birthday")
	//us.RoleID         = this.PostForm("roleID")
	//us.English_name   = this.PostForm("english_name")
	//us.Department     = this.PostForm("department")
	////us.Sex            = this.PostForm("sex")
	//us.Email          = this.PostForm("email")
	//us.Header_img     = this.PostForm("header_img")
	////is_deleted     := this.PostForm("is_deleted")
	//us.Last_name      = this.PostForm("last_name")
	//us.Nickname       = this.PostForm("nickname")
	//us.Openid         = this.PostForm("openid")
	fmt.Println("=========1", us.Phone)
	err = dbmodel.CreatUser(&us)
	fmt.Println(err)
	utilGin.Response(1000, "pong", us)
	//return
}

func UserUpdate(this *gin.Context) {
	utilGin := response.Gin{Ctx: this}
	var us dbmodel.User
	user_unique_id := this.Param("unique_id")
	us.Name = this.PostForm("name")
	//us.User_unique_id = this.PostForm("user_unique_id")
	us.Phone = this.PostForm("phone")
	//us.Password       = this.PostForm("password")
	//us.Post           = this.PostForm("post")
	////us.Birthday       = this.PostForm("birthday")
	//us.RoleID         = this.PostForm("roleID")
	//us.English_name   = this.PostForm("english_name")
	//us.Department     = this.PostForm("department")
	////us.Sex            = this.PostForm("sex")
	//us.Email          = this.PostForm("email")
	//us.Header_img     = this.PostForm("header_img")
	////is_deleted     := this.PostForm("is_deleted")
	//us.Last_name      = this.PostForm("last_name")
	//us.Nickname       = this.PostForm("nickname")
	//us.Openid         = this.PostForm("openid")
	//fmt.Println("=========1",string(data))
	err := dbmodel.UpdateUser(user_unique_id, &us)
	fmt.Println(err)
	utilGin.Response(1000, "pong", us)
}

func UserObj(this *gin.Context) {
	utilGin := response.Gin{Ctx: this}
	var us dbmodel.User
	id := this.Param("id")
	//id64, err := strconv.ParseInt(id, 10, 64)
	_, err := dbmodel.GetUserById(id, &us)
	fmt.Println(err)
	utilGin.Response(1000, "pong", us)

	//
	//if err != nil {
	//	utilGin.Response(500, "pong", err)
	//	return
	//}
	////userList := make([]*dbmodel.Role2, 0)
	////err = dbmodel.X.Find(&userList)
	////fmt.Println("=============err", err)
	////map1Byte, _ := json.Marshal(userRole1)
	//utilGin.Response(200, "用户列表", userList)

}
func UserList(this *gin.Context) {
	utilGin := response.Gin{Ctx: this}
	//var ro dbmodel.Role
	//role_unique_id := this.Param("unique_id")
	//id64, err := strconv.ParseInt(id, 10, 64)
	ro, err := dbmodel.GetUserListModule()
	fmt.Println(err)
	fmt.Println("---------------------", ro)
	utilGin.Response(1000, "pong", ro)
}
