package query

type PageQuery struct {
	Page *uint32 `json:"page"`
	Rows *uint32 `json:"rows"`
}
