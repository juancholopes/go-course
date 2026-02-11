package types

// CONSTANTES
const y = 3 // son constantes No inmutables es decir su valor no cambia en el tiempo y solo recibe strings booleanos y enteros

// VARIABLES
var Entero int         // por defecto es nil que es el equivalente a null
var OtroNumero int = 0 // también se puede asiganar el valor de una vez

// Cuando inicializamos la varible con un valor
var String = "hola"   // Infiere también el tipo
var AnotherNumber = 2 // Infiere el tipo
var Pointer = AnotherNumber

// var iniciarlizarSinTipo := "hola" // Solo funciona dentro de funciones

const Multiplicador = 100000 // Go lo que en escencia hace en tiempo de compilación es copiar y pegar ese valor en el lugar donde se llame la variable

func Calcular() int {
	// INICIALIZADOR
	Numero := 10 // Infiere el tipo pero solo se usa dentro del scope de funciones no el scope global
	return Numero * Multiplicador / AnotherNumber * Pointer - y // Aquí en realidad Go remplaza por el valor 10000 o el valor que tenga asignado
}
