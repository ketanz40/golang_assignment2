package q3

import "fmt"

func Mainq3() {
	var word string = "This is a recursive call function"
	recursiveCall(word, 0)
}

func recursiveCall(s string, n int) rune{
	runes := []rune(s)
	var r rune
	if n == len(s) - 1{
		r = runes[n]
		return r
	}
	r = runes[n]
	fmt.Printf("%x\n", r)	
	return recursiveCall(s, n+1)
}


