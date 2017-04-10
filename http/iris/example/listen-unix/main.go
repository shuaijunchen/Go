// main.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

const host = "127.0.0.1:443"

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from the server")
	})

	app.ListenUNIX("/tmp/srv.sock", 0666)
}
