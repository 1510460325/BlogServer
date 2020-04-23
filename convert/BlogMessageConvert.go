package convert

import (
	"blog/dal/domain"
	"blog/dal/model"
)

func BlogMessageModelToDomain(blogMessageModel *model.BlogMessageModel) *domain.BlogMessageDomain {
	if blogMessageModel == nil {
		return nil
	}
	return &domain.BlogMessageDomain{
		BaseDomain: BaseModelToDomain(blogMessageModel.BaseModel),
		Name:       blogMessageModel.Name,
		Content:    blogMessageModel.Content,
		Email:      blogMessageModel.Email,
		BlogId:     blogMessageModel.BlogId,
	}
}

func BlogMessageDomainToModel(blogMessageDomain *domain.BlogMessageDomain) *model.BlogMessageModel {
	if blogMessageDomain == nil {
		return nil
	}
	return &model.BlogMessageModel{
		BaseModel: BaseDomainToModel(blogMessageDomain.BaseDomain),
		Name:      blogMessageDomain.Name,
		Content:   blogMessageDomain.Content,
		Email:     blogMessageDomain.Email,
		BlogId:    blogMessageDomain.BlogId,
	}
}
