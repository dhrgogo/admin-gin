package dbmodel

import (
	//"fmt"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 用户表
type User struct {
	Id             int64     `xorm:"autoincr pk" json:"id"`
	User_unique_id string    `xorm:"pk varchar(32) notnull unique" json:"user_unique_id" description:"唯一id"`
	Name           string    `xorm:"varchar(32)" json:"name" description:"名字"`
	Phone          string    `xorm:"varchar(32) notnull unique" json:"phone" description:"手机号"`
	Password       string    `xorm:"varchar(32)" json:"Password" description:"密码"`
	Salt           string    `xorm:"varchar(32)" json:"salt" description:"盐"`
	Post           string    `xorm:"varchar(32)" json:"post" description:"职位"`
	Birthday       time.Time `json:"birthday"`
	RoleID         string    `json:"roleid" xorm:"roleid"`
	English_name   string    `xorm:"varchar(32)" json:"english_name" description:"英文名称"`
	Department     string    `xorm:"varchar(100)" json:"department" description:"部门"`
	Reorder        string    `xorm:"varchar(11)" json:"reorder" description:"自定义排序"`
	Sex            int       `xorm:"int(4) default 2" json:"sex" description:"0女、1男、2未知"`
	Email          string    `xorm:"varchar(80)" json:"email" description:"邮箱"`
	Header_img     string    `xorm:"varchar(255)" json:"header_img" description:"头像"`
	Is_deleted     int       `xorm:"int(4) default 1" json:"is_deleted" description:"0删除、1未删除"`
	Last_name      string    `xorm:"varchar(32)" json:"last_name" description:"姓氏"`
	Nickname       string    `xorm:"varchar(32)" json:"nickname" description:"昵称"`
	Openid         string    `xorm:"varchar(32)" json:"openid" description:"第三方登录唯一id"`
	Status         int       `xorm:"int(4) default 0" json:"status" description:"0未激活1已激活,2:待审核,3:已拒绝"`
	CreateTime     time.Time `xorm:"created" json:"create_time" description:"创建时间"`
	UpdateTime     time.Time `xorm:"updated" json:"update_time"description:"更新时间"`
}

func CreatUser(user *User) error {
	a, err := X.Insert(user)
	fmt.Println("=============a", a)
	fmt.Println("=============user", err)
	return err
}
func UpdateUser(unique_id string, user *User) (err error) {
	a, err := X.Where("user_unique_id = ?", unique_id).Update(user)
	fmt.Println("=============a", a)
	fmt.Println("=============err", err)
	return err
}
func GetUserById(unique_id string, user *User) (has *User, err error) {
	fmt.Println("============", unique_id)
	//user := new(User)
	_, err = X.Where("user_unique_id=?", unique_id).Get(user)
	fmt.Println("============", user)
	return user, err
	//return user, err
	//userList := make([]*User, 0)
	////userRole := new([]*Role2)
	//err = X.Find(&userList)
	//if err != nil {
	//	fmt.Println("-----------err", err)
	//}
	//fmt.Println("-------------pEveryOnepEveryOne", userList)
	////json.Unmarshal(userRole, &addrList)
	//return userList, err
}
func getUserList() {

}
func GetUserListModule() (userList []*User, err error) {
	//list := make([]*Role, 0)
	//userRole := new([]*Role2)
	err = X.Find(&userList)
	if err != nil {
		fmt.Println("-----------err", err)
	}
	fmt.Println("-------------pEveryOnepEveryOne", userList)
	//json.Unmarshal(userRole, &addrList)
	return userList, err
}
