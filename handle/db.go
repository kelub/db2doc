package handle

import (
	"database/sql"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/go-sql-driver/mysql"
	"time"
)

type MyEngine struct {
	DBname 	string
	DB     	*sql.DB
}

func getMysqlConfigs(opt *Options) (conf mysql.Config) {
	conf = mysql.Config{
		User:                 opt.user,
		Passwd:               opt.passwd,
		Net:                  "tcp",
		Addr:                 opt.addr,
		DBName:               opt.dbname,
		AllowNativePasswords: true,
		//Params:               *mapArgs["params"],
		Loc: time.Local,
	}
	return
}

func NewMyEngine(opt *Options) (*MyEngine, error) {
	dsn := getMysqlConfigs(opt)
	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		logrus.Errorln("Open db error", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logrus.Errorln("Ping db error", err)
		return nil, err
	}

	return &MyEngine{
		DBname: opt.dbname,
		DB:     db,
	}, nil
}

func (me *MyEngine) Close() {
	me.DB.Close()
}

type DBInfo struct {
	dbname 		string
	tables   []string
	*MyEngine
	Opt	*Options
	datas 	map[string][]map[string]string	// table:rows
}

func NewDBInfo() *DBInfo {
	return &DBInfo{}
}

func (info *DBInfo) GetTables() {
	e,err := NewMyEngine(info.Opt)
	if err != nil{
		fmt.Println("获取 MyEngine 出错",err)
	}
	rows, err := e.DB.Query("select "+
		"TABLE_NAME "+
		"from INFORMATION_SCHEMA.TABLES "+
		"where TABLE_SCHEMA = ?",
		info.Opt.dbname)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer e.Close()
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		rows.Scan(scanArgs...)
		for _, v := range values {
			for _,ex := range info.Opt.exclude{
				if ex != string(v){
					info.tables = append(info.tables, string(v))
				}
			}
		}
		//fmt.Println("tables",info.tables)
	}
	fmt.Println("tables",info.tables)
}

func (info *DBInfo) GetDataFromTable(table string){
	e,err := NewMyEngine(info.Opt)
	if err != nil{
		fmt.Println("获取 MyEngine 出错",err)
	}
	rows, err := e.DB.Query("select "+
		"COLUMN_NAME, COLUMN_TYPE,COLUMN_KEY ,COLUMN_DEFAULT , IS_NULLABLE, COLUMN_COMMENT "+
		"from INFORMATION_SCHEMA.COLUMNS "+
		"where TABLE_SCHEMA = ? and TABLE_NAME = ?",
		info.Opt.dbname, table)
	defer e.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var results  []map[string]string
	for rows.Next(){
		rows.Scan(scanArgs...)
		row := make(map[string]string)
		for k,v := range values{
			key := columns[k]
			row[key] = string(v)
		}
		results = append(results,row)
	}
	info.datas = make(map[string][]map[string]string)
	info.datas[table] = results
	fmt.Println("info.datas:",info.datas)
}