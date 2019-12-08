package dbmodel

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 权限+角色表
type RolePermissions struct {
	Id                   int64     `xorm:"pk autoincr unique notnull" json:"id,omitempty"`
	Role_unique_id       string    `xorm:"pk varchar(32) notnull unique" json:"role_unique_id" description:"角色id"`
	Permission_unique_id string    `xorm:"pk varchar(32) notnull unique" json:"permission_unique_id" description:"名字"`
	Is_deleted           int       `xorm:"int(4) default 1" json:"is_deleted" description:"0删除、1未删除"` // 姓名
	CreateTime           time.Time `xorm:"created" json:"create_time" description:"创建时间"`
	UpdateTime           time.Time `xorm:"updated" json:"update_time"description:"更新时间"`
}
