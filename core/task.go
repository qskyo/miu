package core

// Task 任务接口
type Task interface {
	Run()
	GetTaskId() string
}

type DefaultTask struct {
	Id string
}

func (t *DefaultTask) Run() {

}

func (t *DefaultTask) GetTaskId() string {
	return t.Id
}
