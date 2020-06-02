package main

import (
	"fmt"
)

type vertex struct{
	x,y int
}
var m map[string]vertex
var n = map[string]vertex{
	"a":vertex{2,2},
	"b":vertex{3,3},
}
var b=map[int]vertex{
	1:{2,3},
	2:{4,5},
}

func maps(){
	m=make(map[string]vertex)
	m["first"]=vertex{1,1}
	fmt.Println(m["first"].x)
	fmt.Println(n["a"])
	fmt.Println(b)

	delete(n,"a")
	elem, ok := n["a"]
	fmt.Println("The value:", elem, "Present?", ok)
}