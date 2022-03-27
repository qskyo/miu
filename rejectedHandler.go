package miu

import "fmt"

// RejectedHandler 拒绝策略接口
type RejectedHandler interface {
	RejectedExecution(task Task)
}

type DiscardPolicy struct {
}

func (r *DiscardPolicy) RejectedExecution(task Task) {
	fmt.Printf("the task[id=%s] has been discarded\n", task.GetTaskId())
}
