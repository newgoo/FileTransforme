package main

import (
	"net/http"
	"os"
	"path/filepath"
	"flag"
	"fmt"
)

func main() {
	var (
		err error
		tip = `当前路径为:%s  \r\n
当前端口号为:%s \r\n
[CTRL + C]停止服务 \r\n`
	)

	port := flag.String("port", "8080", "http listen port")
	path := flag.String("path", "/", "path")
	flag.Parse()
	if *path == "/" {
		*path = pathVal()
	}
	fmt.Println(fmt.Sprintf(tip, *path, *port))

	err = http.ListenAndServe(fmt.Sprintf(":%s", *port), http.FileServer(http.Dir(*path)))
	if err != nil {
		fmt.Sprintln("监听端口失败,当前端口被占用或路径不正确,", err)
	}

}

func pathVal() string {
	fmt.Sprintln("路径为更目录或者未输入目录!自动获取当前目录")
	path, err := getPath()
	if err != nil {
		fmt.Sprintln("获取当前路径失败,error:", err)
	}
	return path
}

func getPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
