package user

import (
	"context"
	"demo/api/user/v1"
	"demo/internal/model/do"
	"demo/internal/repository"
	"demo/internal/service"
	"demo/utility"
)

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

type sUser struct {
}

func (s *sUser) AddUser(ctx context.Context, req *v1.CreateReq) (res int64, err error) {
	encryptPassword, err := utility.PasswordEncrypt(req.Password)
	if err != nil {
		return 0, err
	}
	var data = do.Users{
		Username:     req.Username,
		IsActive:     v1.StatusOK,
		Email:        req.Email,
		PasswordHash: encryptPassword,
	}

	userID, err := repository.AddUser(ctx, data)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
