//Q2
package q2

import (
	"fmt"
)

var x int
var y int

var sum int
var sub int
var mul int
var quo float32

func RunQ2(){
	fmt.Print("Enter two numbers: ")
	fmt.Scanf("%d %d", &x, &y)
	sum, sub, mul, quo = calculate(x,y)
}

func calculate(x ,y int) (sum, sub, mul int, quo float32){
	sum = x + y
	sub = x - y
	mul = x * y
	quo = (float32(x) / (float32(y)))
	return
}

func PrintValues(){
	fmt.Println("Sum:", sum,"Difference:",sub,"Product:",mul,"Quotient:",quo)
}