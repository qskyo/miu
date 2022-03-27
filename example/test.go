package main

import (
	"fmt"
	"github.com/qskyo/miu"
	"sync"
	"time"
)

var wg sync.WaitGroup

type TestTask struct {
	Id int
}

func (t *TestTask) Run() {
	fmt.Printf("%d 执行任务\n", t.Id)
	time.Sleep(1 * time.Second)
}

func (t *TestTask) GetTaskId() int {
	return t.Id
}

func main() {
	workPool := miu.NewDefaultWorkPool(2, 10)
	workPool.Start()
	go exec(workPool)
	time.Sleep(3 * time.Second)
	workPool.Shutdown()
	time.Sleep(5 * time.Second)
}

func exec(workPool miu.WorkPool) {
	for i := 0; i < 30; i++ {
		task := &TestTask{
			Id: i,
		}
		workPool.Execute(task)
	}
}
