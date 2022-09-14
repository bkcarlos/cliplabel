package models

import (
	"context"
	"time"

	"github.com/bkcarlos/cliplabel/dao"
	"github.com/bkcarlos/cliplabel/logger"

	"gorm.io/gorm"
)

// Model 基础模型
type Model struct {
	ID        string `gorm:"primarykey;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// getDB 获取数据库
//  @param ctx
//  @param model
//  @return *gorm.DB
func getDB(ctx context.Context, model interface{}, db *gorm.DB) *gorm.DB {
	if db != nil {
		return db.WithContext(ctx).Model(model)
	}

	return dao.ClipLabelDB.WithContext(ctx).Model(model)
}

// Begin 开启事物
//  @param ctx
func Begin(ctx context.Context) *gorm.DB {
	logger.Info(ctx, "open transaction")
	return dao.ClipLabelDB.Begin()
}

// End 关闭事物
//  @param ctx
//  @param tx
//  @param err
func End(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		logger.Error(ctx, "tx has err:%v, rollback", err)
		tx.Rollback()
	} else {
		logger.Info(ctx, "tx succ, commit")
		tx.Commit()
	}
}
