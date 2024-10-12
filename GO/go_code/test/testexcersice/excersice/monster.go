package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (lwk *Monster) Store() bool {
	data, err := json.Marshal(lwk)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}

	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
		return false
	}
	return true
}

func (lwk *Monster) ReStore() bool {
	filePath := "d:/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err = ", err)
		return false
	}

	err = json.Unmarshal(data, lwk)
	if err != nil {
		fmt.Println("Unmarshal err = ", err)
		return false
	}
	return true
}
