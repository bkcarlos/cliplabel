package server

import "context"

// 文件夹相关的操作

// AddFolderReq 添加
type AddFolderReq struct {
	Name string `json:"name" form:"name" binding:"required"`
}

// AddFolderRsp 添加文件夹返回
type AddFolderRsp struct {
	ID   string `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

// AddFolder 添加文件夹
//  @param ctx
//  @param req
//  @return rsp
//  @return err
func AddFolder(ctx context.Context, req AddFolderReq) (rsp AddFolderRsp, err error) {
	return
}
