package convert

import (
	"blog/dal/domain"
	"blog/dal/model"
)

func BlogModelToDomain(blogModel *model.BlogModel) *domain.BlogDomain {
	if blogModel == nil {
		return nil
	}
	return &domain.BlogDomain{
		BaseDomain: BaseModelToDomain(blogModel.BaseModel),
		Title:      blogModel.Title,
		Content:    blogModel.Content,
		ViewNum:    blogModel.ViewNum,
	}
}

func BlogDomainToModel(blogDomain *domain.BlogDomain) *model.BlogModel{
	if blogDomain == nil {
		return nil
	}
	return  &model.BlogModel{
		BaseModel: BaseDomainToModel(blogDomain.BaseDomain),
		Title:      blogDomain.Title,
		Content:    blogDomain.Content,
		ViewNum:    blogDomain.ViewNum,
	}
}