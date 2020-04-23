package dao

import (
	"blog/dal/model"
	"blog/info/query"
	"github.com/jinzhu/gorm"
)

type blogMessageDao struct{}

func (*blogMessageDao) QueryPage(db *gorm.DB, query *query.BlogMessageQuery) []*model.BlogMessageModel {
	modelList := make([]*model.BlogMessageModel, 0)
	cond := db.Where("blog_id = ?", query.BlogId)
	if query.Page != nil && query.Rows != nil {
		cond = cond.Offset(*query.Rows * (*query.Page - 1)).Limit(query.Rows)
	}
	cond.Find(&modelList)
	return modelList
}

func (*blogMessageDao) QueryTotal(db *gorm.DB, query *query.BlogMessageQuery) (total uint32) {
	db.Model(&model.BlogMessageModel{}).Where("blog_id = ?", query.BlogId).Count(&total)
	return
}

func (*blogMessageDao) Create(db *gorm.DB, record *model.BlogMessageModel) error {
	return db.Create(record).Error
}

func (*blogMessageDao) Delete(db *gorm.DB, ids []uint32) error {
	return db.Where("id in (?)", ids).Delete(&model.BlogMessageModel{}).Error
}
