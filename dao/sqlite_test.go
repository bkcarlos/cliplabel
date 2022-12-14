package dao

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bkcarlos/cliplabel/configs"
	"github.com/bkcarlos/cliplabel/logger"
)

func TestMain(m *testing.M) {
	os.Setenv("CLIP_LABEL_PATH", "/Users/bkcarlos/code/clip_tools/cliplabel/configs")
	os.Setenv("CLIP_LABEL_SIMPLE", "/Users/bkcarlos/code/clip_tools/cliplabel")

	configs.Init()
	logger.InitLogger(configs.Conf.Log.FilePath, configs.Conf.Log.Level, configs.Conf.Server.Name)
	Init()

	m.Run()
}

func TestInsert(t *testing.T) {
	// rsp := ClipLabelDB.Exec(`select simple_query('pinyin')`)
	// if rsp.Error != nil {
	// 	t.Error(rsp.Error)
	// 	return
	// }
	// fmt.Println(json.Marshal(rsp))

	// rsp := ClipLabelDB.Exec(`CREATE VIRTUAL TABLE t1 USING fts5(x, tokenize = 'simple');`)
	// if rsp.Error != nil {
	// 	t.Error(rsp.Error)
	// 	return
	// }

	// fmt.Printf("%v", rsp)

	// rsp = ClipLabelDB.Exec(`insert into t1(x) values ('周杰伦 Jay Chou:最美的不是下雨天，是曾与你躲过雨的屋檐'),
	// ('I love China! 我爱中国!'),
	// ('@English &special _characters."''bacon-&and''-eggs%');`)
	// if rsp.Error != nil {
	// 	t.Error(rsp.Error)
	// 	return
	// }

	// // fmt.Printf("%v\n", rsp)

	rspmap := make([]map[string]interface{}, 0)
	err := ClipLabelDB.Table("t1").Scan(&rspmap).Error
	if err != nil {
		t.Error(err)
		return
	}
	rspbs, _ := json.Marshal(rspmap)
	fmt.Printf("%v\n", string(rspbs))

	// defer rows.Close()
	// for rows.Next() {
	// 	var query string
	// 	rows.Scan(&query)
	// 	fmt.Printf("%s\n", query)
	// }

	// select '    ', simple_highlight(t1, 0, '[', ']') from t1 where x match simple_query('杰伦');
	rsp := ClipLabelDB.Exec(`select simple_highlight(t1, 0, '[', ']') from t1 where x match sample_query('杰伦');`)
	if rsp.Error != nil {
		t.Error(rsp.Error)
		return
	}

	fmt.Printf("%v\n", rsp)

	// defer rows.Close()
	// for rows.Next() {
	// 	var query string
	// 	rows.Scan(&query)
	// 	fmt.Printf("%s\n", query)
	// }

	// rows, err = ClipLabelDB.Query(`select simple_highlight(t1, 0, '[', ']') from t1 where x match simple_query('I');`)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// defer rows.Close()
	// for rows.Next() {
	// 	var query string
	// 	rows.Scan(&query)
	// 	fmt.Printf("%s\n", query)
	// }
}
