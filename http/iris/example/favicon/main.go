// main.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	app.Favicon("./static/favicons/iris_favicon_32_32.ico")

	app.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, `You should see the favicon now the side of your browser, if not, please refrsh or clear the browser's cache.`)
	})

	app.Listen(":8090")
}
