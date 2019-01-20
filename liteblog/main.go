package main

import (
	_ "liteblog/models" //初始化数据库
	_ "liteblog/routers" // 初始化路由
	"github.com/astaxie/beego"
	"strings"
	"encoding/gob"
	"liteblog/models"
	"fmt"
)

func main() {
	initSession()
	initTemplate()
	beego.Run()
}


func initSession() {
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(models.UserModel{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}


func initTemplate() {

	// 在template 中添加 比较方法
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})

	beego.AddFuncMap("add", func(x, y int) int {
		return x + y
	})

	beego.AddFuncMap("eq", func(x, y interface{}) bool {
		s1 := fmt.Sprintf("%v", x)
		s2 := fmt.Sprintf("%v", y)
		return strings.Compare(s1, s2) == 0
	})
}

