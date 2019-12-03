package dbmodel

import (
	"fmt"
	//"illness/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid" // "os"
	"github.com/xorm.io/xorm"
	"time"
)

var X *xorm.Engine //定义引擎全局变量
var AllTables []interface{}

// 用户表
type Usr struct {
	Id       string `xorm:"varchar(36) pk index" json:"id,omitempty"`              // 用户ID
	RecordId string `xorm:"varchar(25) notnull unique" json:"record_id,omitempty"` // 档案号
	Name     string `xorm:"varchar(50)" json:"name,omitempty"`                     // 姓名
	Phone    string `xorm:"varchar(16) notnull unique" json:"phone,omitempty"`     // 手机号码
	CardNum  string `xorm:"varchar(20) notnull unique" json:"card_num,omitempty"`  // 身份证号
	Job      string `xorm:"varchar(35)" json:"job,omitempty"`                      // 职业
	//
	//History []string `json:"history,omitempty"` // 个人病史
	//
	//BaseHealthInfo map[string]interface{} `json:"base_health_info,omitempty"` // 基础健康信息
	//
	//ZyYx map[string]interface{} `json:"zy_yx,omitempty"` // 中医眼象辨识
	//ZyTz map[string]interface{} `json:"zy_tz,omitempty"` // 中医体质辨识
	//
	//WenJuan map[string]interface{} `json:"wen_juan,omitempty"` // 问卷
	//
	//Time int64 `xorm:"created" json:"time,omitempty"` // 用户创建时间
}

func (this *Usr) BeforeInsert() {
	fmt.Println("正在处理 BeforeInsert")
	if this.Id == "" {
		id, _ := uuid.NewV1()
		this.Id = id.String()
	}
	if this.Name == "" {
		this.Name = "未设置"
	}
}

// 启动程序后就执行
func init() {
	// 创建 ORM 引擎与数据库
	var err error
	X, err = xorm.NewEngine("mysql", "root:password@tcp(127.0.0.1:3306)/admin")
	if err != nil {
		fmt.Println("引擎创建出错", err)
	}
	err = X.Ping()
	if err == nil {
		fmt.Printf("连接成功\n")
	}
	X.SetConnMaxLifetime(300 * time.Second)
	X.SetMaxIdleConns(40)
	X.SetMaxOpenConns(100)

	// 同步结构体与数据表,和python的 migrate 一样同步数据库
	AllTables = append(AllTables,
		new(Usr),
		new(Role),
		new(User),
		new(UserRole),
	)
	err = X.Sync2(AllTables...)
	if err != nil {
		fmt.Printf("Fail to sync database: %v\n", err)
	}
	cacher := xorm.NewLRUCacher2(xorm.NewMemoryStore(), time.Duration(300)*time.Second, 5000)
	X.MapCacher(new(Usr), cacher)
}
