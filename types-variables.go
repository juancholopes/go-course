package main

const y = 3 // son constantes No inmutables es decir su valor no cambia en el tiempo y solo recibe strings booleanos y enteros

var entero int // por defecto es nil que es el equivalente a null 
var otroNumero int = 0 // también se puede asiganar el valor de una vez 

// Cuando inicializamos la varible con un valor 
var string = "hola" // Infiere también el tipo 
var anotherNumber = 2 // Infiere el tipo 
var pointer = anotherNumber

// var iniciarlizarSinTipo := "hola" // Solo funciona dentro de funciones


const multiplicador = 100000 // Go lo que en escencia hace en tiempo de compilación es copiar y pegar ese valor en el lugar donde se llame la variable


func calcular() int {
	numero := 10 // Infiere el tipo pero solo se usa dentro del scope de funciones no el scope global 
	return numero * multiplicador / anotherNumber * pointer // Aquí en realidad Go remplaza por el valor 10000 o el valor que tenga asignado 
}

