package dao

import (
	"blog/dal/model"
	"github.com/jinzhu/gorm"
)

type blogAboutDao struct{}

func (*blogAboutDao) Query(db *gorm.DB) *model.BlogAboutModel {
	about := model.BlogAboutModel{}
	db.First(&about)
	return &about
}

func (*blogAboutDao) Update(db *gorm.DB, record *model.BlogAboutModel) error {
	record.AdminPwd = nil
	record.ViewNum = nil
	return db.Model(&model.BlogAboutModel{}).Updates(record).Error
}

func (*blogAboutDao) AddViewNum(db *gorm.DB) error {
	return db.Exec("update `Blog_about` set `view_num` = `view_num` + 1 where `id` = 1").Error
}

func (*blogAboutDao) CheckAdmin(db *gorm.DB, pwd string) bool {
	total := -1
	db.Model(&model.BlogAboutModel{}).Where("admin_pwd = ?", pwd).Count(&total)
	return total > 0
}
