package handle

import (
	"bufio"
	"fmt"
	"github.com/Sirupsen/logrus"
	"gitlab.com/golang-commonmark/markdown"
	"os"
	"sort"
	"strings"
)

type MDFile struct {
}

func (mf *MDFile) Show() {
	md := markdown.New(markdown.Tables(true))
	fmt.Println(md.RenderToString([]byte("Header\n===\nText")))
}

func (mf *MDFile) StructToLines(datas map[string][]map[string]string, tables []string) (string, error) {
	var md = []string{}
	tables = tables[:1]
	for _, v := range tables {
		table := datas[v]
		md_h1 := fmt.Sprintf("# %s", v)
		md = append(md, md_h1)

		var keys = []string{}
		for key, _ := range table[0] {
			keys = append(keys, key)
		}
		//排序 字段
		sort.Sort(sort.StringSlice(keys))
		var md_cloumn, md_cloumn_line string
		md_cloumn = strings.Join(keys," | ")

		for i := 0; i < len(keys); i++ {
			md_cloumn_line = md_cloumn_line + "---|"
		}
		md = append(md, md_cloumn)
		md = append(md, md_cloumn_line)
		for i := 0; i < len(table); i++ {
			var line string
			for j := 0; j < len(keys); j++ {
				line = line + table[i][keys[j]] + " | "
			}
			md = append(md, line)
		}
	}
	fmt.Println(md)
	mds := strings.Join(md,"\n") + "\n\n"
	fmt.Println(mds)
	return mds, nil
}

func (mf *MDFile) File(datas string) error {
	fd, err := os.OpenFile("a.md", os.O_RDWR|os.O_CREATE, 0644)
	defer fd.Close()
	if err != nil {
		logrus.Errorln("Open db error", err)
		return err
	}
	w := bufio.NewWriter(fd)
	fmt.Fprintln(w, datas)
	w.Flush()
	return nil
}

func (mf *MDFile) Main(){

}