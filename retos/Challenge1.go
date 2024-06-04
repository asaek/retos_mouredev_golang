package retos

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * Reto #1
 * ¿ES UN ANAGRAMA?
 * Fecha publicación enunciado: 03/01/22
 * Fecha publicación resolución: 10/01/22
 * Dificultad: MEDIA
 *
 * Enunciado: Escribe una función que reciba dos palabras (String) y retorne verdadero o falso (Boolean) según sean o no anagramas.
 * Un Anagrama consiste en formar una palabra reordenando TODAS las letras de otra palabra inicial.
 * NO hace falta comprobar que ambas palabras existan.
 * Dos palabras exactamente iguales no son anagrama.
 */

func Challenge1(palabra1 string, palabra2 string) {
	if palabra1 == palabra2 {
		fmt.Println("Las Palabras son iguales por lo tanto no son un anagrama")
		return
	}

	cadenaRunas1 := []rune(strings.ToLower(palabra1))
	cadenaRunas2 := []rune(strings.ToLower(palabra2))

	sort.Slice(cadenaRunas1, func(i, j int) bool {
		return cadenaRunas1[i] < cadenaRunas1[j]
	})

	sort.Slice(cadenaRunas2, func(i, j int) bool {
		return cadenaRunas2[i] < cadenaRunas2[j]
	})

	if string(cadenaRunas1) == string(cadenaRunas2) {
		println("Si son anagramas")
	}

}
