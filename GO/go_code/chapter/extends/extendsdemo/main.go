package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (stu *Student) Getsum(n1 int, n2 int) int {
	return n1 + n2
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

type Graduate struct {
	Student
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

func main() {
	// var pupil = &Pupil{
	// 	Name: "tom",
	// 	Age: 10,
	// }
	// pupil.testing()
	// pupil.SetScore(90)
	// pupil.ShowInfo()
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	fmt.Println("res=", pupil.Student.Getsum(1, 2))

	graduate := &Graduate{}
	graduate.Student.Name = "merry"
	graduate.Student.Age = 20
	graduate.testing()
	graduate.Student.SetScore(120)
	graduate.Student.ShowInfo()
	fmt.Println("res=", graduate.Student.Getsum(213, 213))
}
