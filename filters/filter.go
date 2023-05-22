package filters

import (
	"github.com/beego/beego/v2/server/web/context"
)

var FilterUser = func(ctx *context.Context) {

	user := ctx.Input.Session("user")
	if user == nil {
		ctx.Redirect(302, "/login")
	}

}
