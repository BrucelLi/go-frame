package token

import (
	"context"
	"demo/utility"

	"demo/api/token/v1"
)

func (c *ControllerV1) GetToken(_ context.Context, req *v1.GetTokenReq) (res *v1.GetTokenRes, err error) {
	token, err := utility.CreateToken(req.UserID)
	if err != nil {
		return nil, err
	}
	res = &v1.GetTokenRes{
		Token: token,
	}
	return
}
