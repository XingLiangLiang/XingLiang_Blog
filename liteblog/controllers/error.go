package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404()  {
	c.Data["content"] = "page not found"
	c.TplName = "error/404.html"
}
