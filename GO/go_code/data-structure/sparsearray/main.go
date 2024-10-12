package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑
	chessMap[2][3] = 2 //蓝
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	var sparseArr []ValNode

	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个vaLNode值结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	fmt.Println("当前的稀疏数组是:::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	//恢复原数组,存盘和恢复(读盘)一定要从文件开始
	var chessMap2 [11][11]int
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	fmt.Println("恢复后的原始数据.......")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
