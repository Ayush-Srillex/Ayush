package main

import (
	"fmt"

	//"example.com/user/hello/morestrings"
	//"github.com/google/go-cmp/cmp"
	"math/rand"
	//"math/cmplx"
	"time"
)

func basic() {
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(add(2,3))
	fmt.Println(f(), f(), f(), f(), f())
	//time.Now().UnixNano() generates constantly changing numbers
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	//fmt.Println(time.Now().UnixNano())
	fmt.Println("Random number",rand.Intn(100))
	a:=2
	b:=3
	x,y:=swap(a,b)
	fmt.Println(a,b," Swap ",x,y)
	fmt.Println(int(k))
	fmt.Println(small,big*0.1)
	loop()
	whileloop()
	fmt.Println(sqrt(400))
	os()
	greet()
	/*Defer
	Deferred function calls are pushed onto a stack.
	When a function returns, its deferred calls are executed
	in last-in-first-out order.
	*/
	defer fmt.Println("world")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	//fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	//fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}