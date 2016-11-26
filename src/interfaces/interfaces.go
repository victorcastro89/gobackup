package main

import "fmt"

type gopher struct {
	Name string
	Age  int
}

func (g gopher) jump() string {
	if g.Age < 50 {
		return g.Name + " Can jump High"
	} else {
		return g.Name + " Still can jump"
	}

}

type horse struct {
	Wheight int
	Name    string
}

func (g horse) jump() string {
	if g.Wheight < 50 {
		return append(g.Name, " Can jump High")
	} else {
		return append(g.Name, " Still can jump")
	}

}

type jumper interface {
	jump() string
}

func main() {

	jumperList := getList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.jump())
	}

}

func getList() []jumper {
	victor := &gopher{Name: "Victor", Age: 26}
	leticia := &gopher{Name: "leticia", Age: 62}
	fusquinha := &horse{Name: "Fusquinha", Wheight: 200}
	list := []jumper{victor, leticia, fusquinha}
	return list
}
