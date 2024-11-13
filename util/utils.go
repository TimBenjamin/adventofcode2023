package util

import "errors"

func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// Pair of functions that return a specific column from a string or int 2D slice.
func GetIntMatrixColumn(matrix [][]int, colIndex int) ([]int, error) {
	// Check if matrix is empty or colIndex is out of bounds
	if len(matrix) == 0 || colIndex < 0 || colIndex >= len(matrix[0]) {
		return nil, errors.New("invalid column index or empty matrix")
	}

	column := make([]int, len(matrix)) // Slice to store the column elements
	for i := 0; i < len(matrix); i++ {
		column[i] = matrix[i][colIndex] // Extract each element from the column
	}

	return column, nil
}

func GetStringMatrixColumn(matrix [][]string, colIndex int) ([]string, error) {
	// Check if matrix is empty or colIndex is out of bounds
	if len(matrix) == 0 || colIndex < 0 || colIndex >= len(matrix[0]) {
		return nil, errors.New("invalid column index or empty matrix")
	}

	column := make([]string, len(matrix)) // Slice to store the column elements
	for i := 0; i < len(matrix); i++ {
		column[i] = matrix[i][colIndex] // Extract each element from the column
	}

	return column, nil
}

// GetStringMatrixAdjacent takes a 2D slice of strings, an index (i, j), and a direction, and returns the string in the specified direction if it's within bounds.
func GetStringMatrixAdjacent(matrix [][]string, i, j int, direction string) (string, error) {
	// Check if the matrix or index is out of bounds
	if len(matrix) == 0 || i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
		return "", errors.New("invalid initial index")
	}

	// Based on direction, calculate new indices and check bounds
	switch direction {
	case "up":
		if i-1 >= 0 {
			return matrix[i-1][j], nil
		}
	case "down":
		if i+1 < len(matrix) {
			return matrix[i+1][j], nil
		}
	case "left":
		if j-1 >= 0 {
			return matrix[i][j-1], nil
		}
	case "right":
		if j+1 < len(matrix[i]) {
			return matrix[i][j+1], nil
		}
	}

	return "", errors.New("specified direction goes out of bounds")
}
