package repo

import (
	"blog/convert"
	"blog/dal/dao"
	"blog/dal/domain"
	"blog/info/query"
	"blog/util"
	"github.com/jinzhu/gorm"
	"time"
)

type blogMessageRepo struct{}

func (*blogMessageRepo) QueryPage(db *gorm.DB, query *query.BlogMessageQuery) ([]*domain.BlogMessageDomain, uint32) {
	page, total := dao.BlogMessageDao.QueryPage(db, query), dao.BlogMessageDao.QueryTotal(db, query)
	messageList := make([]*domain.BlogMessageDomain, 0, len(page))
	for _, d := range page {
		messageList = append(messageList, convert.BlogMessageModelToDomain(d))
	}
	return messageList, total
}

func (*blogMessageRepo) Create(db *gorm.DB, record *domain.BlogMessageDomain) error {
	record.CreateTime = util.Uint64Ptr(uint64(time.Now().Unix()))
	record.ModifyTime = record.CreateTime
	return dao.BlogMessageDao.Create(db, convert.BlogMessageDomainToModel(record))
}

func (*blogMessageRepo) Delete(db *gorm.DB, ids []uint32) error {
	return dao.BlogMessageDao.Delete(db, ids)
}
