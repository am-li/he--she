package main

import (
	"github.com/am-li/he-she/base"
	"github.com/am-li/he-she/countries"
)

func main() {
	c := make(chan string, 1314*520)
	go countries.China(c)
	successStr := <-c
	base.Match(successStr)

}
