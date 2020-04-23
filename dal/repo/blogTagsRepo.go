package repo

import (
	"blog/convert"
	"blog/dal/dao"
	"blog/dal/domain"
	"blog/info/query"
	"github.com/jinzhu/gorm"
)

type blogTagsRepo struct{}

func (*blogTagsRepo) Query(db *gorm.DB) []*domain.BlogTagsDomain {
	modelList := dao.BlogTagsDao.Query(db)
	result := make([]*domain.BlogTagsDomain, 0, len(modelList))
	for _, d := range modelList {
		result = append(result, convert.BlogTagsModelToDomain(d))
	}
	return result
}

func (*blogTagsRepo) Update(db *gorm.DB, record *domain.BlogTagsDomain) error {
	return dao.BlogTagsDao.Update(db, convert.BlogTagsDomainToModel(record))
}

func (*blogTagsRepo) Create(db *gorm.DB, record *domain.BlogTagsDomain) error {
	return dao.BlogTagsDao.Create(db, convert.BlogTagsDomainToModel(record))
}

func (*blogTagsRepo) Delete(db *gorm.DB, ids []uint32) error {
	_ = dao.BlogTagsDao.DeleteTagsAttachByTagsIds(db, ids)
	return dao.BlogTagsDao.Delete(db, ids)
}

func (*blogTagsRepo) QueryIdsPageByTagsId(db *gorm.DB, query *query.BlogQuery) ([]uint32, uint32) {
	dataPage := dao.BlogTagsDao.QueryPageByTagsId(db, query)
	ids := make([]uint32, 0)
	for _, d := range dataPage {
		ids = append(ids, *d.BlogId)
	}
	return ids, dao.BlogTagsDao.QueryTotalByTagsId(db, query)
}
