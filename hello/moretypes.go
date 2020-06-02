package main

import (
	"fmt"
)

func pointer(){
	i,j:=2,3
	p:=&i
	fmt.Println(*p)
	*p=21
	fmt.Println(i)

	p=&j
	*p=*p/3
	fmt.Println(j)
}

type grid struct {
	x int
	y int
}

func structure(){
	v := grid{1,2}
	fmt.Println(v)
	fmt.Println(v.x+1)
	p:=&v
	p.x=4 //No need of (*p).x
	w:=grid{x:5} //y:0 implicit
	fmt.Println(w)
}

func arrays(){
	arr:= [10]int{1,2,3,4,5,6,7}
	fmt.Println(arr)
	s:=arr[1:4] //slice is dynamic in size. They are references, so change in them changes og array
	fmt.Println(s);
	s2:=[]int{1,2}
	fmt.Println(s2)
	//s=s2[:]
	fmt.Println(s,cap(s),len(s))//capacity from first element of slice 
	a:=make([]int,2,5) //make([]type,len,cap)
	a[1]=4
	fmt.Println(a,cap(a),len(a))
	/*board := [][]string{ // just for example
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)*/
	s=append(s,2,3,4,3,4,6,7)
	fmt.Println(s,cap(s),len(s),arr) //read how new capacity allocated. new array is allocated.
	for i,v := range a {
		fmt.Println(i, v)
	}
	
}