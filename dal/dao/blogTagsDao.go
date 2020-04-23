package dao

import (
	"blog/dal/model"
	"blog/info/query"
	"fmt"

	"github.com/jinzhu/gorm"
)

type blogTagsDao struct{}

func (*blogTagsDao) Query(db *gorm.DB) []*model.BlogTagsModel {
	modelList := make([]*model.BlogTagsModel, 1)
	db.Find(&modelList)
	return modelList
}

func (*blogTagsDao) QueryByIds(db *gorm.DB, ids []uint32) []*model.BlogTagsModel {
	modelList := make([]*model.BlogTagsModel, 1)
	db.Where("id in (?)", ids).Find(&modelList)
	return modelList
}

func (*blogTagsDao) Update(db *gorm.DB, record *model.BlogTagsModel) error {
	return db.Model(&model.BlogTagsModel{}).Update(record).Error
}

func (*blogTagsDao) Create(db *gorm.DB, record *model.BlogTagsModel) error {
	return db.Create(record).Error
}

func (*blogTagsDao) Delete(db *gorm.DB, ids []uint32) error {
	return db.Where("id in (?)", ids).Delete(&model.BlogTagsModel{}).Error
}

/**   BlogTagsAttach    **/
func (*blogTagsDao) DeleteTagsAttachByBlogIds(db *gorm.DB, blogIds []uint32) error {
	return db.Where("blog_id in (?) ", blogIds).Delete(&model.BlogTagsAttachModel{}).Error
}

func (*blogTagsDao) DeleteTagsAttachByTagsIds(db *gorm.DB, tagsIds []uint32) error {
	return db.Where("tags_id in (?) ", tagsIds).Delete(&model.BlogTagsAttachModel{}).Error
}

func (*blogTagsDao) SaveTagsAttach(db *gorm.DB, records []*model.BlogTagsAttachModel) error {
	sql := "INSERT INTO `Blog_tags_attach` (`blog_id`,`tags_id`,`create_time`, `modify_time`) VALUES "
	// 循环data数组,组合sql语句
	for key, value := range records {
		sql += fmt.Sprintf("(%v, %v, %v, %v)", *value.BlogId, *value.TagsId, *value.CreateTime, *value.ModifyTime)
		if key != len(records)-1 {
			sql += ","
		}
	}
	return db.Exec(sql).Error
}

func (*blogTagsDao) QueryByBlogIds(db *gorm.DB, blogIds []uint32) []*model.BlogTagsAttachModel {
	attachList := make([]*model.BlogTagsAttachModel, 0)
	_ = db.Where("blog_id in (?) ", blogIds).Find(&attachList).Error
	return attachList
}

func (*blogTagsDao) QueryPageByTagsId(db *gorm.DB, query *query.BlogQuery) []*model.BlogTagsAttachModel {
	attachList := make([]*model.BlogTagsAttachModel, 0)
	cond := db.Where("tags_id = ?", query.TagsId)
	if query.Page != nil && query.Rows != nil {
		cond = cond.Offset(*query.Rows * (*query.Page - 1)).Limit(query.Rows)
	}
	cond.Find(&attachList)
	return attachList
}

func (*blogTagsDao) QueryTotalByTagsId(db *gorm.DB, query *query.BlogQuery) (total uint32) {
	db.Model(&model.BlogTagsAttachModel{}).Where("tags_id = ?", query.TagsId).Count(&total)
	return
}
