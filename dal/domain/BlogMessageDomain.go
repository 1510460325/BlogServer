package domain

type BlogMessageDomain struct {
	BaseDomain
	Content *string `json:"content"`
	Email   *string `json:"email"`
	Name    *string `json:"name"`
	BlogId  *uint32 `json:"blogId"`
}
