package handle

import (
	"bufio"
	"fmt"
	"github.com/Sirupsen/logrus"
	"gitlab.com/golang-commonmark/markdown"
	"os"
	"strings"
)

type MDFile struct {
}

func (mf *MDFile) Show() {
	md := markdown.New(markdown.Tables(true))
	fmt.Println(md.RenderToString([]byte("Header\n===\nText")))
}

func (mf *MDFile) StructToLines(datas map[string][]map[string]string, dbname string, tables []string, columns []string) (string, error) {
	var md = []string{}
	//tables = tables[:1]

	md_h1 := fmt.Sprintf("# %s", dbname)
	md = append(md, md_h1)

	var md_cloumn, md_cloumn_line string
	md_cloumn = strings.Join(columns, " | ")
	for i := 0; i < len(columns); i++ {
		md_cloumn_line = md_cloumn_line + "---|"
	}
	for _, v := range tables {
		table := datas[v]
		md_h2 := fmt.Sprintf("\n\n## %s \n", v)
		md = append(md, md_h2)

		md = append(md, md_cloumn)
		md = append(md, md_cloumn_line)

		for i := 0; i < len(table); i++ {
			var line string
			for j := 0; j < len(columns); j++ {
				line = line + table[i][columns[j]] + " | "
			}
			md = append(md, line)
		}
	}
	mds := strings.Join(md, "\n") + "\n\n"
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

func (mf *MDFile) Main(datas map[string][]map[string]string, dbname string, tables []string, columns []string) error {
	Lines, err := mf.StructToLines(datas, dbname, tables, columns)
	if err != nil {
		fmt.Println("得到结构数据错误", err)
		return err
	}
	if err := mf.File(Lines); err != nil {
		fmt.Println("写入数据出错", err)
		return err
	}
	return nil
}
