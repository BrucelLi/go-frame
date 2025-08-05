package cmd

import (
	"context"
	"demo/internal/cmd/jobs"
	"demo/internal/cmd/migrate"
	"demo/internal/router"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http servers",
		Func: func(_ context.Context, _ *gcmd.Parser) (err error) {
			s := g.Server()
			router.RegisterRoutes(s)
			s.Run()
			return nil
		},
	}
)

func init() {
	if err := Main.AddCommand(&migrate.Migrate); err != nil {
		panic(err)
	}
	if err := Main.AddCommand(&jobs.Jobs); err != nil {
		panic(err)
	}
}
