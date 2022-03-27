package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/qskyo/miu"
	"sync"
	"time"
)

var wg sync.WaitGroup

type TestTask struct {
	miu.BaseTask
}

func (t *TestTask) Run() {
	fmt.Printf("%s 执行任务\n", t.GetTaskId())
	time.Sleep(1 * time.Second)
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
		task := new(TestTask)
		task.Id = uuid.NewString()
		workPool.Execute(task)
	}
}
