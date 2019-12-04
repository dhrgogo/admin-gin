package admin

import (
	"admin-gin/app/model/dbmodel"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	//"admin-gin/app/model/dbmodel"
	"admin-gin/app/util/response"
)

func RoleCreat(this *gin.Context) {
	utilGin := response.Gin{Ctx: this}
	var ro dbmodel.Role
	Pid := this.PostForm("pid")
	pid, err := strconv.Atoi(Pid)
	ro.Pid = pid
	ro.Role_unique_id = this.PostForm("role_unique_id")
	ro.Role_name = this.PostForm("role_name")
	ro.RoleDesc = this.PostForm("roleDesc")
	ro.PName = this.PostForm("pName")
	//ro.Password       = this.PostForm("password")
	//ro.Post           = this.PostForm("post")
	////ro.Birthday       = this.PostForm("birthday")
	//ro.RoleID         = this.PostForm("roleID")
	//ro.English_name   = this.PostForm("english_name")
	//ro.Department     = this.PostForm("department")
	////ro.Sex            = this.PostForm("sex")
	//ro.Email          = this.PostForm("email")
	//ro.Header_img     = this.PostForm("header_img")
	////is_deleted     := this.PostForm("is_deleted")
	//ro.Last_name      = this.PostForm("last_name")
	//ro.Nickname       = this.PostForm("nickname")
	//ro.Openid         = this.PostForm("openid")
	fmt.Println("=========1", ro.Role_name)
	err = dbmodel.CreatRoleModule(&ro)
	fmt.Println(err)
	utilGin.Response(1000, "pong", ro)
	//return
}

func RoleUpdate(this *gin.Context) {
	utilGin := response.Gin{Ctx: this}
	var ro dbmodel.Role
	role_unique_id := this.Param("unique_id")
	//ro.Role_unique_id = this.PostForm("role_unique_id")
	ro.Role_name = this.PostForm("role_name")
	ro.RoleDesc = this.PostForm("roleDesc")
	Pid := this.PostForm("pid")
	pid, err := strconv.Atoi(Pid)
	ro.Pid = pid
	//ro.Password       = this.PostForm("password")
	//ro.Post           = this.PostForm("post")
	////ro.Birthday       = this.PostForm("birthday")
	//ro.RoleID         = this.PostForm("roleID")
	//ro.English_name   = this.PostForm("english_name")
	//ro.Department     = this.PostForm("department")
	////ro.Sex            = this.PostForm("sex")
	//ro.Email          = this.PostForm("email")
	//ro.Header_img     = this.PostForm("header_img")
	////is_deleted     := this.PostForm("is_deleted")
	//ro.Last_name      = this.PostForm("last_name")
	//ro.Nickname       = this.PostForm("nickname")
	//ro.Openid         = this.PostForm("openid")
	//fmt.Println("=========1",string(data))
	err = dbmodel.UpdateRoleModule(role_unique_id, &ro)
	fmt.Println(err)
	utilGin.Response(1000, "pong", ro)
}

func RoleObj(this *gin.Context) {
	utilGin := response.Gin{Ctx: this}
	var ro dbmodel.Role
	role_unique_id := this.Param("unique_id")
	//id64, err := strconv.ParseInt(id, 10, 64)
	_, err := dbmodel.GetRoleByIdModule(role_unique_id, &ro)
	fmt.Println(err)
	utilGin.Response(1000, "pong", ro)
}
func RoleList(this *gin.Context) {
	utilGin := response.Gin{Ctx: this}
	//var ro dbmodel.Role
	//role_unique_id := this.Param("unique_id")
	//id64, err := strconv.ParseInt(id, 10, 64)
	ro, err := dbmodel.GetRoleListModule()
	fmt.Println(err)
	fmt.Println("---------------------", ro)
	utilGin.Response(1000, "pong", ro)
}
