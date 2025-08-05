package user

import (
	"context"
	"demo/internal/service"

	"demo/api/user/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	insertID, err := service.User().AddUser(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{
		ID: insertID,
	}
	return
}
