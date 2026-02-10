package main

import (
	"fmt"
	"github.com/juancholopes/go-course/test"
)

const float = multiplicador
const url = test.TestEnpoint

func main() {
	test.Mapa[0] = "Estados Unidos"
	print("Hello from a module")
	fmt.Println("\n", y, "\n", entero, "\n", otroNumero, "\n", calcular(), "\n", string, float, url, test.Mapa[0])
	test.Colections()
}
