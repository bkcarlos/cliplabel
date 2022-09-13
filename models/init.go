package models

func Init() {

}

// func CreateFilesTable() {
// 	sqlStmt := `
// 	create table foo (id varchar(20) not null primary key, name varchar(1024) not null, parent_id varchar(20) not null, path varchar(4096) not null, is );
// 	delete from foo;
// 	`

// 	rsp, err := dbExec(sqlStmt)
// 	if err != nil {
// 		logger.Errors("db exec (%v) is err:%v", sqlStmt, err)
// 		panic(fmt.Errorf("db exec(%v) is err:%v", sqlStmt, err))
// 	}

// 	logger.Debugs("db exec rsp:%v", rsp)
// 	logger.Infos("(%v) is success", sqlStmt)
// }

// func dbExec(sqlStmt string) (rsp sql.Result, err error) {
// 	return dao.ClipLabelDB.Exec(sqlStmt)
// }
