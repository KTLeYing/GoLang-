package main

import (
	"bufio"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

func main() {
	f, _ := excelize.OpenFile("D:\\MyInstall\\Python\\files\\十万热歌结果(音频指纹两两匹配).xlsx")
	//获取sheet上的所有单元格
	rows, _ := f.GetRows("十万热歌结果(音频指纹两两匹配)")
	var arr [11]int
	fileName := "D:\\MyInstall\\Python\\files\\degree0.txt"
	//创建文件
	f1, _ := os.Create(fileName)
	//创建写文件流
	w := bufio.NewWriter(f1)
	for _, row := range rows {
		//fmt.Println(row)
		//fmt.Println(row[2])
		matchingDegree := row[2]
		matchingDegree1, _ := strconv.ParseFloat(matchingDegree, 64)
		for i := 0; i < 10; i++ {
			min := float64(i) * 0.1
			max := float64(i+1) * 0.1
			if matchingDegree1 >= min && matchingDegree1 < max {
				arr[i]++
			}
			if matchingDegree1 == 1 {
				arr[10]++
			}
		}
		if matchingDegree1 == 0 {
			w.WriteString(row[0] + "\t" + row[1] + "\n")
		}
		fmt.Println()
	}
	//关闭写流和文件流
	w.Flush()
	f1.Close()

	for _, res := range arr {
		fmt.Println(res)
	}

}
