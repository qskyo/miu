package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/qskyo/miu/core"
	"time"
)

type TestTask struct {
	core.DefaultTask
}

func (t *TestTask) Run() {
	fmt.Printf("%s 执行任务\n", t.GetTaskId())
	time.Sleep(1 * time.Second)
}

func main() {
	workPool := core.NewFixedWorkPool(2, 10)
	workPool.Start()
	go exec(workPool)
	time.Sleep(3 * time.Second)
	workPool.Shutdown()
	time.Sleep(5 * time.Second)
}

func exec(workPool core.WorkPool) {
	for i := 0; i < 30; i++ {
		task := &TestTask{}
		task.Id = uuid.NewString()
		workPool.Execute(task)
	}
}
