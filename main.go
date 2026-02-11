package main

import (
	"fmt"

	"github.com/juancholopes/go-course/collections"
	"github.com/juancholopes/go-course/functions"
	"github.com/juancholopes/go-course/pointers"
	"github.com/juancholopes/go-course/types"
	"github.com/juancholopes/go-course/visibility"
)

const float = types.Multiplicador

var salary float64 = 20000000.0

func main() {
	collections.Mapa[0] = "Estados Unidos"
	print("Hello from a module")

	fmt.Println("Taxes calculate: ", functions.CalculateTaxes(23000))

	// VISIBILITY
	fmt.Println("Usar funciones privadas:")
	fmt.Println(visibility.UsarVarloresPrivados())
	// Extraer valores
	miEnteroPrivado, miSlicePrivado, miMapaPrivado, miArrayPrivado := visibility.UsarVarloresPrivados()

	fmt.Println(miEnteroPrivado, miSlicePrivado[0], miMapaPrivado["id"], miArrayPrivado[0])

	// POINTERS
	age := 19
	pointers.AgeCalculator(&age)
	fmt.Println(age)

	collections.Colections()
	SalaryAfterTaxes, SalaryBeforeTaxes := functions.CalculateSalary(salary)

	fmt.Println("\n", "\n", types.Entero, "\n", types.OtroNumero, "\n", types.Calcular(), "\n", float, collections.Mapa[0])
	fmt.Println("Salary before taxes: ", SalaryBeforeTaxes)
	fmt.Println("Salary after taxes: ", SalaryAfterTaxes)
}
