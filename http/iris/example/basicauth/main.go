// main.go
package main

import (
	"time"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/basicauth"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	authConfig := basicauth.Config{
		Users:      map[string]string{"admin": "123456", "zhangsan": "pwd"},
		Realm:      "Authorization Required", // defaults to "Authorization Required"
		ContextKey: "mycustomkey",            // defaults to "user"
		Expires:    time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	app.Get("/", func(ctx *iris.Context) {
		ctx.Redirect("/admin")
	})

	needAuth := app.Party("/admin", authentication)
	{
		needAuth.Get("/", func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey")
			ctx.Writef("Hello authenticated user: %s from: %s ", username, ctx.Path())
		})

		needAuth.Get("/profile", func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey")
			ctx.Writef("Hello authenticated user: %s from: %s ", username, ctx.Path())
		})

		needAuth.Get("/settings", func(ctx *iris.Context) {
			username := authConfig.User(ctx)
			ctx.Writef("Hello authenticated user: %s from: %s ", username, ctx.Path())
		})
	}

	app.Listen(":8080")
}
