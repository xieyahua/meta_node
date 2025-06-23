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



// Goroutine 1

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

// Goroutine 2
