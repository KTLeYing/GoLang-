package main

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bitly/go-simplejson"
	"os"
)

/**
片段音频指纹匹配(多结果)日志结果处理
*/
func main() {
	//打开文件
	f, _ := excelize.OpenFile("D:\\MyInstall\\Python\\files\\gsa(片段音频指纹匹配(多结果)).xlsx")

	//创建文件
	fileName := "D:\\MyInstall\\Python\\files\\gsa的结果1(片段音频指纹匹配(多结果)).txt"
	file, _ := os.Create(fileName)
	//创建写文件流
	writer := bufio.NewWriter(file)
	writer.WriteString("id" + "\t" + "ids" + "\n")

	//读取工作表所有的行
	rows := f.GetRows("所有有结果日志")
	//遍历excel文件
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		id := row[1]
		result := row[10]

		//处理json数据得到ids
		json, _ := simplejson.NewJson([]byte(result))
		//转化主json为多个子map数组
		resultArr, _ := json.Get("results").Array()
		fmt.Println(resultArr)
		ids := ""
		//遍历子resultArr, 处理每个子result对象
		for i, arr := range resultArr {
			fmt.Println(i, arr)
			//取出各个子result对象的字段值
			subMap := json.Get("results").GetIndex(i)
			id1 := subMap.Get("id").MustString()
			if i == 0 {
				ids += id1
				continue
			}
			ids += ", " + id1
		}
		writer.WriteString(id + "\t" + ids + "\n")
	}
	writer.Flush()
	file.Close()

}
