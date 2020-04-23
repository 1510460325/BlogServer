package convert

import (
	"blog/dal/domain"
	"blog/dal/model"
)

func BlogTagsModelToDomain(blogTagsModel *model.BlogTagsModel) *domain.BlogTagsDomain {
	if blogTagsModel == nil {
		return nil
	}
	return &domain.BlogTagsDomain{
		BaseDomain: BaseModelToDomain(blogTagsModel.BaseModel),
		Name:       blogTagsModel.Name,
	}
}

func BlogTagsDomainToModel(blogTagsDomain *domain.BlogTagsDomain) *model.BlogTagsModel {
	if blogTagsDomain == nil {
		return nil
	}
	return &model.BlogTagsModel{
		BaseModel: BaseDomainToModel(blogTagsDomain.BaseDomain),
		Name:      blogTagsDomain.Name,
	}
}
