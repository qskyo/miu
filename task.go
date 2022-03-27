package miu

// Task 任务接口
type Task interface {
	Run()
	GetTaskId() int
}
