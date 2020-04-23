package repo

import (
	"blog/convert"
	"blog/dal/dao"
	"blog/dal/domain"
	"github.com/jinzhu/gorm"
)

type blogAboutRepo struct{}

func (*blogAboutRepo) Query(db *gorm.DB) *domain.BlogAboutDomain {
	return convert.BlogAboutModelToDomain(dao.BlogAboutDao.Query(db))
}

func (*blogAboutRepo) Update(db *gorm.DB, record *domain.BlogAboutDomain) error {
	return dao.BlogAboutDao.Update(db, convert.BlogAboutDomainToModel(record))
}

func (*blogAboutRepo) AddViewNum(db *gorm.DB) error {
	return dao.BlogAboutDao.AddViewNum(db)
}

func (*blogAboutRepo) CheckAdmin(db *gorm.DB, pwd string) bool {
	return dao.BlogAboutDao.CheckAdmin(db, pwd)
}
