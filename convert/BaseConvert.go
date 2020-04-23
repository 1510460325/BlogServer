package convert

import (
	"blog/dal/domain"
	"blog/dal/model"
	"blog/info"
)

func BaseModelToDomain(model model.BaseModel) domain.BaseDomain {
	return domain.BaseDomain{
		ID:         model.ID,
		CreateTime: model.CreateTime,
		ModifyTime: model.ModifyTime,
	}
}

func BaseDomainToModel(baseDomain domain.BaseDomain) model.BaseModel {
	return model.BaseModel{
		ID:         baseDomain.ID,
		CreateTime: baseDomain.CreateTime,
		ModifyTime: baseDomain.ModifyTime,
	}
}

func BaseDomainToInfo(baseDomain domain.BaseDomain) info.BaseInfo {
	return info.BaseInfo{
		ID:         baseDomain.ID,
		CreateTime: baseDomain.CreateTime,
		ModifyTime: baseDomain.ModifyTime,
	}
}
