// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for the table users.
type UsersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UsersColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UsersColumns defines and stores column names for the table users.
type UsersColumns struct {
	Id           string // 用户唯一标识，自增主键
	Username     string // 用户名，最长50字符
	Email        string // 用户邮箱，唯一，最长100字符
	PasswordHash string // 用户密码的哈希值
	IsActive     string // 用户是否激活，默认为激活，true：激活；false: 未激活
	CreatedAt    string // 创建时间，默认为当前时间
	UpdatedAt    string // 更新时间，默认为当前时间
}

// usersColumns holds the columns for the table users.
var usersColumns = UsersColumns{
	Id:           "id",
	Username:     "username",
	Email:        "email",
	PasswordHash: "password_hash",
	IsActive:     "is_active",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao(handlers ...gdb.ModelHandler) *UsersDao {
	return &UsersDao{
		group:    "default",
		table:    "users",
		columns:  usersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
