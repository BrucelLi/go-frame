package v1

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta         `path:"/user" method:"post" tags:"User" summary:"Create user"`
	Username       string `v:"required#ReuiredUserName|max-length:50" dc:"user name"`
	Email          string `v:"required|email" dc:"user email"`
	Password       string `v:"required|password3|ci|same:RepeatPassword" dc:"user password"`
	RepeatPassword string `v:"required|password3" dc:"user RepeatPassword"`
}
type CreateRes struct {
	ID int64 `json:"id" dc:"user id"`
}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"Delete user"`
	ID     int64 `v:"required" dc:"user id"`
}
type DeleteRes struct{}

// Status marks user status.
type Status bool

const (
	StatusOK       Status = true  // User is OK.
	StatusDisabled Status = false // User is disabled.
)

type UpdateReq struct {
	g.Meta   `path:"/user/{id}" method:"put" tags:"User" summary:"Update user"`
	ID       int64   `v:"required" dc:"user id"`
	Username string  `v:"required|max-length:50" dc:"user name"`
	Email    string  `v:"required|email" dc:"user email"`
	IsActive *Status `v:"bool" dc:"user status"`
}
type UpdateRes struct{}

type GetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get one user"`
	ID     int64 `v:"required" dc:"user id"`
}
type GetOneRes struct {
	*entity.Users `dc:"user"`
}

type GetListReq struct {
	g.Meta   `path:"/user" method:"get" tags:"User" summary:"Get users"`
	IsActive *Status `v:"bool" dc:"user status"`
}
type GetListRes struct {
	List []*entity.Users `json:"list" dc:"user list"`
}
