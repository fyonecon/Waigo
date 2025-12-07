package main

import "fmt"

type AppGin struct {
}

func TestGin() *AppGin {
	fmt.Println("TEST-GIN========================")
	return &AppGin{}
}

// RunGin Http服务
func (ag *AppGin) RunGin() {
	fmt.Println("RUN_GIN")
}
