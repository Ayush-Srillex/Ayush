package main

// import "golang.org/x/tour/pic"
// Contour visualisation only on Go Playground
import (
	"strings"
)
func Pic(dx, dy int) [][]uint8 {
	p:=make([][]uint8,0,dy+1)
	for j:=0;j<dy;j++ {
		q:=make([]uint8,0,1)
		for i:=0;i<dx;i++ {
			q=append(q,uint8((i+j)/2))
		}
		p=append(p,q)
	}
	return p
}

func WordCount(s string) map[string]int {
	words:=strings.Fields(s)
	m:=make(map[string]int)
	for _,v := range(words){
		_,ok := m[v]
		if ok==true {
			m[v]+=1
		} else{
			m[v]=1
		}
	}
	return m
}


