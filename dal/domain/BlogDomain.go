package domain

type BlogDomain struct {
	BaseDomain
	Title      *string           `json:"title"`
	Content    *string           `json:"content"`
	ViewNum    *uint32           `json:"viewNum"`
	CommentNum *uint32           `json:"commentNum"`
	Tags       []*BlogTagsDomain `json:"tags"`
}
