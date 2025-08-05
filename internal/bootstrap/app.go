package bootstrap

import (
	"demo/internal/cmd"
	"errors"
	// 引入 PostgreSQL 驱动，触发 init 注册
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	// 引入 Redis 驱动
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("database ping master err:" + err.Error())
	}
	return nil
}

func init() {
	ctx := gctx.GetInitCtx()
	var err error
	// default i18n
	lang := g.Cfg().MustGet(ctx, "local.lang").String()
	g.I18n().SetLanguage(lang)

	// test db connect
	err = connDb()
	if err != nil {
		panic(err)
	}
	// set db TIME ZONE
	_, err = g.DB().Exec(ctx, "SET TIME ZONE 'Asia/Shanghai'")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
