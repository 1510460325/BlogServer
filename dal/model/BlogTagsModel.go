package model

type BlogTagsModel struct {
	BaseModel
	Name *string `gorm:"column:name"`
}

func (BlogTagsModel) TableName() string {
	return "Blog_tags"
}
