// main.go
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"log"
	"os"
)

var myLogFile *os.File

func init() {
	f, err := os.Create("logs.txt")
	if err != nil {
		panic(err)
	}

	myLogFile = f
}

func myFileLogger() iris.LoggerPolicy {
	myLogger := log.New(myLogFile, "", log.LstdFlags)

	return func(mode iris.LogMode, message string) {
		if mode == iris.ProdMode {
			myLogger.Println(message)
		}
	}
}

func main() {
	defer func() {
		if err := myLogFile.Close(); err != nil {
			panic(err)
		}
	}()

	app := iris.New()
	app.Adapt(myFileLogger())
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		app.Log(iris.ProdMode, "You have requested: http://localhost:8080/"+ctx.Path())
		ctx.Writef("hello")
	})

	app.Listen(":8080")
}
