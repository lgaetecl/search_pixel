package search_matrix

import "fmt"

var (
	mat [16][16]int
	k   int = 0
)

func init() {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			if i == 0 || i%2 == 0 {
				mat[i][j] = k
			} else {
				mat[i][15-j] = k
			}
			k++
		}
	}
}

func SearchPixel(pixel int) (int, int) {
	var x, y int
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			if mat[i][j] == pixel {
				y = i
				x = j
			}
		}
	}
	return x, y
}

func main() {

	var pixel int
	fmt.Println("Ingrese el pixel a buscar:")
	fmt.Scanf("%d", &pixel)
	i, j := SearchPixel(pixel)
	fmt.Printf("La posición es [%d][%d] para el valor %d \n", i, j, pixel)

}
