package main

import (
	"fmt"
	"search_matrix"
)

func main() {

	var pixel int
	fmt.Println("Ingrese el pixel a buscar:")
	fmt.Scanf("%d", &pixel)
	i, j := search_matrix.SearchPixel(pixel)
	fmt.Printf("La posición es [%d][%d] para el valor %d \n", i, j, pixel)

}
