package miu

// Task 任务接口
type Task interface {
	Run()
	GetTaskId() string
}

type BaseTask struct {
	Id string
}

func (t *BaseTask) Run() {

}

func (t *BaseTask) GetTaskId() string {
	return t.Id
}
