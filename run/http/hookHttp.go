package http

import (
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestCookie"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpRequestFormValue"
	_ "github.com/HXSecurity/DongTai-agent-go/core/http/httpServeHTTP"
	"github.com/HXSecurity/DongTai-agent-go/global"
	"github.com/HXSecurity/DongTai-agent-go/hook"
)

func init() {
	h := new(hook.Http)
	global.AllHooks = append(global.AllHooks, h)
	h.HookAll()
}
