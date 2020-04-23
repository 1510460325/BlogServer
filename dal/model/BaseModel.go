package model

type BaseModel struct {
	ID         *uint32 `gorm:"column:id"`
	CreateTime *uint64 `gorm:"column:create_time"`
	ModifyTime *uint64 `gorm:"column:modify_time"`
}
