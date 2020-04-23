package service

import (
	"blog/dal/domain"
	"blog/dal/repo"
	"blog/info"
	"blog/info/query"
	"blog/util"
)

type blogMessageService struct{}

func (*blogMessageService) QueryPage(query *query.BlogMessageQuery) *info.PagingInfo {
	page, total := repo.BlogMessageRepo.QueryPage(util.GetDB(), query)
	return &info.PagingInfo{
		List:  page,
		Total: total,
	}
}

func (*blogMessageService) Create(record *domain.BlogMessageDomain) error {
	return repo.BlogMessageRepo.Create(util.GetDB(), record)
}

func (*blogMessageService) Delete(ids []uint32) error {
	return repo.BlogMessageRepo.Delete(util.GetDB(), ids)
}
