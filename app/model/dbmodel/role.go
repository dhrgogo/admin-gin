package dbmodel

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 角色表
type Role struct {
	Id             int64     `xorm:"autoincr pk" json:"id"`
	Role_unique_id string    `xorm:"pk varchar(32) notnull unique" json:"role_unique_id" description:"角色id"`
	Role_name      string    `xorm:"varchar(32)" json:"role_name" description:"名字"`
	RoleDesc       string    `xorm:"varchar(32) notnull unique" json:"roleDesc" description:"手机号"`
	Is_deleted     int       `xorm:"int(4) default 1" json:"is_deleted" description:"0删除、1未删除"` // 姓名
	Pid            int       `json:"pId" form:"pid" xorm:"pId"`
	PName          string    `json:"pName" xorm:"<- p_name"`
	CreateTime     time.Time `xorm:"created" json:"create_time" description:"创建时间"`
	//CreateTime     time.Time `json:"createTime" xorm:"created 'createtime'"`
	UpdateTime time.Time `xorm:"updated" json:"update_time"description:"更新时间"`
}

func CreatRoleModule(role *Role) (err error) {
	a, err := X.Insert(role)
	fmt.Println("=============a", a)
	fmt.Println("=============user", err)
	return err
}
func UpdateRoleModule(unique_id string, role *Role) (err error) {
	fmt.Println("=============unique_id", unique_id)

	a, err := X.Where("role_unique_id = ?", unique_id).Update(role)
	fmt.Println("=============a", a)
	fmt.Println("=============err", err)
	return err
}
func GetRoleByIdModule(unique_id string, role *Role) (has *Role, err error) {
	fmt.Println("============", unique_id)
	//user := new(User)
	_, err = X.Where("role_unique_id=?", unique_id).Get(role)
	fmt.Println("============", has)
	return has, err
}
func GetRoleListModule() (roleList []*Role, err error) {
	//list := make([]*Role, 0)
	//userRole := new([]*Role2)
	err = X.Find(&roleList)
	if err != nil {
		fmt.Println("-----------err", err)
	}
	fmt.Println("-------------pEveryOnepEveryOne", roleList)
	//json.Unmarshal(userRole, &addrList)
	return roleList, err
}
