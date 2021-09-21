package main

import (
	"fmt"
	"math/rand"
)

//任务实体类
type Job struct {
	Id      int
	RandNum int
}

//结果实体类
type Result struct {
	job *Job
	sum int
}

func main() {
	//需2个通道
	//1. job通道
	jobChan := make(chan *Job, 128)
	//2. 结果通道
	resultChan := make(chan *Result, 128)
	//3. 创建工作池
	createPool(64, jobChan, resultChan)
	//4. 开一个打印的协程
	go func(resultChan chan *Result) {
		//遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	//循环创建job，输入到管道
	for {
		id++
		//生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		//把得到的随机数发送到管道中方便后面进行处理
		jobChan <- job
	}
}

//创建工作池
//参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	//根据线程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			//执行运算
			//遍历job管道所有的数据，进行相加
			for job := range jobChan {
				//随机数接过来
				r_num := job.RandNum
				//随机数每位相加
				//定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				//想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
