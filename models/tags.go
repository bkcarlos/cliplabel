package models

import (
	"context"

	"gorm.io/gorm"
)

// Tags 标签
type Tags struct {
	Model
	Key   string // tag key
	Value string // tag value
}

// GetTagsTable 获取 tags 表
//  @param ctx
//  @param db
//  @return *gorm.DB
func GetTagsTable(ctx context.Context, db *gorm.DB) *gorm.DB {
	return getDB(ctx, &Tags{}, db)
}

// FileTag file tag
type FileTag struct {
	Model
	FileID string `gorm:"idx_file_tags"`
	TagsID string `gorm:"idx_file_tags"`
}

// GetFileTagTable
//  @param ctx
//  @param db
//  @return *gorm.DB
func GetFileTagTable(ctx context.Context, db *gorm.DB) *gorm.DB {
	return getDB(ctx, &FileTag{}, db)
}
