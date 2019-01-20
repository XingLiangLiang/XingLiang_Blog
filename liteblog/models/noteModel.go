package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type NoteModel struct {
	gorm.Model

	NoteKey string `gorm:"unique_index;not null;"`
	UserID  int
	User    UserModel
	Title   string
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Files   string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func (db *DB) QueryNoteByKeyAndUserId(key string, userid int) (note NoteModel, err error) {
	return note, db.db.Model(&NoteModel{}).Where("`note_key` = ? and user_id = ?", key, userid).Take(&note).Error
}

func (db *DB) QueryNoteByKey(key string) (note NoteModel, err error) {
	return note, db.db.Model(&NoteModel{}).Where("`note_key` = ? ", key).Take(&note).Error
}

func (db *DB) SaveNote(n *NoteModel) error {
	return db.db.Save(n).Error
}

func (db *DB) QueryAllNotes() (notes []*NoteModel, err error) {
	return notes, db.db.Find(&notes).Error
}

//func (db *DB) QueryNotesByPage(page, limit int, title string) (notes []*NoteModel, err error) {
//	return notes, db.db.Offset((page - 1) * limit).Limit(limit).Find(&notes).Error
//}

func (db *DB) QueryNotesByPage(page, limit int, title string) (note []*NoteModel, err error) {
	return note, db.db.Model(&NoteModel{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1) * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}

func (db *DB) QueryNotesCount(title string) (cnt int, err error) {
	return cnt, db.db.Model(&NoteModel{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset(-1).Limit(-1).Count(&cnt).Error
}
