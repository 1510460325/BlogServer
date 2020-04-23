package model

type BlogAboutModel struct {
	ID       *uint32 `gorm:"column:id"`
	Content  *string `gorm:"column:content"`
	ViewNum  *uint32 `gorm:"column:view_num"`
	AdminPwd *string `gorm:"column:admin_pwd"`
}

func (BlogAboutModel) TableName() string {
	return "Blog_about"
}
