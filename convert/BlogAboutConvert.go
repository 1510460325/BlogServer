package convert

import (
	"blog/dal/domain"
	"blog/dal/model"
)

func BlogAboutModelToDomain(blogAboutModel *model.BlogAboutModel) *domain.BlogAboutDomain {
	if blogAboutModel == nil {
		return nil
	}
	return &domain.BlogAboutDomain{
		ID:      blogAboutModel.ID,
		Content: blogAboutModel.Content,
		ViewNum: blogAboutModel.ViewNum,
	}
}

func BlogAboutDomainToModel(blogAboutDomain *domain.BlogAboutDomain) *model.BlogAboutModel {
	if blogAboutDomain == nil {
		return nil
	}
	return &model.BlogAboutModel{
		ID:      blogAboutDomain.ID,
		Content: blogAboutDomain.Content,
		ViewNum: blogAboutDomain.ViewNum,
	}
}
