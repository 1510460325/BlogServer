package query

type BlogMessageQuery struct {
	PageQuery
	BlogId *uint32 `json:"blogId"`
}
