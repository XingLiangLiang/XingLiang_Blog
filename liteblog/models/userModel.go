package models

import (
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model

	Name   string `gorm:"unique_index"`
	Email  string `gorm:"unique_index"`
	Avatar string
	Pwd    string
	Role   int `gorm:"default:1"` // 0 管理员 1正常用户
}

func (db *DB) QueryUserByEmailAndPassword(email, password string) (user UserModel, err error) {
	return user, db.db.Model(UserModel{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error
}

func (db *DB) QueryUserByName(name string) (user UserModel, err error) {
	return user, db.db.Where("name = ?", name).Take(&user).Error

}

func (db *DB) QueryUserByEmail(email string) (user UserModel, err error) {
	return user, db.db.Where("email = ?", email).Take(&user).Error
}

func SaveUser(user *UserModel) error {
	return dbValue.Create(user).Error
}
