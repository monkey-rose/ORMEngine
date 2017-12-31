package main

import (
    "github.com/monkey-rose/ORMEngine/entities"
    "fmt"
)

func main() {
	engine := entities.NewORMEngine()
	
	usr := entities.NewUser("ricky", "cs")
	_ = engine.Insert(usr)

	usr2 := entities.NewUser("mini", "law")
	_ = engine.Insert(usr2)

	usr3 := entities.NewUser("cathy", "biology")
	_ = engine.Insert(usr3)

	pEveryOne := make([]entities.UserInfo, 0)
	_ = engine.Find(&pEveryOne)
	fmt.Println(pEveryOne)
}