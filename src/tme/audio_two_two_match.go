package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

//定义一个结果体(相当于实体类)，可以用map来间接代替处理
type ExcelData struct {
	Ret            int
	MatchingDegree string
	Message        string
}

type partExcelData struct {
	MatchingDegree string
}

/**
音频两两匹配日志结果处理
*/
func main() {
	f, _ := excelize.OpenFile("D:\\MyInstall\\Python\\files\\十万热歌(音频指纹两两匹配).xlsx")
	//获取sheet的所有数据
	rows, _ := f.GetRows("有结果日志")
	fileName := "D:\\MyInstall\\Python\\files\\十万热歌结果1(音频指纹两两匹配).txt"
	file, _ := os.Create(fileName) //创建文件
	//file1, _ := os.OpenFile(fileName, os.O_WRONLY | os.O_APPEND, 0666)
	writer := bufio.NewWriter(file)
	for _, row := range rows {
		result := row[10]
		str := strings.Split(row[2], ">")
		//json转map
		var data map[string]interface{}
		err := json.Unmarshal([]byte(result), &data)
		if err == nil {
			fmt.Println(data["MatchingDegree"])
			match := data["MatchingDegree"].(string)
			str1 := str[0] + "\t" + str[1] + "\t" + match + "\n"
			//写入时，使用带缓存的 *Writer
			writer.WriteString(str1)
		}
		fmt.Println(result)
	}
	writer.Flush()
	file.Close()
}
