package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
    //beego.Router("/", &controllers.IndexController{})
	//beego.Router("/message", &controllers.IndexController{} , "*:GetMessage")
	//beego.Router("/about", &controllers.IndexController{},"*:GetAbout")
    //beego.ErrorController(&controllers.ErrorController{})
	//beego.Include(&controllers.IndexController{})
	//

	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
	)

	beego.AddNamespace(
		beego.NewNamespace(
			"note",
			beego.NSInclude(&controllers.NoteController{}),
		),
		)
}


