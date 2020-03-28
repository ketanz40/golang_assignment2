package main

import (
	//"fmt"
	"assignment2/q1"
	"assignment2/q2"
	"assignment2/q3"
)

func main(){
	q1.Main1()
	defer q2.PrintValues()
	q2.RunQ2()
	q3.Mainq3()
}
