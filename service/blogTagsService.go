package service

import (
	"blog/dal/domain"
	"blog/dal/repo"
	"blog/util"
	"time"
)

type blogTagsService struct{}

func (*blogTagsService) Query() []*domain.BlogTagsDomain {
	return repo.BlogTagsRepo.Query(util.GetDB())
}

func (*blogTagsService) Update(record *domain.BlogTagsDomain) error {
	return repo.BlogTagsRepo.Update(util.GetDB(), record)
}

func (*blogTagsService) Create(record *domain.BlogTagsDomain) error {
	record.CreateTime = util.Uint64Ptr(uint64(time.Now().Unix()))
	record.ModifyTime = record.CreateTime
	return repo.BlogTagsRepo.Create(util.GetDB(), record)
}

func (*blogTagsService) Delete(ids []uint32) error {
	return repo.BlogTagsRepo.Delete(util.GetDB(), ids)
}
