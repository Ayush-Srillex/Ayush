package main

import (
	"fmt"
	"runtime"
	"time"
)

func loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	
}

func whileloop() { //for is Go's while
	sum:=0
	i:=0
	
	for i<100 {
		if  i%2==0 { //read if initialiser
			sum=sum+i
		}
		
		i+=1
	}
	fmt.Println(sum)
}

func sqrt(x float64) float64{
	var g float64 = 1
	for ((g*g)-x)>0.0000001 || (x-(g*g))>0.0000001 {
		g-=(g*g-x)/(2*g)
	}
	return g
}

func os(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func greet(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}