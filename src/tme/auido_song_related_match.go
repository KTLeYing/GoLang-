package main

import (
	"bufio"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/xuri/excelize/v2"
	"os"
)

/**
音频歌曲关联配日志结果处理
*/
func main() {
	f, _ := excelize.OpenFile("D:\\MyInstall\\Python\\files\\ORC_Recognition_Result(音频歌曲关联).xlsx")
	//获取sheet的所有数据
	rows, _ := f.GetRows("有结果日志")
	//创建文件
	fileName := "D:\\MyInstall\\Python\\files\\QRC_Recognition_Result结果1(音频歌曲关联).txt"
	file, _ := os.Create(fileName) //创建文件
	//file1, _ := os.OpenFile(fileName, os.O_WRONLY | os.O_APPEND, 0666)
	//创建写文件流
	writer := bufio.NewWriter(file)
	writer.WriteString("url" + "\t" + "ids" + "\n")

	//遍历所有行
	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		//处理url
		url := row[2]
		//处理结果返回体
		resultBody := row[10]
		fmt.Println(url + "\t" + resultBody)
		//处理json数据
		json, _ := simplejson.NewJson([]byte(resultBody))
		//从主json中获取results字段(数组)
		resultArr, _ := json.Get("results").Array()
		//拼接ids
		var ids string = ""
		//遍历results数组
		for idx, _ := range resultArr {
			//取出results数组中各个子result的字段值
			subRes := json.Get("results").GetIndex(idx) //用下标来获取子body
			//从results数组获取body字段（数组）
			bodyArr := subRes.Get("body").MustArray() //转为数组
			//遍历bodyArr数组
			for idx1, _ := range bodyArr {
				//取出bodyArr数组中各个子body的字段值
				subBody := subRes.Get("body").GetIndex(idx1) //用下标来获取子body
				//或
				//subBody := json.Get("results").GetIndex(idx).Get("body").GetIndex(idx1)
				trackId := subBody.Get("track_id").MustString() //转为字符串
				fmt.Println(trackId)
				if idx == 0 {
					ids += trackId
					continue
				}
				ids += ", " + trackId
			}
		}
		writer.WriteString(url + "\t" + ids + "\n")

	}
	writer.Flush()
	file.Close()

}
