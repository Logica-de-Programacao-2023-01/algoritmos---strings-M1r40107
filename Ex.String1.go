package main

import "fmt"

func main() {
	var x, y string

	fmt.Print("Digite uma String:")
	fmt.Scan(&x)
	fmt.Print("Digite outra String:")
	fmt.Scan(&y)

	s := x + ", " + y
	fmt.Println(s)
}
