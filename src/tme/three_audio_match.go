package main

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bitly/go-simplejson"
	"os"

	//"github.com/bitly/go-simplejson"
	"strings"
)

/**
三段(多段)音频匹配日志结果处理
*/
func main() {
	//打开excel文本
	f, err := excelize.OpenFile("D:\\MyInstall\\Python\\files\\三段(多段)音频指纹.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	//创建文件
	fileName := "D:\\MyInstall\\Python\\files\\三段指纹并集伴奏结果1(三段(多段)音频指纹).txt"
	file, _ := os.Create(fileName)
	writer := bufio.NewWriter(file)
	writer.WriteString("id" + "\t" + "ids" + "\n")
	//获取工作表的所有行
	rows := f.GetRows("所有有结果日志")
	//遍历所有行
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		//处理id
		url := row[2]
		result := row[10]
		fmt.Println(url + "\t" + result)
		start := strings.Index(url, "id=")
		end := strings.Index(url, "&format")
		id := url[start+3 : end]
		fmt.Println(id)

		//使用simplejson来解析处理
		//转化为json来处理
		json, _ := simplejson.NewJson([]byte(result))
		//转化为多个子map数组
		resultArr, _ := json.Get("results").Array()
		fmt.Println(resultArr)
		ids := ""
		//遍历子map数组
		for i, arr := range resultArr {
			fmt.Println(i, arr)
			//取出各个子object的字段值
			subMap := json.Get("results").GetIndex(i)
			confidence := subMap.Get("confidence").MustInt()
			id := subMap.Get("id").MustString()
			fmt.Println(confidence, id)
			fmt.Println()
			if i == 0 {
				ids += id
				continue
			}
			ids += ", " + id
		}
		writer.WriteString(id + "\t" + ids + "\n")
	}
	writer.Flush()
	file.Close()

}
