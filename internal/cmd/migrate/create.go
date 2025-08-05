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
		Name:  "create",
		Usage: "migrate create -name=<migration_name>",
		Brief: "Create a new migration file",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			name := parser.GetOpt("name").String()
			if name == "" {
				return gerror.New("missing migration name. Usage: migrate create -name=<migration_name>")
			}

			dir, err := g.Cfg().Get(ctx, "migration.dir")
			if err != nil {
				return gerror.Wrap(err, "failed to get migration.dir from config")
			}
			migrationsDir := dir.String()

			g.Log().Infof(ctx, "Creating migration: %s", name)
			tmpCmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", migrationsDir, "-seq", name)

			out, err := tmpCmd.CombinedOutput()
			if err != nil {
				return gerror.Wrapf(err, "failed to create migration file: %s", string(out))
			}

			g.Log().Infof(ctx, "Migration created successfully:\n%s", string(out))
			return nil
		},
	})
	if err != nil {
		panic(err)
	}
}
