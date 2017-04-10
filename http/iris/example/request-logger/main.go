// main.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"
)

func main() {
	app := iris.New()

	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	customLogger := logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
	})

	app.Use(customLogger)

	app.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello")
	})

	app.Get("/1", func(ctx *iris.Context) {
		ctx.Writef("hello")
	})

	app.Get("/2", func(ctx *iris.Context) {
		ctx.Writef("hello")
	})

	errorLogger := logger.New()
	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Writef("My Custom 404 error page")
	})

	app.Listen(":8080")
}
