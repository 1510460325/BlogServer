package service

import (
	"blog/dal/domain"
	"blog/dal/repo"
	"blog/util"
)

type blogAboutService struct{}

func (*blogAboutService) Query() *domain.BlogAboutDomain {
	return repo.BlogAbout.Query(util.GetDB())
}

func (*blogAboutService) Update(record *domain.BlogAboutDomain) error {
	return repo.BlogAbout.Update(util.GetDB(), record)
}

func (*blogAboutService) CheckAdmin(pwd string) bool {
	return repo.BlogAbout.CheckAdmin(util.GetDB(), pwd)
}

func (*blogAboutService) AddViewNum() error {
	return repo.BlogAbout.AddViewNum(util.GetDB())
}
