package controllers

import (
	"github.com/astaxie/beego"
	"liteblog/models"
	"liteblog/syserrors"
	"errors"
	"github.com/satori/go.uuid"
)

const SESSION_USER_KEY = "SESSION_USER_KEY1"

type BaseController struct {
	beego.Controller

	IsLogin bool
	User    models.UserModel
	Dao     *models.DB
}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// 接口设计
type NestPreparer interface {
	NestPrepare()
}

func (ctx *BaseController) Prepare() {

	IsUserLogin(ctx)

	myPath := ctx.Ctx.Request.RequestURI
	ctx.Dao = models.NewDB()

	ctx.Data["Path"] = myPath
	//log.Print(">>>> BaseController <<<< --->" + myPath)
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

// 判断用户是否登录
func IsUserLogin(ctx *BaseController) {
	ctx.IsLogin = false
	tu := ctx.GetSession(SESSION_USER_KEY)
	if tu != nil {
		if user ,ok := tu.(models.UserModel) ; ok {
			ctx.User = user
			ctx.IsLogin = true
			ctx.Data["User"] = user

		}
	}

	ctx.Data["IsLogin"] = ctx.IsLogin
}




// 验证工具
func (ctx *BaseController) MustLogin() {
	if !ctx.IsLogin {
		ctx.Abort500(syserrors.NoUserError{})
	}
}

func (c *BaseController) GetMustString(key string, msg string) string {
	email := c.GetString(key, "")
	if len(email) == 0 {
		c.Abort500(errors.New(msg))
	}
	return email
}

func (ctx *BaseController) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}

func (ctx *BaseController) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	ctx.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}

func (this *BaseController) UUID() string {
	u,err:=uuid.NewV4()
	if err!=nil{
		this.Abort500(syserrors.NewError("系统错误",err))
	}
	return u.String()
}