package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Digite uma palavra:")
	fmt.Scan(&x)

	fmt.Println(strings.ToUpper(x))

}
