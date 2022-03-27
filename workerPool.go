package miu

import (
	"fmt"
	"github.com/google/uuid"
)

type WorkPool interface {
	Execute(task Task)
	Start()
	Shutdown()
	GetTaskQueue() chan Task
	GetWorkers() []*Worker
}

type Worker struct {
	id       string
	exit     chan bool
	workPool WorkPool
}

func (worker *Worker) start() {
	for {
		select {
		case <-worker.exit: //收到结束信号
			fmt.Printf("worker[id=%s] exit\n", worker.id)
			return
		// 如果有消息，处理业务
		case task := <-worker.workPool.GetTaskQueue():
			task.Run()
		}
	}
}

func (worker *Worker) stop() {
	worker.exit <- true
}

type DefaultWorkPool struct {
	// 工作池大小
	workerSize int
	// 任务队列大小
	taskQueueSize int
	// 任务（消息）队列
	taskQueue chan Task
	// 工人集合
	workers []*Worker
	// 拒绝策略
	rejectedHandler RejectedHandler
}

func NewDefaultWorkPool(workerSize, taskQueueSize int) *DefaultWorkPool {
	defaultHandler := &DiscardPolicy{}
	return &DefaultWorkPool{
		workerSize:      workerSize,
		taskQueueSize:   taskQueueSize,
		taskQueue:       make(chan Task, taskQueueSize),
		rejectedHandler: defaultHandler,
	}
}

func (w *DefaultWorkPool) Execute(task Task) {
	if len(w.taskQueue) >= w.taskQueueSize {
		fmt.Println("the taskQueue is is full")
		w.rejectedHandler.RejectedExecution(task)
		return
	}
	w.taskQueue <- task
}

func (w *DefaultWorkPool) Start() {
	workers := make([]*Worker, w.workerSize)
	for i := 0; i < w.workerSize; i++ {
		worker := &Worker{
			id:       uuid.NewString(),
			exit:     make(chan bool),
			workPool: w,
		}
		workers[i] = worker
		go worker.start()
	}
	w.workers = workers
}

func (w *DefaultWorkPool) Shutdown() {
	defer close(w.taskQueue)
	workers := w.workers
	for _, worker := range workers {
		worker.stop()
	}
	fmt.Println("workPool exit")
}

func (w *DefaultWorkPool) GetTaskQueue() chan Task {
	return w.taskQueue
}

func (w *DefaultWorkPool) GetWorkers() []*Worker {
	return w.workers
}
