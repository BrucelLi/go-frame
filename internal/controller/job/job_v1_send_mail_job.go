package job

import (
	"context"
	"demo/api/job/v1"
	"demo/internal/jobs/tasks"
)

func (c *ControllerV1) SendMailJob(_ context.Context, req *v1.SendMailJobReq) (res *v1.SendMailJobRes, err error) {
	err = tasks.EnqueueSendEmail(req.Email, "", "")
	if err != nil {
		return nil, err
	}
	return
}
