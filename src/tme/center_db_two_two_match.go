package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"strings"
)

//定义一个结果体(相当于实体类)，可以用map来间接代替处理
type ExcelData1 struct {
	Ret            int
	MatchingDegree string
	Message        string
}

type partExcelData1 struct {
	MatchingDegree string
}

/**
音频两两匹配日志结果处理
*/
func main() {
	f1, _ := excelize.OpenFile("D:\\TencentDevelop\\推送文件\\中央曲库\\task_天琴-两两匹配结果.xlsx")
	//获取sheet的所有数据(结果)
	rows, _ := f1.GetRows("识别日志")
	f, _ := excelize.OpenFile("D:\\TencentDevelop\\推送文件\\中央曲库\\task_天琴(原).xlsx")
	//获取sheet的所有数据(原)
	rows1, _ := f.GetRows("result_1")
	fileName := "D:\\TencentDevelop\\推送文件\\中央曲库\\task_天琴-两两匹配结果.txt"
	file, _ := os.Create(fileName) //创建文件
	//file1, _ := os.OpenFile(fileName, os.O_WRONLY | os.O_APPEND, 0666)
	writer := bufio.NewWriter(file)
	strHeader := "流水id" + "\t" + "相似度" + "\n"
	writer.WriteString(strHeader)
	for i, row := range rows { //先遍历结果excel
		if i == 0 {
			continue
		}
		//再遍历原excel
		result := row[10]
		str := row[2]
		//去除字符串首尾空格
		str = strings.TrimSpace(str)
		fmt.Println("识曲后拼接的url：" + str)
		flowId := "-1"
		for j, row1 := range rows1 {
			if j == 0 {
				continue
			}
			str1 := row1[1] + ">" + row1[2]
			//去除字符串首尾空格
			str1 = strings.TrimSpace(str1)
			fmt.Println("原资源拼接的url：" + str1)
			if str == str1 {
				fmt.Println("匹配到...")
				flowId = row1[0]
				//匹配到则删除这个元素，减少遍历的次数
				rows1 = append(rows1[:j], rows1[j+1:]...)
				break
			}
		}
		fmt.Println("流水id：" + flowId)
		//json转map(获取识曲结果json格式数据的MatchingDegree字段)
		var data map[string]interface{}
		err := json.Unmarshal([]byte(result), &data)
		if err == nil {
			if data["MatchingDegree"] != nil {
				fmt.Println("识曲匹配的结果：", data["MatchingDegree"])
				match := data["MatchingDegree"].(string)
				str1 := flowId + "\t" + match + "\n"
				//写入时，使用带缓存的 *Writer
				writer.WriteString(str1)
				fmt.Println("此次写入的结果：" + str1)
			} else {
				fmt.Println("识曲匹配的结果：", data["ret"])
				match := data["ret"].(float64)
				str1 := flowId + "\t" + strconv.FormatFloat(match, 'f', 0, 64) + "\n"
				//写入时，使用带缓存的 *Writer
				writer.WriteString(str1)
				fmt.Println("此次写入的结果：" + str1)
			}
		}
		fmt.Println("----------------------------------------")
	}
	writer.Flush()
	file.Close()
}
