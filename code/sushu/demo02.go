package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	data_source_chan := make(chan int,500)
	data_result_chan := make(chan int,2000)
	gen_num := 8
	time1 := time.Now().Unix()
	var wg sync.WaitGroup
	go generate_source2(data_source_chan)

	// 开启8个协程
	for i:=0;i<gen_num;i++{
		wg.Add(1)
		go generate_sushu2(data_source_chan,data_result_chan,&wg)
	}
	//workpool2(data_source_chan, data_result_chan,&wg, gen_num)

	wg.Wait()
	close(data_result_chan)
	fmt.Println("spend timeis ", time.Now().Unix()-time1)

	for data_result := range data_result_chan{
		fmt.Println(data_result)
	}

}

func generate_source2(data_source_chan chan int)  {
	for i := 1; i <= 80000; i++ {
		data_source_chan <- i
	}
	fmt.Println("写入协程结束")
	close(data_source_chan)
}

func generate_sushu2(data_source_chan, data_result_chan chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
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
}

func workpool2(data_source_chan chan int, data_result_chan chan int,wg *sync.WaitGroup, gen_num int){
	// 开启8个协程
	for i:=0;i<gen_num;i++{
		wg.Add(1)
		go generate_sushu2(data_source_chan, data_result_chan, wg)
	}
}