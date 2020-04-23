package dao

import (
	"blog/dal/model"
	"blog/info/query"

	"github.com/jinzhu/gorm"
)

type blogDao struct{}

func (*blogDao) QueryPage(db *gorm.DB, query *query.BlogQuery) []*model.BlogModel {
	modelList := make([]*model.BlogModel, 0, 10)
	cond := db.Where("1 = 1")
	if query.Title != nil {
		cond = cond.Where("title like ?", "%"+*query.Title+"%")
	}
	if query.ID != nil {
		cond = cond.Where("id = ?", query.ID)
	}
	if query.Page != nil && query.Rows != nil {
		cond = cond.Offset(*query.Rows * (*query.Page - 1)).Limit(query.Rows)
	}
	cond.Find(&modelList)
	return modelList
}

func (*blogDao) QueryByIds(db *gorm.DB, ids []uint32) []*model.BlogModel {
	modelList := make([]*model.BlogModel, 0, 10)
	db.Where("id in (?)", ids).Find(&modelList)
	return modelList
}

func (*blogDao) QueryTotal(db *gorm.DB, query *query.BlogQuery) (total uint32) {
	cond := db.Model(&model.BlogModel{}).Where("1 = 1")
	if query.Title != nil {
		cond = cond.Where("title like ?", "%"+*query.Title+"%")
	}
	if query.ID != nil {
		cond = cond.Where("id = ?", query.ID)
	}
	cond.Count(&total)
	return
}

func (*blogDao) Update(db *gorm.DB, record *model.BlogModel) error {
	return db.Model(&model.BlogModel{}).Update(record).Error
}

func (*blogDao) Create(db *gorm.DB, record *model.BlogModel) error {
	return db.Create(record).Error
}

func (*blogDao) Delete(db *gorm.DB, ids []uint32) error {
	return db.Where("id in (?)", ids).Delete(&model.BlogModel{}).Error
}

func (*blogDao) AddViewNum(db *gorm.DB, id uint32) error {
	return db.Exec("update `Blog` set `view_num` = `view_num` + 1 where `id` = ?", id).Error
}
