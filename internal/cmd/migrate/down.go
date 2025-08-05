// cmd/migrate/down.go
package migrate

import (
	"context"
	"os/exec"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

func init() {
	err := Migrate.AddCommand(&gcmd.Command{
		Name:  "down",
		Usage: "migrate down",
		Brief: "Roll back the last migration",
		Func: func(ctx context.Context, _ *gcmd.Parser) error {
			g.Log().Infof(ctx, "Migration down")

			dsn, _ := g.Cfg().Get(ctx, "database.default.link")
			g.Log().Infof(ctx, "Migration down dsn: %s", dsn.String())

			dir, err := g.Cfg().Get(ctx, "migration.dir")
			if err != nil {
				return gerror.Wrap(err, "failed to get migration.dir from config")
			}
			migrationsDir := dir.String()
			g.Log().Infof(ctx, "Migration down migrationsDir: %s", migrationsDir)

			tmpCmd := exec.Command("migrate",
				"-path", migrationsDir,
				"-database", dsn.String(),
				"down", "1",
			)

			out, err := tmpCmd.CombinedOutput()
			if err != nil {
				return gerror.Wrapf(err, "failed migration down: %s", string(out))
			}

			g.Log().Infof(ctx, "Migration down successfully:\n%s", string(out))
			return nil
		},
	})
	if err != nil {
		panic(err)
	}
}
