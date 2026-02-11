package main

import (
	"fmt"
	"github.com/juancholopes/go-course/test"
)

const float = multiplicador
const url = test.TestEnpoint

var salary float64 = 20000000.0

func main() {
	test.Mapa[0] = "Estados Unidos"
	print("Hello from a module")

	fmt.Println("Taxes calculate: ", test.CalculateTaxes(23000))

	test.Colections()
	SalaryAfterTaxes, SalaryBeforeTaxes := test.CalculateSalary(salary)

	fmt.Println("\n", y, "\n", entero, "\n", otroNumero, "\n", calcular(), "\n", string, float, url, test.Mapa[0])
	fmt.Println("Salary before taxes: ", SalaryBeforeTaxes)
	fmt.Println("Salary after taxes: ", SalaryAfterTaxes)
}
