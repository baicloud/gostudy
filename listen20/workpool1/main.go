package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	id      int
	randnum int
}

type Result struct {
	task   *Task
	result int
}

func calc(task *Task, result chan *Result) {
	var sum int
	num := task.randnum
	for num != 0 {
		tmp := num % 10
		sum += tmp
		num /= 10
	}
	r := &Result{
		task:   task,
		result: sum,
	}
	result <- r
}

func Worker(taskChan chan *Task, resultChan chan *Result) {
	for task := range taskChan {
		fmt.Println("计算任务开始")
		time.Sleep(3 * time.Second)
		calc(task, resultChan)
	}
}

func startWorkerPool(num int, taskChan chan *Task, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go Worker(taskChan, resultChan)
	}
}

func printResult(resultChan chan *Result) {
	for result := range resultChan {
		fmt.Printf("job id:%v number:%v result:%d\n", result.task.id, result.task.randnum, result.result)
	}
}

func main() {
	var taskChan = make(chan *Task, 100)
	var resultChan = make(chan *Result, 100)
	startWorkerPool(8, taskChan, resultChan)
	go printResult(resultChan)
	var id int
	for {
		id++
		num := rand.Intn(899) + 100
		task := &Task{
			id:      id,
			randnum: num,
		}
		taskChan <- task
	}
}
