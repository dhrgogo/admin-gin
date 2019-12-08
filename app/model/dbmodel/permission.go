package dbmodel

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 权限表
type Permission struct {
	Id                   int64  `xorm:"pk autoincr unique notnull" json:"id,omitempty"`
	Permission_unique_id string `xorm:"pk varchar(32) notnull unique" json:"permission_unique_id" description:"权限id"`
	Permission_name      string `xorm:"varchar(32) notnull unique" json:"permission_name" description:"名字"`
	PermissionDesc       string `xorm:"varchar(32)" json:"permissionDesc" description:"描述"`
	Is_deleted           int    `xorm:"int(4) default 1" json:"is_deleted" description:"0删除、1未删除"` // 姓名
	//Pid            int       `json:"pId" form:"pid" xorm:"pId"`
	//PName          string    `json:"pName" xorm:"<- p_name"`

	CreateTime time.Time `xorm:"created" json:"create_time" description:"创建时间"`
	//CreateTime     time.Time `json:"createTime" xorm:"created 'createtime'"`
	UpdateTime time.Time `xorm:"updated" json:"update_time"description:"更新时间"`
}

func CreatPermissionModule(permission *Permission) (err error) {
	a, err := X.Insert(permission)
	fmt.Println("=============a", a)
	fmt.Println("=============user", err)
	return err
}
func UpdatePermissionModule(unique_id string, permission *Permission) (err error) {
	fmt.Println("=============unique_id", unique_id)

	a, err := X.Where("role_unique_id = ?", unique_id).Update(permission)
	fmt.Println("=============a", a)
	fmt.Println("=============err", err)
	return err
}
func GetPermissionByIdModule(unique_id string, permission *Permission) (has *Permission, err error) {
	fmt.Println("============", unique_id)
	//user := new(User)
	_, err = X.Where("role_unique_id=?", unique_id).Get(permission)
	fmt.Println("============", has)
	return has, err
}
func GetPermissionListModule() (roleList []*Permission, err error) {
	//list := make([]*Permission, 0)
	//userRole := new([]*Role2)
	err = X.Find(&roleList)
	if err != nil {
		fmt.Println("-----------err", err)
	}
	fmt.Println("-------------pEveryOnepEveryOne", roleList)
	//json.Unmarshal(userRole, &addrList)
	return roleList, err
}
