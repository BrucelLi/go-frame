package user

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/do"

	"demo/api/user/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	_, err = dao.Users.Ctx(ctx).Data(do.Users{
		Username: req.Username,
		IsActive: req.IsActive,
		Email:    req.Email,
	}).WherePri(req.ID).Update()
	return
}
