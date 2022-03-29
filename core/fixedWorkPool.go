package core

import (
	"fmt"
	"github.com/google/uuid"
)

type FixedWorkPool struct {
	DefaultWorkPool
}

func NewFixedWorkPool(workerSize, taskQueueSize int) *FixedWorkPool {
	w := &FixedWorkPool{}
	w.workerSize = workerSize
	w.taskQueue = make(chan Task, taskQueueSize)
	w.rejectedHandler = &DiscardPolicy{}

	return w
}

func NewFixedWorkPoolWithRejectedHandler(workerSize, taskQueueSize int, handler RejectedHandler) *FixedWorkPool {
	w := &FixedWorkPool{}
	w.workerSize = workerSize
	w.taskQueue = make(chan Task, taskQueueSize)
	w.rejectedHandler = handler

	return w
}

func (w *FixedWorkPool) Execute(task Task) {
	if len(w.taskQueue) >= cap(w.taskQueue) {
		fmt.Println("the taskQueue is is full")
		w.rejectedHandler.RejectedExecution(task)
		return
	}
	w.taskQueue <- task
}

func (w *FixedWorkPool) Start() {
	workers := make([]*Worker, w.workerSize)
	for i := 0; i < w.workerSize; i++ {
		worker := &Worker{
			id:        uuid.NewString(),
			exit:      make(chan bool),
			taskQueue: w.taskQueue,
		}
		workers[i] = worker
		go worker.start()
	}
	w.workers = workers
}

func (w *FixedWorkPool) Shutdown() {
	defer close(w.taskQueue)
	workers := w.workers
	for _, worker := range workers {
		worker.stop()
	}
	fmt.Println("workPool exit")
}

func (w *FixedWorkPool) GetTaskQueue() chan Task {
	return w.taskQueue
}
