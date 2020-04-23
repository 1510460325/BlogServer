package domain

type BaseDomain struct {
	ID         *uint32 `json:"id"`
	CreateTime *uint64 `json:"createTime"`
	ModifyTime *uint64 `json:"modifyTime"`
}
