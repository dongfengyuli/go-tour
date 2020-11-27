package main

import (
	"fmt"
	"time"
)
//利用chan + 多协程并发处理
func main(){
	// 任务数据
	data_source_chan := make(chan int, 2000)
	// 结果数据
	data_result_chan := make(chan int, 2000)
	gen_num := 8
	// 所有任务协程是否结束
	gen_chan := make(chan bool,gen_num)
	time1 := time.Now().Unix()

	go generate_source(data_source_chan)
	// 协程池,任务分发
	workpool(data_source_chan,data_result_chan,gen_chan,gen_num)
	// 所有协程结束后关闭结果数据channel
	go func() {
		for i:=0;i<gen_num;i++{
			<-gen_chan
		}
		close(data_result_chan)
		fmt.Println("spend timeis ", time.Now().Unix()-time1)
	}()

	fmt.Println("######")
	for data_result := range data_result_chan{
		fmt.Println(data_result)
	}
}


func generate_source(data_source_chan chan int){
	for i:=1;i<=80000;i++{
		data_source_chan <-i
	}
	fmt.Println("写入协程结束")
	close(data_source_chan)
}

func generate_sushu(data_source_chan chan int,data_result_chan chan int,gen_chan chan bool){
	for num := range data_source_chan{
		falg := true
		for i:=2;i<num;i++{
			if num%i == 0{
				falg = false
				break
			}
		}
		if falg == true{
			data_result_chan <- num
		}
	}
	fmt.Println("该协程结束")
	gen_chan <- true
}

func workpool(data_source_chan chan int,data_result_chan chan int,gen_chan chan bool,gen_num int){
	// 开启8个协程
	for i:=0;i<gen_num;i++{
		go generate_sushu(data_source_chan,data_result_chan,gen_chan)
	}
}
