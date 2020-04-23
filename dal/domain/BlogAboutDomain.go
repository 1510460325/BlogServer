package domain

type BlogAboutDomain struct {
	ID      *uint32 `json:"id"`
	Content *string `json:"content"`
	ViewNum *uint32 `json:"viewNum"`
}
