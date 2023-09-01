package model

import "gorm.io/gorm"

// BaseModel 模型基类
// 用于定义公共字段
// CreatedAt 创建时间
// UpdatedAt 更新时间
// DeletedAt 删除时间（软删除，为 null 值时表示未被删除）
type BaseModel struct {
	CreatedAt int64          `gorm:"column:created_at;type:bigint(20);autoCreateTime:milli;autoUpdateTime:false;comment:创建时间"`
	UpdatedAt int64          `gorm:"column:updated_at;type:bigint(20);autoUpdateTime:milli;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
