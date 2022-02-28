package main

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
)

/**
导出各个十万热歌matchDegree范围的数据到txt文件（通用，只需修改度的范围大小）
*/
func main() {
	//打开excel
	f, _ := excelize.OpenFile("D:\\\\MyInstall\\\\Python\\\\files\\\\十万热歌结果(音频指纹两两匹配).xlsx")
	//创建存储文件txt
	fileName := "D:\\MyInstall\\Python\\files\\matchDegree_0.txt"
	f1, _ := os.Create(fileName)
	//创建写文件流
	w := bufio.NewWriter(f1)
	w.WriteString("url1" + "\t" + "url2" + "\n")

	count := 0

	//获取工作表所有的行
	rows := f.GetRows("十万热歌结果(音频指纹两两匹配)")
	//遍历行
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		matchDegree := row[2]
		matchDegree1, _ := strconv.ParseFloat(matchDegree, 64)

		//matchDegree：0
		if matchDegree1 == 0 {
			w.WriteString(row[0] + "\t" + row[1] + "\n")
			count++
		}

		//matchDegree：0.2~0.3
		//if matchDegree1 >= 0.2 && matchDegree1 < 0.3{
		//	w.WriteString(row[0] + "\t" + row[1] + "\n")
		//	count++
		//}
	}
	//关闭写文件流和文件(顺序相反)
	w.Flush()
	f1.Close()

	fmt.Println("符合行数：" + strconv.FormatInt(int64(count), 10))
}
