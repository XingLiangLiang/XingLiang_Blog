package controllers

import (
	"liteblog/syserrors"
	"github.com/jinzhu/gorm"
	"time"
	"bytes"
	"liteblog/models"
	"github.com/PuerkitoBio/goquery"
)

type NoteController struct {
	BaseController

}

func (this *NoteController) NestPrepare() {
	//
	this.MustLogin()
	if this.User.Role != 0 {
		this.Abort500(syserrors.NewError("您没有该权限",nil))
	}

}

// @router /new [get]
func (this *NoteController) NewNote()  {
	this.Data["key"] =this.UUID()
	this.TplName = "note_new.html"
}

// @router /save/:key [post]
func (this *NoteController) Save()  {
	noteKey := this.Ctx.Input.Param(":key")
	title := this.GetMustString("title", "标题不能为空！")
	content := this.GetMustString("content", "内容不能为空！")
	files := this.GetString("files", "")
	summary, _ := getSummary(content)
	note, err := this.Dao.QueryNoteByKeyAndUserId(noteKey, int(this.User.ID))
	var n models.NoteModel
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			this.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = models.NoteModel{
			NoteKey:     noteKey,
			Summary: summary,
			Title:   title,
			Files:   files,
			Content: content,
			UserID:  int(this.User.ID),
		}
	} else {
		n = note
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.Files = files
		n.UpdatedAt = time.Now()
	}
	if err := this.Dao.SaveNote(&n); err != nil {
		this.Abort500(syserrors.NewError("保存失败！", err))
	}
	this.JSONOk("成功", "/details/"+noteKey)

}

func getSummary(content string) (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(content))
	doc, err := goquery.NewDocumentFromReader(&buf)
	if err != nil {
		return "", err
	}
	str := doc.Find("body").Text()
	strRune := []rune(str)
	if len(strRune) > 400 {
		strRune = strRune[:400]
	}
	return string(strRune) + "...", nil
}
