package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "yistore/models"
	_ "yistore/routers"
)

func main() {
	//坑
	beego.InsertFilter("http://127.0.0.1:8080", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Accept", "Authorization", "Access-Control-Allow-Origin",
									"Access-Control-Allow-Headers", "Access-Control-Allow-Credentials", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Run()
}
