// 环形链表要留一个标志位，容量=数组大小-1
package main

import (
	"errors"
	"fmt"
	"os"
)

type CircularQueue struct {
	maxSize int //4
	array   [5]int
	head    int //0
	tail    int //0
}

// 入队列
func (lwk *CircularQueue) Push(val int) (err error) {
	if lwk.IsFull() {
		return errors.New("queue full")
	}

	//lwk.tail再队列尾部，但不包含最后的元素
	lwk.array[lwk.tail] = val // 把值给尾部
	lwk.tail = (lwk.tail + 1) % lwk.maxSize
	return
}

// 出队列
func (lwk *CircularQueue) Pop() (val int, err error) {
	if lwk.IsEmpty() {
		return 0, errors.New("queue empty")
	}

	//head指向队首并且包含队首元素
	val = lwk.array[lwk.head]
	lwk.head = (lwk.head + 1) % lwk.maxSize
	return
}

// 显示队列
func (lwk *CircularQueue) ListQueue() {
	fmt.Println("环形队列的情况：")
	//取出当前队列有多少个元素
	size := lwk.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设计一格辅助变量
	tempHead := lwk.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, lwk.array[tempHead])
		tempHead = (tempHead + 1) % lwk.maxSize
	}
	fmt.Println()
}

// 判断环形队列为满
func (lwk *CircularQueue) IsFull() bool {
	return (lwk.tail+1)%lwk.maxSize == lwk.head
}

// 判断环形队列为空
func (lwk *CircularQueue) IsEmpty() bool {
	return lwk.tail == lwk.head
}

// 取出环形队列的元素
func (lwk *CircularQueue) Size() int {
	//关键算法
	return (lwk.tail + lwk.maxSize - lwk.head) % lwk.maxSize
}

func main() {
	queue := &CircularQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列中获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入exit表示退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("进度队列的数据：")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数据=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
