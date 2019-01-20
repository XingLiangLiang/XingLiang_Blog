package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type DB struct {
	db *gorm.DB
}

func (db *DB) Begin() {
	db.db = db.db.Begin()
}

func (db *DB) Rollback() {
	db.db = db.db.Rollback()
}

func (db *DB) Commit() {
	db.db = db.db.Commit()
}

func NewDB() *DB {
	return &DB{db: dbValue}
}

var (
	dbValue *gorm.DB
)

func init() {

	if err := os.MkdirAll("data", 0777); err != nil {
		panic("failed to mkdir data," + err.Error())
	}

	var err error
	dbValue, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/beego_blog?charset=utf8&parseTime=true")

	//defer dbValue.Close()

	if err != nil {
		panic("failed to connect database =" + err.Error())
	}

	dbValue.AutoMigrate(&UserModel{},&NoteModel{})

	insertManager(dbValue)
}


func insertManager(db *gorm.DB) {
	var count int

	if err := db.Model(&UserModel{}).Count(&count).Error; err == nil && count == 0 {
		// 新增管理员
		db.Create(
			&UserModel{
				Name:   "admin",
				Email:  "admin@126.com",
				Avatar: "/static/images/info-img.png",
				Pwd:    "123456",
				Role:   0,
			})
	}

}

