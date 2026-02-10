package test

import "fmt"

// Map que es un tipo de almacenamiento de dato
// var Mapa = map[int]string // Se define tanto el tipo de dato como los strings

var Mapa = map[int]string{}

// Para utilizarlo

var NombreEstudiantes = map[int]string{0: "Juan", 1: "Pedro"}

// Slice
var slice []string // Declaración del slice sin asignar ningún valor aún

// Array
var array [10]bool

func Colections() {

	Mapa[0] = "Colombia"
	array[0] = true                 // Asignamos al array un valor en la posición 0
	slice = append(slice, "Carlos") // Utilizamos la función global append para añadir un elemento al slice
	slice = append(slice, "Lopez")  // Asignamos un nuevo valor, este quedará en la posición [1] porque en la posición [0] ya hay un valor asignado

	fmt.Println(len(NombreEstudiantes), len(Mapa), NombreEstudiantes[0], slice[0], slice[1], Mapa[0], array[0])
}
