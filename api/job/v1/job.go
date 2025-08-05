package v1

import "github.com/gogf/gf/v2/frame/g"

type SendMailJobReq struct {
	g.Meta `path:"/sendEmailJob" method:"post" tags:"job" summary:"send email job info"`
	Email  string `v:"required|email" dc:"email"`
}

type SendMailJobRes struct{}
