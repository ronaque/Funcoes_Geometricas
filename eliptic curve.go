package main

import (
	"fmt"
	"math"
)

func equacaoSegundoGrauMatrix() {
	const tam = 11
	var matrix [tam][tam]string
	var x, y float64

	a, b := 1.0, 1.0
	c := 0.0

	// Preencher Matriz
	for i := tam - 1; i >= 0; i-- {
		for j := 0; j < tam; j++ {
			matrix[i][j] = fmt.Sprintf("%dx%d ", i, j)
			// matrix[i][j] = " "
		}
	}

	// x, y = 0.0, 0.0
	// matrix[int(x)][int(y)] = "a\t"
	// fmt.Println(matrix[0][10])
	for x = 0.00; x <= 2.00; x += 0.01 {
		y = a*(math.Pow(x, 2)) + (b * x) + c
		// y = math.Sqrt((a * math.Pow(x, 3)) + a*x + b)
		// fmt.Printf("x: %f x*10: %f intx: %d y: %f y*10: %f inty: %d\n", x, x*10, int(x*10), y, y*10, int(y*10))
		// fmt.Printf("x*10: %f intx: %d y*10: %f inty: %d\n", x*10, int(x*10), y*10, int(y*10))
		matrix[(tam/2)+int(y)][(tam/2)+int(x)] = "a   "
	}

	//i = y; j = x
	// Printar Matriz
	for i := tam - 1; i >= 0; i-- {
		for j := 0; j < tam; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func equacaoSegundoGrau() {
	var x, y float64

	a, b, c := 1.0, 1.0, 0.0

	for x := 2.00; x >= -2.00; x -= 0.10 {
		y = a*(math.Pow(x, 2)) + (b * x) + c
		for auxY := math.Abs(y * 10); auxY > 0.0; auxY-- {
			print(" ")
		}
		print(".\n")
		// fmt.Printf("%0.3f %0.3f\n", x, y)
	}

	fmt.Printf("X: %0.3f Y: %0.3f\n", x, y)
}

func elipticCurveMatrix() {

}

func elipticCurve() {
	var x, y float64

	a, b := 1.0, 1.0

	for x := 2.00; x >= -0.50; x -= 0.10 {
		y = math.Sqrt(math.Pow(x, 3) + (a * x) + b)
		for auxY := math.Abs(y * 10); auxY > 0.0; auxY-- {
			print(" ")
		}
		print(".\n")
		// fmt.Printf("%0.3f %0.3f\n", x, y)
	}

	fmt.Printf("X: %0.3f Y: %0.3f\n", x, y)
}

func main() {
	// equacaoSegundoGrau()
	// elipticCurve()
	equacaoSegundoGrauMatrix()
}
