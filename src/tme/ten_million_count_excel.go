package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

/**
在匹配度帅选结果上来根据匹配度来统计matchDegree各范围的数量
*/
func main() {
	f, _ := excelize.OpenFile("D:\\MyInstall\\Python\\files\\十万热歌结果(音频指纹两两匹配).xlsx")
	//获取sheet上的所有单元格
	rows, _ := f.GetRows("十万热歌结果(音频指纹两两匹配)")
	var arr [11]int

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

		fmt.Println()
	}

	for _, res := range arr {
		fmt.Println(res)
	}

}
