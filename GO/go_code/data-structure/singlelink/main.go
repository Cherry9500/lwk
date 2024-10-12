// 一般都会有头结点，本身不存放数据，用来标识链表头
//单链表查找只能是一个方向，从头到尾，不能反向，也不能自我删除
package main

import (
	"fmt"
)

type HeroNode struct {
	no       int //编号
	name     string
	nickname string
	next     *HeroNode //表示指向下一个结点
}

// 给链表插入一个结点
// 第一种：在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//1.先找到该链表最后的结点
	//2.创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { //到了最后的结点,所以下一个结点是空
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	//3.将newHHeroNode加入到链表的最后
	temp.next = newHeroNode
}

// 第二种，根据编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//1.先找到适当的结点
	//2.创建一个辅助结点
	temp := head
	flag := true
	//让插入的结点的no和temp的下一个结点的no进行比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode一个插入到temp的后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明id已存在,不插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("编号已存在=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

// 删除一个结点
func DelHeroNode(head *HeroNode, id int) {
	temp := head

	flag := false
	//找到要删除的结点
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			//说明id已存在,不插入
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("要删除的id不存在")
	}
}

// 显示练链表的所有信息
func ListHeroNode(head *HeroNode) {
	//1.创建一个辅助结点
	temp := head

	//先判断是不是空链表
	if temp.next == nil {
		fmt.Println("链表为空...")
		return
	}

	//2.遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		//判断是否到尾部
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	//创建头结点
	head := &HeroNode{}
	//创建一个新的HeroNode
	hero1 := *&HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := *&HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := *&HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	hero4 := *&HeroNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}
	//3.加入
	InsertHeroNode2(head, &hero3)
	InsertHeroNode2(head, &hero1)
	InsertHeroNode2(head, &hero2)
	InsertHeroNode2(head, &hero4)

	//4.显示
	ListHeroNode(head)

	//5.删除
	fmt.Println()
	DelHeroNode(head, 2)
	ListHeroNode(head)
}