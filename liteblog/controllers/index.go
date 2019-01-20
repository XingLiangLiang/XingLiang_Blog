package controllers

import (
	"log"
	"liteblog/syserrors"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Get() {
	log.Println("-----Get-----")
	//notes, err := c.Dao.QueryAllNotes()
	limit := 5
	page , err := c.GetInt("page" , 1)
	if err != nil || page < 1 {
		page = 1
	}

	title := c.GetString("title" ,"")

	notes,err := c.Dao.QueryNotesByPage(page , limit , title)
	if err != nil {
		log.Println("-----查询错误-----")
		c.Abort500(syserrors.NewError("查询错误", nil))
	} else {
		log.Println("notes size = %v",len(notes))
	}
	if notes != nil {
		c.Data["notes"] = notes
	}

	totalPage := 0
	totalCount,_ := c.Dao.QueryNotesCount(title)
	if totalCount % limit == 0 {
		totalPage = totalCount/limit
	} else {
		totalPage = totalCount/limit + 1
	}

	c.Data["totpage"] = totalPage
	c.Data["page"] = page
	c.Data["title"] = title

	c.TplName = "index.html"
}

// @router /message [get]
func (c *IndexController) GetMessage() {
	log.Println("-----message-----")
	c.TplName = "message.html"
}

// @router /about [get]
func (c *IndexController) GetAbout() {
	log.Println("-----about-----")
	c.TplName = "about.html"
}

// @router /user [get]
func (c *IndexController) GetUser() {
	c.TplName = "user.html"
}

// @router /reg [get]
func (c *IndexController) GetReg() {
	c.TplName = "reg.html"
}
