package main

import "fmt"

// Num struct
type Num struct {
	N int
}

func main() {
	// Variables -- reference variable(a) when declaraing pointer(p); asterisk pointer(*p) when manipulating it
	a := 1
	p := &a
	*p = 3
	fmt.Println("a(Var):", a)
	fmt.Println("")

	b := Num{2}
	b = squareVal(b)
	fmt.Println("b1(Val):", b.N)

	squareRef(&b)
	fmt.Println("b1(Ref):", b.N)
	fmt.Println("")

	c := 3
	c = doubleVal(c)
	fmt.Println("c1(Val):", c)

	doubleRef(&c)
	fmt.Println("c1(Ref):", c)
	fmt.Println("")

	d := quadVal(b)
	fmt.Println("d2(Val):", d.N)

	quadRef(&d)
	fmt.Println("d2(Ref):", d.N)
}

// Struct Types Value -- No pointing; need a return value
func squareVal(n Num) Num {
	n.N = n.N * n.N
	return n
}

// Struct Types Reference -- Only need to point to type declaration for structs
func squareRef(n *Num) {
	n.N = n.N * n.N
}

// Built-in types (int, string, bool ) -- Values need to be returned
func doubleVal(n int) int {
	n = n + n
	return n
}

// Built-in types (int, string, bool) -- Need to point to type declaration and all instances of that type's variable; no return value needed
func doubleRef(n *int) {
	*n = *n + *n
}

// Struct Types Value -- No pointing; need return
func quadVal(n Num) Num {
	n.N = n.N * 4
	return n
}

// Struct Types Reference -- Pointing; no return
func quadRef(n *Num) {
	n.N = n.N * 4
}
