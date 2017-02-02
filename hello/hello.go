package main

import (

	"fmt"
	"github.com/mattrice12/go-practice/stringutil"

)

func main() {
	var thisVariable float32

	fmt.Println(stringutil.Reverse("!oG olleH\n"))
	fmt.Println("!oG olleH"[1])
	fmt.Println("!oG" + " olleH")
	fmt.Println(thisVariable)
	fmt.Println("1 + 1 =", 1 + 1)

	thisVariable = 2.0 + 1.3

	fmt.Println(thisVariable)

}
