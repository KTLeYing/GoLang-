package main

import (
	"bufio"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

/**
导出十万热歌识曲无结果得日志存到txt
*/
func main() {
	f, _ := excelize.OpenFile("D:\\MyInstall\\Python\\files\\十万热歌(音频指纹两两匹配).xlsx")
	//获取sheet的所有数据
	rows, _ := f.GetRows("识别日志")
	fileName := "D:\\MyInstall\\Python\\files\\十万热歌失败结果(音频指纹两两匹配).txt"
	file, _ := os.Create(fileName) //创建文件
	writer := bufio.NewWriter(file)
	count := 0
	for _, row := range rows {
		resCode := row[6]
		fmt.Println(resCode)
		if resCode != "0.0" {
			//字符串转数组化字符串
			var result string
			for _, data := range row { //遍历数组中所有元素追加成string
				result += data + "\t"
			}
			count++
			writer.WriteString(result + "\n")
		}
	}
	fmt.Println(count)
	writer.Flush()
	file.Close()
}
