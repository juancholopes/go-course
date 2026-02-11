package visibility

/*
Go maneja diferentes scopes:
	- Scope de paquete => Scope global (en otros lenguajes como JavaScript)
	- Scope de función => func
	- Scope de bloque => if, while, etc.

Cuando se declaran funciones o variables en el Scope de paquete, estas funciones
y variables se pueden usar en otros archivos SIEMPRE Y CUANDO esten en el mismo
directorio o mejor dicho paquete

Para usar este paquete y sus funciones en otros paquetes fuera del mismo directorio,
los podemos exportar dependiendo de como se capitalice su nombre, por ejemplo una variable
que inicia con minúsculas y es declarada en el scope de paquete es una variable privada
si inicia con mayúsculas es una variable pública - LO MISMO OCURRE CON LAS FUNCIONES.

*/

// Variable en el scope de paquete privada

var miVariablePrivada = 893298372
var MiVariablePublica = "Esto es una variable pública"

// También con constantes 
const miConstantePrivada = 2006 
const MiConstantePública = 19

// También con arrays - mapas - slice 
var MiMapaPublico = map[int]string{} // Inicializado - No usado aún 
var miMapaPrivado = map[string]int{"id": 10959860000}

var MiArrayPublico [10]int 
var miArrayPrivado = [...]int{1,2,3,4,5,6}

var MiSlicePublico []int
var miSlicePrivado []string // Recordar que para añadir a un slice se usa la función append()

// Función privada para exportar pero usable en el mismo paquete
func dividirEntreTres(numero int) int { 
	return numero * 3
}

// Función pública que se puede usar tanto en el paquete como cuando se exporta 
func MultiplicarPorDos (numero int) int {
	return numero * 2
}

// Podemos crear funciones públicas que usen esos valores privados 

func UsarVarloresPrivados () (int, []string, map[string]int, [6]int){
	//Usamos mapa privado
	miMapaPrivado["id"] = 12939032
	miSlicePrivado = append(miSlicePrivado, "Colombia")

	resultadoDividirEntreTres := dividirEntreTres(miVariablePrivada)
	return miConstantePrivada - resultadoDividirEntreTres , miSlicePrivado, miMapaPrivado, miArrayPrivado
}
