// main.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type Company struct {
	Name  string
	City  string
	Other string
}

func MyHandler(ctx *iris.Context) {
	c := &Company{}
	if err := ctx.ReadJSON(c); err != nil {
		ctx.Log(iris.ProdMode, err.Error())
		return
	}

	ctx.Writef("Company: %#v\n", c)
}

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.Post("/", MyHandler)
	app.Listen(":8080")
}
