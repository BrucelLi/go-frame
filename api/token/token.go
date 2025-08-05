// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package token

import (
	"context"

	"demo/api/token/v1"
)

type ITokenV1 interface {
	GetToken(ctx context.Context, req *v1.GetTokenReq) (res *v1.GetTokenRes, err error)
}
