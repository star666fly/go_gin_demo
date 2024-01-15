package main

import (
	_ "go_demo/docs"
)

// @title			接囗文档
// @version		1.0
// @description	demo项目
// @microservice	https://github.com/xxxx
// @contact.name	demo
// @contact.email	xxxx@126.com
func main() {
	r := SetRouter()
	r.Run(":8081")
}
