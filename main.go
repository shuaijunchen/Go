package main

import (
	// "fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("upload.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		// 获取上传信息
		f, h, err := r.FormFile("file")
		log.Println(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 获取文件名
		filename := h.Filename
		log.Println(filename)
		defer f.Close() // 关闭

		// 文件路径
		path := UPLOAD_DIR + "/" + filename
		log.Println(path) //文件路径

		// 获取当前目录
		dir, _ := os.Getwd()
		log.Println(dir) //当前文件目录

		// a, err := os.Stat(UPLOAD_DIR) // 检查文件是否存在
		isV := isExists(UPLOAD_DIR)
		if !isV {
			// 创建文件夹
			e := os.Mkdir(dir+"/"+UPLOAD_DIR, os.ModePerm)
			if e != nil {
				log.Println(e)
			}
			log.Println("创建成功")
		}

		// 创建文件
		t, err := os.Create(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close() //关闭

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 重定向到上传成功页面
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// 获取文件信息
	fileId := r.FormValue("id")
	// 拼接文件路径
	filePath := UPLOAD_DIR + "/" + fileId
	if exists := isExists(filePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "file")
	http.ServeFile(w, r, filePath)
}

// true:文件存在，false:文件不存在
func isExists(path string) bool {
	// 检查文件是否存在
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func main() {
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	// 监听端口
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
