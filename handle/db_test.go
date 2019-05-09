package handle

import (
	"database/sql"
	"fmt"
	"testing"
)

var user = "root"
var passwd = "SDF-XSP-0056"
var addr = "192.168.9.241:3306"
var dbname = "192_168_9_230_player"
var params = ""
var dsn = ""

var exclude = "t_player_new_red_ticket,t_hall_info"
var file_dir = ""
var mapArgs = map[string]*string{
	"user":     &user,
	"passwd":   &passwd,
	"addr":     &addr,
	"dbname":   &dbname,
	"exclude":  &exclude,
	"file_dir": &file_dir,
	"params":   &params,
	"dsn":      &dsn,
}

/*
COLUMN_NAME 字段名
COLUMN_TYPE 字段类型。比如float(9,3)，varchar(50)。
COLUMN_KEY 索引类型 PRI，代表主键，UNI，代表唯一键，MUL，可重复。
COLUMN_DEFAULT 字段默认值。
IS_NULLABLE 字段是否可以是NULL。
COLUMN_COMMENT 字段注释
EXTRA 其他信息。

*/

func showdata(t *testing.T, rows *sql.Rows, me *MyEngine) {
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	var results []map[string]string
	//i := 0
	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			if err != nil {
				fmt.Println(err)
				t.FailNow()
			}
		}
		row := make(map[string]string)
		for k, v := range values {
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}
}

func Test_db(t *testing.T) {
	opt := NewOptions(mapArgs)
	me, err := NewMyEngine(opt)
	defer me.Close()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	//rows, err := me.db.Query("select * from t_player_id")
	//columns, err := rows.Columns()
	//if err != nil {
	//	fmt.Println(err)
	//	t.FailNow()
	//}
	//fmt.Println("columns:", columns)
	//fmt.Println("rows:", rows)

	rows, err := me.DB.Query("select "+
		"COLUMN_NAME, COLUMN_TYPE,COLUMN_KEY ,COLUMN_DEFAULT , IS_NULLABLE, COLUMN_COMMENT "+
		"from INFORMATION_SCHEMA.COLUMNS "+
		"where TABLE_SCHEMA = ? and TABLE_NAME = ?",
		*mapArgs["dbname"], "t_player_id")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	showdata(t, rows, me)
	//for i,_ := range results{
	//	fmt.Println(results[i])
	//}
}

func Test_db_table(t *testing.T) {
	opt := NewOptions(mapArgs)
	me, err := NewMyEngine(opt)
	defer me.Close()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	rows, err := me.DB.Query("select "+
		"TABLE_NAME "+
		"from INFORMATION_SCHEMA.TABLES "+
		"where TABLE_SCHEMA = ?",
		opt.dbname)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	showdata(t, rows, me)
}

func Test_GetTables(t *testing.T) {
	opt := NewOptions(mapArgs)
	fmt.Println(opt)
	//me, err := NewMyEngine(opt)
	//if err != nil {
	//	fmt.Println(err)
	//	t.FailNow()
	//}
	info := NewDBInfo()
	//info.MyEngine = me
	info.Opt = opt
	info.GetTables()
}

func Test_GetDataFromTable(t *testing.T) {
	opt := NewOptions(mapArgs)
	//me, err := NewMyEngine(opt)
	//if err != nil {
	//	fmt.Println(err)
	//	t.FailNow()
	//}
	info := NewDBInfo()
	//info.MyEngine = me
	info.Opt = opt
	info.GetTables()
	info.GetDataFromTable(info.tables[0])
}
