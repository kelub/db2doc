package handle

import (
	"database/sql"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type MyEngine struct {
	DBname string
	DB     *sql.DB
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
	dbname string
	tables []string
	Columns []string
	//*MyEngine
	Opt   *Options
	datas map[string][]map[string]string // table:rows
}

func NewDBInfo() *DBInfo {
	return &DBInfo{}
}

func (info *DBInfo) Main() (map[string][]map[string]string, error) {
	info.GetTables()
	for i := 0;i< len(info.tables);i++{
		info.GetDataFromTable(info.tables[i])
	}
	return info.datas, nil
}

func (info *DBInfo) GetTables() {
	e, err := NewMyEngine(info.Opt)
	if err != nil {
		fmt.Println("获取 MyEngine 出错", err)
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
			info.tables = append(info.tables, string(v))
		}
	}
	ex := info.Opt.exclude
	for i := 0;i<len(ex);i++{
		for j := 0;j<len(info.tables);j++{
			if ex[i] == info.tables[j]{
				info.tables = append(info.tables[:j],info.tables[j+1:]... )
			}
		}
	}
}

func (info *DBInfo) GetDataFromTable(table string) (map[string][]map[string]string, error) {
	e, err := NewMyEngine(info.Opt)
	if err != nil {
		fmt.Println("获取 MyEngine 出错", err)
	}
	rows, err := e.DB.Query("select "+ strings.Join(info.Columns, ", ") +
		" from INFORMATION_SCHEMA.COLUMNS "+
		" where TABLE_SCHEMA = ? and TABLE_NAME = ?",
		info.Opt.dbname, table)
	defer e.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var results []map[string]string
	for rows.Next() {
		rows.Scan(scanArgs...)
		row := make(map[string]string)
		for k, v := range values {
			key := columns[k]
			row[key] = string(v)
		}
		//row["columns"] = strings.Join(columns_info,",")
		results = append(results, row)
	}
	info.datas[table] = results
	return info.datas, nil
}
