// 队列是一个有序列表，可以用数组和链表实现
// 先入先出原则
// 基本队列：无法复用
package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int //数组是一种数据类型，数组=>模拟队列
	front   int    //指向队列的队首，不包含第一个元素
	rear    int    //指向队列的尾部
}

// 添加数据到队列
func (lwk *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if lwk.rear == lwk.maxSize-1 { //指向并包含最后一个元素
		return errors.New("queue full")
	}

	lwk.rear++ //rear后移
	lwk.array[lwk.rear] = val
	return
}

// 从队列中取出数据
func (lwk *Queue) GetQueue() (val int, err error) {
	//先判断是否为空
	if lwk.front == lwk.rear {
		return -1, errors.New("queue empty")
	}

	lwk.front++
	val = lwk.array[lwk.front]
	return val, err
}

// 显示队列,找到队首，然后遍历到队尾
func (lwk *Queue) ShowQueue() {
	fmt.Println("当前队列是：")
	for i := lwk.front + 1; i <= lwk.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, lwk.array[i])
	}
	fmt.Println()
}

func main() {
	//创建队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数据=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
