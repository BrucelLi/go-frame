package middleware

import (
	"demo/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

func I18nMiddleware(r *ghttp.Request) {
	langHeader := r.Header.Get("Accept-Language")
	if langHeader == "" {
		langHeader = "zh-CN"
	}

	if lang, ok := utility.SupportedLang[langHeader]; ok {
		utility.SetLocaleLang(lang)
	} else {
		utility.SetLocaleLang(utility.LangCn)
	}

	r.Middleware.Next()
}
