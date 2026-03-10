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
	// COLLECTIONS
	fmt.Println("COLLECTIONS")
	collections.Mapa[0] = "Estados Unidos"
	collections.Colections()

	// VISIBILITY
	fmt.Println("VISIBILITY")
	fmt.Println("Usar funciones privadas:")
	fmt.Println(visibility.UsarVarloresPrivados())
	fmt.Println("Exportar función privada:")
	fmt.Println(visibility.ExportarFuncionPrivada())
	// Extraer valores de  una función que retorna multiples valores a la vez
	miEnteroPrivado, miSlicePrivado, miMapaPrivado, miArrayPrivado := visibility.UsarVarloresPrivados()

	valor := visibility.MultiplicarPorDos(4)
	fmt.Println(valor)
	fmt.Println(miEnteroPrivado, miSlicePrivado[0], miMapaPrivado["id"], miArrayPrivado[0])

	// POINTERS
	fmt.Println("POINTERS")
	age := 19
	pointers.AgeCalculator(&age)
	fmt.Println(age)

	//FUNCTIONS
	SalaryAfterTaxes, SalaryBeforeTaxes := functions.CalculateSalary(salary)

	fmt.Println("\n", "\n", types.Entero, "\n", types.OtroNumero, "\n", types.Calcular(), "\n", float, collections.Mapa[0])
	fmt.Println("Salary before taxes: ", SalaryBeforeTaxes)
	fmt.Println("Salary after taxes: ", SalaryAfterTaxes)
}
