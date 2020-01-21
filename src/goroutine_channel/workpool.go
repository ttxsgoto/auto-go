package main

import (
	//"time"
	"sync"
	"math/rand"
	"fmt"
	"time"
)
// 工作池,通常为一组线程,它们正等待分配给它们的任务,一旦完成分配的任务，他们就会再次为下一个任务提供服务
// https://www.yuque.com/ksco/ogg7um/vs1irw

/*
	• 创建一个goroutine池，该池在等待分配任务的输入缓冲信道上监听
	• 向输入缓冲信道添加任务
	• 任务完成后将结果写入输出缓冲信道
	• 从输出缓冲信道读取和输出结果
*/


type Job struct {
	id       int
	randomno int // 单个数字的和
}

type Result struct {
	job         Job // 任务
	sumofdigits int // 保存结果
}

var jobs = make(chan Job, 10)       // 任务
var results = make(chan Result, 10) // 结果

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	// 创建worker Goroutine
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noofWorkers int) {
	// 创建线程池
	//var wg sync.WaitGroup
	//for i:=0;i<num;i++ {
	//	wg.Add(1)
	//	go process(i, &wg)
	//}
	//wg.Wait()
	var wg sync.WaitGroup
	for i := 0; i < noofWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		//fmt.Println("------>", result)
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	fmt.Println("-xxxx-", digits(123))
	//startTime := time.Now()
	//noofJobs := 100
	//go allocate(noofJobs)
	//
	//done := make(chan bool)
	//go result(done)
	//noofWorkers := 10
	//createWorkerPool(noofWorkers)
	//<-done
	//
	//endTime := time.Now()
	//diff := endTime.Sub(startTime)
	//fmt.Println("Total time taken ", diff)
}















