package model

type BlogTagsAttachModel struct {
	BaseModel
	BlogId *uint32 `gorm:"column:blog_id"`
	TagsId *uint32 `gorm:"column:tags_id"`
}

func (BlogTagsAttachModel) TableName() string {
	return "Blog_tags_attach"
}
