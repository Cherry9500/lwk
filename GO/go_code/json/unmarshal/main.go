package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

// 反序列化后的数据格式要和原数据格式一样，主要针对字段
// 反序列化map时不需要make，因为make操作被封装到Unmarshal函数里了
func unmarshalStruct() {
	str := "{\"Name\":\"张三\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":80000,\"Skill\":\"吃东西\"}" //直接写str要用转义字符
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)

	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后monster=%v\n monster.Name = %v \n", monster, monster.Name)
}

func main() {
	unmarshalStruct()
}
