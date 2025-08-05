package servers

import (
	"context"
	"demo/internal/jobs/tasks"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hibiken/asynq"
)

func StartJobsWorker(ctx context.Context) error {
	redisAddr := g.Cfg().MustGet(ctx, "jobs.redis.address").String()
	concurrency := g.Cfg().MustGet(ctx, "jobs.redis.concurrency").Int()
	queues := g.Cfg().MustGet(ctx, "jobs.redis.queues").Int()
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: concurrency,
			Queues: map[string]int{
				"default": queues,
			},
		},
	)

	mux := asynq.NewServeMux()

	// 注册任务处理器
	mux.HandleFunc(tasks.TaskSendEmail, tasks.HandleSendEmail)
	mux.HandleFunc(tasks.TaskLogEvent, tasks.HandleLogEvent)

	g.Log().Info(ctx, "starting jobs worker")
	return srv.Run(mux)
}
