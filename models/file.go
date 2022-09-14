package models

import (
	"context"

	"gorm.io/gorm"
)

// File 文件相关的定义
type File struct {
	Model
	Name     string // 文件名
	Path     string // 文件真实路径
	IsDir    bool   // 是否是目录
	ParentID string // 父 ID
}

// GetFileTable 获取文件表
//  @param ctx
//  @param db
//  @return *gorm.DB
func GetFileTable(ctx context.Context, db *gorm.DB) *gorm.DB {
	return getDB(ctx, &File{}, db)
}
