package handle

import (
	"fmt"
)

func Main(mapArgs map[string]*string) error {
	for k, v := range mapArgs {
		fmt.Printf("%s=%s\n", k, *v)
	}
	opt := NewOptions(mapArgs)
	dbinfo := NewDBInfo()
	dbinfo.Opt = opt
	data, err := dbinfo.Main()
	if err != nil {
		return err
	}
	mdfile := MDFile{}
	mdfile.StructToLines(data,dbinfo.tables)
	return nil
}
