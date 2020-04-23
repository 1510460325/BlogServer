package info

type BaseInfo struct {
	ID         *uint32 `json:"id"`
	CreateTime *uint64 `json:"createTime"`
	ModifyTime *uint64 `json:"modifyTime"`
}

type JsonResult struct {
	Data    interface{} `json:"data"`
	Code    string      `json:"code"`
	Message *string     `json:"message"`
}
