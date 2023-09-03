package main

import (
	"errors"
	"fmt"
)

func matrixAdd(a [][]int, b [][]int) ([][]int, error) {
	if len(a) == len(b) && len(a[0]) == len(b[0]) {
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[0]); j++ {
				a[i][j] = a[i][j] + b[i][j]
			}
		}
	} else {
		return [][]int{{}}, errors.New("matrixes are not the same size")
	}
	return a, nil // nil - местный null
}

func calculateCellValue(a []int, b []int) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("slices are different size")
	}
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i] * b[i]
	}
	return total, nil
}

func createVerticalSlice(matrix [][]int, number int) []int {
	vertSlice := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		vertSlice[i] = matrix[i][number]
	}
	return vertSlice
}

func matrixMultiply(a [][]int, b [][]int) ([][]int, error) {
	if len(a[0]) == len(b) {
		newMatrix := make([][]int, len(a))
		for i := 0; i < len(a); i++ {
			newMatrix[i] = make([]int, len(b[0]))
			for j := 0; j < len(b[0]); j++ {
				// just some error handling
				value, err := calculateCellValue(a[i], createVerticalSlice(b, j))
				if err != nil {
					return nil, err
				} else {
					newMatrix[i][j] = value
				}
			}
		}
		return newMatrix, nil
	} else {
		return [][]int{{}}, errors.New("matrixes sizes do not match (m * n) * (n * k)")
	}
}

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	b := [][]int{{1}, {2}, {3}}
	fmt.Print(matrixMultiply(a, b))
}
