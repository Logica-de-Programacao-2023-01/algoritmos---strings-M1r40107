package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, y string
	fmt.Print("Digite um String:")
	fmt.Scan(&x)
	fmt.Print("Digite o caractere desejado:")
	fmt.Scan(&y)

	count := strings.Count(x, y)
	fmt.Printf("o caractere '%s' aparece '%d' vezes na string '%s' .\n", y, count, x)
}
