// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id           int64       `json:"id"           orm:"id"            description:"用户唯一标识，自增主键"`                     // 用户唯一标识，自增主键
	Username     string      `json:"username"     orm:"username"      description:"用户名，最长50字符"`                      // 用户名，最长50字符
	Email        string      `json:"email"        orm:"email"         description:"用户邮箱，唯一，最长100字符"`                 // 用户邮箱，唯一，最长100字符
	PasswordHash string      `json:"passwordHash" orm:"password_hash" description:"用户密码的哈希值"`                        // 用户密码的哈希值
	IsActive     bool        `json:"isActive"     orm:"is_active"     description:"用户是否激活，默认为激活，true：激活；false: 未激活"` // 用户是否激活，默认为激活，true：激活；false: 未激活
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间，默认为当前时间"`                    // 创建时间，默认为当前时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间，默认为当前时间"`                    // 更新时间，默认为当前时间
}
