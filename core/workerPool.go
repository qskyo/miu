package core

import (
	"fmt"
)

// WorkPool 工作池抽象
type WorkPool interface {
	// Execute 执行任务
	Execute(task Task)
	// Start 启动工作池
	Start()
	// Shutdown 关闭工作池
	Shutdown()
	GetTaskQueue() chan Task
}

type DefaultWorkPool struct {
	// 工作池大小
	workerSize int
	// 任务（消息）队列
	taskQueue chan Task
	// 工人集合
	workers []*Worker
	// 拒绝策略
	rejectedHandler RejectedHandler
}

type Worker struct {
	id   string
	exit chan bool
	// 任务（消息）队列
	taskQueue chan Task
}

func (worker *Worker) start() {
	for {
		select {
		// 退出消息
		case <-worker.exit:
			fmt.Printf("worker[id=%s] exit\n", worker.id)
			return
		// 业务消息
		case task := <-worker.taskQueue:
			task.Run()
		}
	}
}

func (worker *Worker) stop() {
	worker.exit <- true
}
