package models

import "github.com/bkcarlos/cliplabel/dao"

func Init() {

	// file
	dao.ClipLabelDB.AutoMigrate(&File{})

	// tag
	dao.ClipLabelDB.AutoMigrate(&Tags{}, &FileTag{})

}
