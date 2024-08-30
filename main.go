package main

import ("fmt")

func sum(x int, y int) int {
	sum := x + y // Shortcut for return variable, instead of in header (sum int) {} 
	return sum
}
// returning a pack

func sub(x int, y int) (diff int, origin int) {
	

	diff = x - y
	origin = 10
	return 

}
func main() {
	var s string = "Let's Go!"
	

	pack1, pack2 := sub(20, 10)

	fmt.Println("Hello ", pack1, pack2) 
	arra :=[4]int{1, 2, 3, 4}

	// var S int= 10; public
	// var s int = 10; private

	// i := 10; infered type

	for i := 0; i < len(arra); i ++ {

		fmt.Println(arra[i])
	}

	j := 0
	for j < len(arra) {

		fmt.Println(arra[j])
		j++
	}


	// Index value,
	for i, v := range arra {
		fmt.Println(i, v)
	}


	// Struct

	type Animal struct {
		Name string  // public
		animalType string // private
	}

	shark := Animal{"Foo", "Type1"}
	fmt.Println(shark)
	fmt.Println(s);
}
