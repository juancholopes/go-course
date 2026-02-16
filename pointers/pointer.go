package pointers

/*
En Go en las funciones cuando reciben un parámetro este esta por defecto toma una copia del valor para ser utilizado
dentro de la función pero no modifica el valor original, pero si queremos modificar el valor de la variable lo que se suele
utilizar son los punteros que permiten modificar el valor de una función pasando la dirección en memoria de la variable
*/

// Esta función recibe como parámetro la dirección en memoria de la variable
func AgeCalculator(age *int) {
	*age++ // Después toma el valor de la variable y lo incrementa * => significa que va a tomar el valor no su dirección
}
