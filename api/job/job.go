// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package job

import (
	"context"

	"demo/api/job/v1"
)

type IJobV1 interface {
	SendMailJob(ctx context.Context, req *v1.SendMailJobReq) (res *v1.SendMailJobRes, err error)
}
