package main

import (
	"fmt"
	human "wildberies/L1/L1.1/humanStruct"
)

type Action struct {
	newField string
	human.Human
}

func main() {
	newAction := Action{
		newField: "test1",
		Human: human.Human{
			Age:    18,
			Name:   "Alex",
			Adress: "Moscow",
			Email:  "alex_new@gmail.com",
			Family: human.FamilyInfo{
				Mother:  "Alena",
				Father:  "Boris",
				Sister:  "Irina",
				Brother: "Fedot",
			},
		},
	}
	fmt.Println(newAction.Human.GetFamily(), "\n", newAction.newField)
}
