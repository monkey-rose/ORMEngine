package main

import (
    "database/sql"
    "github.com/monkey-rose/ORMEngine/entities"
)

func main() {
	/*
	usr := entities.NewUser("ricky", "cs")
	engine := NewORMEngine()
	_ = engine.Insert(usr)*/

	pEveryOne := make([]entities.UserInfo, 0)
	engine := NewORMEngine()
	_ = engine.Find(&pEveryOne)
	fmt.Println(pEveryOne)
}