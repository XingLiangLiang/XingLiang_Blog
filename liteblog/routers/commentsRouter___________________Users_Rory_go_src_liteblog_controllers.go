package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetAbout",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetMessage",
			Router: `/message`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetReg",
			Router: `/reg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:IndexController"] = append(beego.GlobalControllerRouter["liteblog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "NewNote",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:NoteController"] = append(beego.GlobalControllerRouter["liteblog/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/save/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["liteblog/controllers:UserController"] = append(beego.GlobalControllerRouter["liteblog/controllers:UserController"],
		beego.ControllerComments{
			Method: "Reg",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
