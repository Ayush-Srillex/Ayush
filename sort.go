package main

import(
	"fmt"
	"sort"
)

func cosort(a []int,c chan []int){
	sort.Ints(a)
	c<-a
}

func merge(a []int,b []int,c chan []int) {
	n1,n2:=len(a),len(b)
	s:=make([]int,0)
	l,m:=0,0
	for l<n1 && m<n2 {
		if a[l]<b[m] {
			s=append(s, a[l])
			l++
		} else{
			s=append(s, b[m])
			m++
		}
	}
	if l<n1{
		for l<n1 {
			s=append(s, a[l])
			l++
		}
	} else{
		for m<n2 {
			s=append(s, b[m])
			m++
		}
	}
	c<-s
}
func main() {
	//Input Elements
	fmt.Println("Enter no. of elements")
	var n int
	fmt.Scan(&n)
	fmt.Println("Enter elements")
	a:=make([]int, n)
	for i := 0; i<n; i++ {
		fmt.Scan(&a[i])
	}

	//Dividing in 4 arrays
	a1:=a[:n/4]
	a2:=a[n/4:n/2]
	a3:=a[n/2:3*n/4]
	a4:=a[3*n/4:]
	
	//sorting in concurrency
	c:=make(chan []int,4)
	go cosort(a1,c)
	go cosort(a2,c)
	go cosort(a3,c)
	go cosort(a4,c)

	//retrieve arrays
	aa1:=<-c
	aa2:=<-c
	aa3:=<-c
	aa4:=<-c

	//Print Individual sorted arrays
	//Maynot be original sequence because of no deterministic nature of concurrency
	fmt.Println(aa1,aa2,aa3,aa4)

	//merging
	go merge(aa1,aa2,c)
	go merge(aa3, aa4,c)
	s1:=<-c
	s2:=<-c
	go merge(s1,s2,c)
	s:=<-c

	//Final array
	fmt.Println(s)
	fmt.Println("Enter any number to exit")
	var exit int
	fmt.Scan(&exit)
}