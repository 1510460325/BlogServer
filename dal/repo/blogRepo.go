package repo

import (
	"github.com/jinzhu/gorm"
	"time"

	"blog/convert"
	"blog/dal/dao"
	"blog/dal/domain"
	"blog/dal/model"
	"blog/info/query"
	"blog/util"
)

type blogRepo struct{}

func (b *blogRepo) QueryPage(db *gorm.DB, query *query.BlogQuery) ([]*domain.BlogDomain, uint32) {
	dataPage := dao.BlogDao.QueryPage(db, query)
	result := make([]*domain.BlogDomain, 0, len(dataPage))
	for _, d := range dataPage {
		result = append(result, convert.BlogModelToDomain(d))
	}
	b.fillTags(db, result)
	return result, dao.BlogDao.QueryTotal(db, query)
}

func (b *blogRepo) QueryByIds(db *gorm.DB, ids []uint32) []*domain.BlogDomain {
	dataPage := dao.BlogDao.QueryByIds(db, ids)
	result := make([]*domain.BlogDomain, 0, len(dataPage))
	for _, d := range dataPage {
		result = append(result, convert.BlogModelToDomain(d))
	}
	b.fillTags(db, result)
	return result
}

func (b *blogRepo) Update(db *gorm.DB, record *domain.BlogDomain) error {
	if err := dao.BlogDao.Update(db, convert.BlogDomainToModel(record)); err != nil {
		return err
	}
	if err := b.saveTagsAttach(db, record); err != nil {
		return err
	}
	return nil
}

func (b *blogRepo) Create(db *gorm.DB, record *domain.BlogDomain) error {
	recordModel := convert.BlogDomainToModel(record)
	if err := dao.BlogDao.Create(db, recordModel); err != nil {
		return err
	}
	record.ID = recordModel.ID
	if err := b.saveTagsAttach(db, record); err != nil {
		return err
	}
	return nil
}

func (*blogRepo) saveTagsAttach(db *gorm.DB, record *domain.BlogDomain) error {
	if err := dao.BlogTagsDao.DeleteTagsAttachByBlogIds(db, []uint32{*record.ID}); err != nil {
		return err
	}
	if len(record.Tags) > 0 {
		tagsIdSet := make(map[*uint32]bool)
		models := make([]*model.BlogTagsAttachModel, 0, len(record.Tags))
		for _, single := range record.Tags {
			if tagsIdSet[single.ID] == false {
				models = append(models, &model.BlogTagsAttachModel{
					BaseModel: model.BaseModel{
						CreateTime: util.Uint64Ptr(uint64(time.Now().Unix())),
						ModifyTime: util.Uint64Ptr(uint64(time.Now().Unix())),
					},
					BlogId: record.ID,
					TagsId: single.ID,
				})
				tagsIdSet[single.ID] = true
			}
		}
		if err := dao.BlogTagsDao.SaveTagsAttach(db, models); err != nil {
			return err
		}
	}
	return nil
}

func (*blogRepo) Delete(db *gorm.DB, ids []uint32) error {
	_ = dao.BlogTagsDao.DeleteTagsAttachByBlogIds(db, ids)
	return dao.BlogDao.Delete(db, ids)
}

func (*blogRepo) AddViewNum(db *gorm.DB, id uint32) error {
	return dao.BlogDao.AddViewNum(db, id)
}

func (*blogRepo) fillTags(db *gorm.DB, data []*domain.BlogDomain) {
	blogIds := make([]uint32, 0, len(data))
	for _, d := range data {
		blogIds = append(blogIds, *d.ID)
	}
	// attach
	attachModels := dao.BlogTagsDao.QueryByBlogIds(db, blogIds)
	tagsIds := make([]uint32, 0, len(attachModels))
	// blog attach map
	attachMap := make(map[uint32][]uint32)
	for _, d := range attachModels {
		tagsIds = append(tagsIds, *d.TagsId)
		attachMap[*d.BlogId] = append(attachMap[*d.BlogId], *d.TagsId)
	}
	// tags map
	tags := dao.BlogTagsDao.QueryByIds(db, tagsIds)
	tagsMap := make(map[uint32]*domain.BlogTagsDomain)
	for _, t := range tags {
		tagsMap[*t.ID] = convert.BlogTagsModelToDomain(t)
	}
	for _, d := range data {
		if tagsAttach, ok := attachMap[*d.ID]; ok {
			for _, t := range tagsAttach {
				d.Tags = append(d.Tags, tagsMap[t])
			}
		}
	}
}
