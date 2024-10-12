package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "张三",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      80000.0,
		Skill:    "吃东西",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("monster序列化后=%v\n", string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "李四"
	a["age"] = 30
	a["address"] = "北京"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "tom"
	m1["age"] = 7
	m1["address"] = [2]string{"北京", "上海"}
	slice = append(slice, m1)

	data, err := json.Marshal(m1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

func testFloat64() {
	var num1 float64 = 213.21

	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
