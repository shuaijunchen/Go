// main.go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())
	app.Adapt(view.HTML("./templates", ".html"))

	app.Get("/upload", func(ctx *iris.Context) {
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		ctx.Render("upload_form.html", token)
	})

	app.Post("/upload", iris.LimitRequestBodySize(10<<20), func(ctx *iris.Context) {
		file, info, err := ctx.FormFile("uploadfile")
		if err != nil {
			// ctx.HTML(iris.StatusInternalServerError, "Error while uploading:<b>"+err.Error()+"</b>")
			myerr(err, ctx)
			return
		}

		defer file.Close()
		fname := info.Filename

		out, err := os.OpenFile("./uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			myerr(err, ctx)
			return
		}

		defer out.Close()
		io.Copy(out, file)
	})

	app.Listen(":8080")
}

func myerr(err error, ctx *iris.Context) {
	ctx.HTML(iris.StatusInternalServerError, "Error while uploading:<b>"+err.Error()+"</b>")
}
