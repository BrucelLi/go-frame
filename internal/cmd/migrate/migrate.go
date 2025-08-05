package migrate

import "github.com/gogf/gf/v2/os/gcmd"

var Migrate = gcmd.Command{
	Name:  "migrate",
	Usage: "migrate [subcommand]",
	Brief: "Migration tool for database",
}
