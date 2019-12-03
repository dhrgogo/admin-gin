package dbmodel

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 用户表
type UserRole struct {
	User_unique_id string    `xorm:"pk varchar(100) notnull pk" json:"user_unique_id" description:"关联role表id"` // 用户ID
	Role_unique_id string    `xorm:"pk varchar(100) notnull pk" json:"role_unique_id" description:"关联role表id"` // 档案号
	Is_deleted     int       `xorm:"int(4) default 0" json:"is_deleted"   description:"是否删除"`                  // 姓名
	CreateTime     time.Time `xorm:"created" json:"create_time" description:"创建时间"`
	UpdateTime     time.Time `xorm:"updated" json:"update_time"description:"更新时间"`
}
