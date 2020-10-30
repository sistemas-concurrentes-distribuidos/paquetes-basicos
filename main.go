package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	condicion := "s"
	var nombre string

	var nombres []string

	for strings.ToLower(condicion) != "n" {
		fmt.Println("Ingresa el nombre: ")
		fmt.Scanln(&nombre)
		fmt.Println()
		nombres = append(nombres, nombre)

		fmt.Println("¿Agregar otro nombre? Sí(cualquier tecla)/No(n)")
		fmt.Scanln(&condicion)
		fmt.Println()
	}

	sort.Strings(nombres)
	guardarArchivo("asecendente.txt", nombres)

	sort.Sort(sort.Reverse(sort.StringSlice(nombres)))
	guardarArchivo("descendente.txt", nombres)
}

func guardarArchivo(nombreArchivo string, nombres []string) {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", nombreArchivo)
		return
	}
	defer archivo.Close()
	_, errEscribiendoArchivo := archivo.WriteString(strings.Join(nombres, "\n"))
	if errEscribiendoArchivo != nil {
		fmt.Println("Error al escribir en el archivo: ", nombreArchivo)
	}

}
