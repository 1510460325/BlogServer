package service

import (
	"blog/dal/domain"
	"blog/dal/repo"
	"blog/info"
	"blog/info/query"
	"blog/util"
	"time"
)

type blogService struct{}

func (*blogService) QueryPage(query *query.BlogQuery) *info.PagingInfo {
	if query.TagsId == nil {
		page, total := repo.BlogRepo.QueryPage(util.GetDB(), query)
		return &info.PagingInfo{
			List:  page,
			Total: total,
		}
	} else {
		ids, total := repo.BlogTagsRepo.QueryIdsPageByTagsId(util.GetDB(), query)
		page := repo.BlogRepo.QueryByIds(util.GetDB(), ids)
		return &info.PagingInfo{
			List:  page,
			Total: total,
		}
	}
}

func (*blogService) Update(record *domain.BlogDomain) error {
	return repo.BlogRepo.Update(util.GetDB(), record)
}

func (*blogService) Create(record *domain.BlogDomain) error {
	record.ViewNum = util.Uint32Ptr(0)
	record.CreateTime = util.Uint64Ptr(uint64(time.Now().Unix()))
	record.ModifyTime = record.CreateTime
	return repo.BlogRepo.Create(util.GetDB(), record)
}

func (*blogService) Delete(ids []uint32) error {
	return repo.BlogRepo.Delete(util.GetDB(), ids)
}

func (*blogService) AddViewNum(id uint32) bool {
	return repo.BlogRepo.AddViewNum(util.GetDB(), id) == nil
}
