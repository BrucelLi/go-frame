// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta       `orm:"table:users, do:true"`
	Id           interface{} // 用户唯一标识，自增主键
	Username     interface{} // 用户名，最长50字符
	Email        interface{} // 用户邮箱，唯一，最长100字符
	PasswordHash interface{} // 用户密码的哈希值
	IsActive     interface{} // 用户是否激活，默认为激活，true：激活；false: 未激活
	CreatedAt    *gtime.Time // 创建时间，默认为当前时间
	UpdatedAt    *gtime.Time // 更新时间，默认为当前时间
}
