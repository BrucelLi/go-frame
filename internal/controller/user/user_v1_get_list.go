package user

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/do"

	"demo/api/user/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.Users.Ctx(ctx).Where(do.Users{
		IsActive: req.IsActive,
	}).Scan(&res.List)
	return
}
