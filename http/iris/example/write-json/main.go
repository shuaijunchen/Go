// main.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	app.Post("/decode", func(ctx *iris.Context) {
		var user User
		ctx.ReadJSON(&user)

		ctx.Writef("%s %s is %d year old!", user.Firstname, user.Lastname, user.Age)
	})

	app.Get("/encode", func(ctx *iris.Context) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       21,
		}

		ctx.JSON(iris.StatusOK, peter)
	})

	app.Listen(":8080")
}
