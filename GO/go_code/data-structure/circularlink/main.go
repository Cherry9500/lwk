// 环形单链表
package main

import "fmt"

type CatNode struct {
	no   int // 编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成一个环状
		fmt.Println(newCatNode, "加入并创建一个环形链表")
		return
	}

	//定义一个临时变量，去找到环形的最后一个结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

//输出环形链表
func ListCircularLink(head *CatNode) {
	fmt.Println("环形链表的情况如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空...")
		return
	}
	for {
		fmt.Printf("猫的信息为[id=%d name=%s] ->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode { //和main中的head不是同一个
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("这个环形链表是空的，无法删除")
		return head
	}

	//如果只有一个结点
	if temp.next == head {
		if temp.no == id {
			temp.next = nil
		}
		return head
	}

	//将helper定位到链表尾部
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//两个以上结点
	flag := true
	for {
		if temp.next == head { //已经比较到最后一个，但是还没比较最后一个
			break
		}
		if temp.no == id {
			if temp == head { //删除的是头结点
				head = head.next
			}
			//找到了要删除的结点
			helper.next = temp.next
			fmt.Printf("删除的猫猫id=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     //在于比较，找结点
		helper = helper.next //在于删除
	}
	//再判断最后一次
	if flag { //如果flag为真则表示在上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起,没有no=%d\n", id)
		}
	}
	return head
}

func main() {
	//初始化一个环形链表头结点
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom1",
	}

	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}

	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}

	cat4 := &CatNode{
		no:   4,
		name: "tom4",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	InsertCatNode(head, cat4)

	ListCircularLink(head)

	head = DelCatNode(head, 30)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	ListCircularLink(head)
}
