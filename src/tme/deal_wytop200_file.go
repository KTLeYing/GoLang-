package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
读取网易top200内容（已补充tmeid）.txt文件  111,2222,2222..... 然后把它们进
行逗号分割，最后输出每个数占一行的形式到一个新文件，并导出到目标目录 [D:\TencentDevelop\lester文件\读取网易top200内容（已补充tmeid）.txt文件]
*/
func main() {
	f, _ := ioutil.ReadFile("D:\\TencentDevelop\\推送文件\\同歌组数据\\Q音ID.txt")
	str := string(f)
	fmt.Println(str)
	strs := strings.Split(str, ",")
	fmt.Println(strs)

	filePath := "D:\\TencentDevelop\\推送文件\\同歌组数据\\Q音ID1.txt"
	file, _ := os.Create(filePath) //创建文件
	//及时关闭file句柄
	defer file.Close()
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for _, s := range strs {
		writer.WriteString(s + "\n")
	}
	writer.Flush()
	file.Close()
}
