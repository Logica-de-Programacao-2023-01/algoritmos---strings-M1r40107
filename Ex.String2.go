package main

import "fmt"

func main() {
	var s string

	fmt.Print("Digite um String:")
	fmt.Scan(&s)
	fmt.Println("len(s) =", len(s))

}
