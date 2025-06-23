package main

import "fmt"

// 指针1
// AddTen 接收整数指针并将值增加10
func AddTen(numPtr *int) {
    if numPtr != nil {
        *numPtr += 10
    }
}

func main() {
    number := 5
    fmt.Println("原始值:", number)
    
    AddTen(&number)
    fmt.Println("修改后:", number)
}

// 指针2
// DoubleSlice 接收切片指针并将每个元素乘以2
func DoubleSlice(slicePtr *[]int) {
    if slicePtr == nil {
        return
    }
    
    slice := *slicePtr
    for i := 0; i < len(slice); i++ {
        slice[i] *= 2
    }
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("原切片:", nums)
    
    DoubleSlice(&nums)
    fmt.Println("处理后:", nums)
}



// -----------------------------------------------------------Goroutine 1-------------------------------------------------------------------------------------------//

import (
	"fmt"
	"sync"
)

func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
	}
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	
	go printOdd(&wg)
	go printEven(&wg)
	
	wg.Wait()
	fmt.Println("打印完成")

// ---------------------------------------------------Goroutine 2----------------------------------------------------------------------------//


package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型
type Task func()

// TaskResult 存储任务执行结果
type TaskResult struct {
	ID       int
	Duration time.Duration
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks []Task
	wg    sync.WaitGroup
}

// NewScheduler 创建调度器实例
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]Task, 0),
	}
}

// AddTask 添加任务
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Run 并发执行所有任务并统计耗时
func (s *Scheduler) Run() []TaskResult {
	results := make([]TaskResult, len(s.tasks))
	resultChan := make(chan TaskResult, len(s.tasks))
	s.wg.Add(len(s.tasks))

	for i, task := range s.tasks {
		go func(id int, t Task) {
			defer s.wg.Done()
			start := time.Now()
			t()
			duration := time.Since(start)
			resultChan <- TaskResult{
				ID:       id,
				Duration: duration,
			}
		}(i, task)
	}

	go func() {
		s.wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		results[result.ID] = result
	}

	return results
}

func main() {
	scheduler := NewScheduler()

	// 添加示例任务
	scheduler.AddTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 1 completed")
	})

	scheduler.AddTask(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task 2 completed")
	})

	scheduler.AddTask(func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Task 3 completed")
	})

	// 执行并获取结果
	results := scheduler.Run()

	// 打印统计结果
	for i, result := range results {
		fmt.Printf("Task %d took %v\n", i+1, result.Duration)
	}
}

