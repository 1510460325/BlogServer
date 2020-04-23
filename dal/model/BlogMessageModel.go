package model

type BlogMessageModel struct {
	BaseModel
	Content *string `gorm:"column:content"`
	Email   *string `gorm:"column:email"`
	Name    *string `gorm:"column:name"`
	BlogId  *uint32 `gorm:"column:blog_id"`
}

func (BlogMessageModel) TableName() string {
	return "Blog_message"
}
