package verify

import "github.com/gogf/gf/net/ghttp"

func Verify(r *ghttp.Request)  {
	r.Response.WriteExit("ok")
}
