package handle

import (
	"fmt"
)

/*
COLUMN_NAME 字段名
COLUMN_TYPE 字段类型。比如float(9,3)，varchar(50)。
COLUMN_KEY 索引类型 PRI，代表主键，UNI，代表唯一键，MUL，可重复。
COLUMN_DEFAULT 字段默认值。
IS_NULLABLE 字段是否可以是NULL。
COLUMN_COMMENT 字段注释
EXTRA 其他信息。

*/

var Columns = []string{
	"COLUMN_NAME",
	"COLUMN_TYPE",
	"COLUMN_KEY",
	"COLUMN_DEFAULT",
	"IS_NULLABLE",
	"COLUMN_COMMENT",
	"EXTRA",
}

func Main(mapArgs map[string]*string) error {
	for k, v := range mapArgs {
		fmt.Printf("%s=%s\n", k, *v)
	}
	opt := NewOptions(mapArgs)
	dbinfo := NewDBInfo()
	dbinfo.Opt = opt
	dbinfo.Columns = Columns
	data, err := dbinfo.Main()
	if err != nil {
		return err
	}
	mdfile := MDFile{}
	return mdfile.Main(data,dbinfo.dbname,dbinfo.tables,dbinfo.Columns)
}
