package client

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/hibiken/asynq"
)

var clientInstance *asynq.Client

func GetJobsClient() *asynq.Client {
	if clientInstance == nil {
		ctx := gctx.GetInitCtx()
		redisAddr := g.Cfg().MustGet(ctx, "jobs.redis.address").String()
		clientInstance = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	}
	return clientInstance
}
