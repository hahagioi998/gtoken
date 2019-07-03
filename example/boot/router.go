package boot

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"gtoken/gtoken"
	"gtoken/utils/resp"
)

/*
绑定业务路由
*/
func bindRouter() {

	s := g.Server()
	// 调试路由
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.WriteJson(resp.Succ("hello"))
	})
	s.BindHandler("/system/user", func(r *ghttp.Request) {
		r.Response.WriteJson(resp.Succ("system user"))
	})

	loginFunc := Login
	// 启动gtoken
	gtoken := &gtoken.GfToken{
		//Timeout:         10 * 1000,
		CacheMode:       g.Config().GetInt8("cache-mode"),
		LoginPath:       "/login",
		LoginBeforeFunc: loginFunc,
		LogoutPath:      "/user/logout",
		AuthPaths:       g.SliceStr{"/user/*", "/system/*"},
	}
	gtoken.Start()

}

/*
统一路由注册
*/
func initRouter() {

	s := g.Server()

	// 绑定路由
	bindRouter()

	// 首页
	s.BindHandler("/", func(r *ghttp.Request) {
		content, err := g.View().Parse("index.html", map[string]interface{}{
			"id":    1,
			"name":  "GTOKEN",
			"title": g.Config().GetString("setting.title"),
		})
		if err != nil {
			glog.Error(err)
		}
		r.Response.Write(content)

	})

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

}

func Login(r *ghttp.Request) (string, interface{}) {
	username := r.GetPostString("username")
	passwd := r.GetPostString("passwd")

	if username == "" || passwd == "" {
		r.Response.WriteJson(resp.Fail("账号或密码错误."))
		r.ExitAll()
	}

	return username, ""
}