package q1

import "fmt"

var WhatIsThe int = AnswerToLife()

func AnswerToLife() int{
	return 42
}

func init(){
	WhatIsThe = 0
}



func Main1(){
	if WhatIsThe == 0{
		fmt.Println("It's all a lie")	
	}
}