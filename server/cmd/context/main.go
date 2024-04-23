package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct {
}

func main() {
	c := context.WithValue(context.Background(),
		paramKey{}, "abc",
	)
	c, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()
	mainTask(c)
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))
	// c1, cancel := context.WithTimeout(c, 1*time.Second)
	// defer cancel()
	// // task1任务 使用c1 超时时间限制在1秒之内 并且 在c的5秒之内
	// smallTask(c1, "task1", 2*time.Second)
	// // task2任务 使用c1超时时间剩余时间
	// smallTask(c, "task2", 4*time.Second)
	go func() {
		c1, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		// task1任务 使用新的context 超时时间
		smallTask(c1, "task1", 10*time.Second)
	}()
	// task2任务 使用c的超时时间
	smallTask(c, "task2", 4*time.Second)
}

func smallTask(c context.Context, name string, d time.Duration) {
	fmt.Printf("%s started with param %q\n", name, c.Value(paramKey{}))
	select {
	case <-time.After(d):
		fmt.Printf("%s done\n", name)
	case <-c.Done():
		fmt.Printf("%s cancelled\n", name)
	}

}
