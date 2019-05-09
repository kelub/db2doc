package handle

import (
	"fmt"
	"strings"
)

type Options struct {
	dsn      string
	dbname   string
	user     string
	passwd   string
	addr     string
	params   string
	exclude  []string
	file_dir string
}

func NewOptions(mapArgs map[string]*string) *Options {
	exclude := strings.Split(*mapArgs["exclude"], ",")
	for i := 0; i < len(exclude); i++ {
		exclude[i] = strings.Replace(exclude[i], " ", "", -1)
	}
	fmt.Println("exclude", exclude)
	return &Options{
		dsn:      *mapArgs["dsn"],
		dbname:   *mapArgs["dbname"],
		user:     *mapArgs["user"],
		passwd:   *mapArgs["passwd"],
		addr:     *mapArgs["addr"],
		params:   *mapArgs["params"],
		exclude:  exclude,
		file_dir: *mapArgs["file_dir"],
	}
}
