package jobs

import (
	"context"
	"demo/internal/jobs/servers"
	"github.com/gogf/gf/v2/os/gcmd"
)

var Jobs = gcmd.Command{
	Name:  "jobs",
	Usage: "jobs run",
	Brief: "jobs run tool for redis",
	Func: func(ctx context.Context, _ *gcmd.Parser) error {
		// run jobs worker
		err := servers.StartJobsWorker(ctx)
		if err != nil {
			panic(err)
		}
		return nil
	},
}
