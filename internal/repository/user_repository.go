package repository

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/do"
)

func AddUser(ctx context.Context, data do.Users) (res int64, err error) {
	insertID, err := dao.Users.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return insertID, nil
}
