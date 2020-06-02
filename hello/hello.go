package main

import (
	"fmt"

	//"example.com/user/hello/morestrings"
	//"github.com/google/go-cmp/cmp"
	//"math/rand"
	"math/cmplx"
	//"golang.org/x/tour/pic"
	//"time"
)

/* Go variable types:
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

Type Conversion = T(v) like var v float64=2, int(v)
*/

var c, python, java bool //package level variables
var k float64=3.14
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
) //factored declaration
const i=3 //not :=
const big = 1 << 100
const small=big>>99
// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func add(a int,b int) int{ //or func add(a,b int) int{}
	return a+b
}
func swap(a,b int) (int,int){ //of func swap(a,b int) (x,y int){}
	return b,a
}

func main() {
	fmt.Println("Maps")
	maps()
	//pic.Show(Pic)
}