package query

type BlogQuery struct {
	PageQuery
	ID     *uint32 `json:"id"`
	Title  *string `json:"title"`
	TagsId *uint32 `json:"tagsId"`
}
