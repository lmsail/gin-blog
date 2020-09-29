package main

import (
	"gin-lmsail/app/Models"
	"gin-lmsail/routes"
)

func main() {
	blog := routes.InitWebRouter()
	defer Models.DB.Close()
	_ = blog.Run(":8888")
}
