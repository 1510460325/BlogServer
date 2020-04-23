package model

type BlogModel struct {
	BaseModel
	Title   *string `gorm:"column:title"`
	Content *string `gorm:"column:content"`
	ViewNum *uint32 `gorm:"column:view_num"`
}

func (BlogModel) TableName() string {
	return "Blog"
}
