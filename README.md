# miu

Miu is a go work pool implementation, which can help you manage gorontine in the project。

## Installation

miu requires a Go version with [modules](https://github.com/golang/go/wiki/Modules) support. So make sure to initialize a Go module:

```
$ go mod init github.com/my/project
```

And then install miu:

```
$ go get github.com/qskyo/miu
```

## Quickstart

```go
package main

import (
    "fmt"
    "time"
    "github.com/google/uuid"
    "github.com/qskyo/miu/core"
)

type TestTask struct {
    core.DefaultTask
}

func (t *TestTask) Run() {
    fmt.Printf("task[id=%s] is running\n", t.GetTaskId())
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

```


