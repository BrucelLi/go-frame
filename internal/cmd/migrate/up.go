package migrate

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"os/exec"

	"github.com/gogf/gf/v2/os/gcmd"
)

func init() {
	err := Migrate.AddCommand(&gcmd.Command{
		Name:  "up",
		Usage: "migrate up",
		Brief: "Apply all up migrations",
		Func: func(ctx context.Context, _ *gcmd.Parser) (err error) {
			g.Log().Infof(ctx, "Migration up")
			dsn, _ := g.Cfg().Get(ctx, "database.default.link")
			g.Log().Infof(ctx, "Migration up dsn %s", dsn.String())
			dir, err := g.Cfg().Get(ctx, "migration.dir")
			if err != nil {
				return gerror.Wrap(err, "failed to get migration.dir from config")
			}
			migrationsDir := dir.String()
			g.Log().Infof(ctx, "Migration up migrationsDir %s", migrationsDir)
			tmpCmd := exec.Command("migrate",
				"-path", migrationsDir,
				"-database", dsn.String(),
				"up",
			)
			out, err := tmpCmd.CombinedOutput()
			if err != nil {
				return gerror.Wrapf(err, "failed migration up: %s", string(out))
			}

			g.Log().Infof(ctx, "Migration up successfully:\n%s", string(out))
			return nil
		},
	})
	if err != nil {
		return
	}
}
