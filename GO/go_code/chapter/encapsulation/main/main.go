package main

import(
	"fmt"
	"goproject/go_code/chapter/encapsulation/model"
)

func main()  {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p.Name,p.GetAge(),p.GetSal())
}