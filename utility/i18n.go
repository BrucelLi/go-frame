package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

type Lang string

const (
	LangEn Lang = "en"    // en
	LangCn Lang = "zh-CN" // zh-CN
)

var (
	SupportedLang = map[string]Lang{
		"en":    LangEn,
		"zh-CN": LangCn,
	}
)

const localeCacheKey = "localeLang"

func SetLocaleLang(lang Lang) {
	err := gcache.Set(gctx.GetInitCtx(), localeCacheKey, lang, 0)
	if err != nil {
		panic(err)
	}
}

func GetLocaleLang() Lang {
	value, err := gcache.Get(gctx.GetInitCtx(), localeCacheKey)
	if err != nil {
		panic(err)
	}
	langStr := value.String()
	if lang, ok := SupportedLang[langStr]; ok {
		return lang
	}
	return LangCn
}

func GetLocaleT(key string) string {
	lang := GetLocaleLang()
	ctx := gi18n.WithLanguage(context.TODO(), string(lang))
	title := g.I18n().T(ctx, key)
	return title
}

func GetLocaleTf(format string, values ...interface{}) string {
	lang := GetLocaleLang()
	ctx := gi18n.WithLanguage(context.TODO(), string(lang))
	return g.I18n().Tf(ctx, format, values...)
}
