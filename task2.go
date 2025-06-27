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



---------------------------------------------面向对象1---------------------------------------------------------------

package main
import (
	"fmt"
	"math"
)

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Area 计算圆形面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建矩形实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形 - 面积: %.2f, 周长: %.2f\n", rect.Area(), rect.Perimeter())

	// 创建圆形实例
	circle := Circle{Radius: 4}
	fmt.Printf("圆形 - 面积: %.2f, 周长: %.2f\n", circle.Area(), circle.Perimeter())
}


---------------------------------------------面向对象2---------------------------------------------------------------

package main

import "fmt"

// Person 基础结构体
type Person struct {
	Name string
	Age  int
}

// Employee 员工结构体，组合Person
type Employee struct {
	Person      // 匿名嵌入Person
	EmployeeID string
}

// PrintInfo 打印员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工ID: %s\n姓名: %s\n年龄: %d\n", 
		e.EmployeeID, e.Name, e.Age)
}

func main() {
	// 创建Employee实例
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  28,
		},
		EmployeeID: "E1001",
	}

	// 调用方法输出信息
	emp.PrintInfo()
}
