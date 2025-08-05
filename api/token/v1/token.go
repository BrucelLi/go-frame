package v1

import "github.com/gogf/gf/v2/frame/g"

type GetTokenReq struct {
	g.Meta `path:"/user/token" method:"get" tags:"User" summary:"Get users token"`
	UserID int64 `v:"required|max-length:50" dc:"user id"`
}

type GetTokenRes struct {
	Token string `json:"token" dc:"user token"`
}

type LoginReq struct {
}
