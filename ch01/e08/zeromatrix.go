// Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
// column are set to 0.
package e08

import "errors"

func zeroMatrixWithArray(matrix [][]int) error {
	if len(matrix) == 0 {
		return errors.New("Invalid matrix of length 0.")
	}

	rowSize := len(matrix)
	columnSize := len(matrix[0])
	rows := make([]bool, rowSize)
	columns := make([]bool, columnSize)

	for i := 0; i < rowSize; i++ {
		for j := 0; j < columnSize; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	for i := 0; i < rowSize; i++ {
		if rows[i] {
			for j := 0; j < columnSize; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < columnSize; i++ {
		if columns[i] {
			for j := 0; j < rowSize; j++ {
				matrix[j][i] = 0
			}
		}
	}

	return nil
}

func zeroMatrixInPlace(matrix [][]int) error {
	if len(matrix) == 0 {
		return errors.New("Invalid matrix of length 0.")
	}

	rowSize := len(matrix)
	columnSize := len(matrix[0])
	rowHasZero := false
	columnHasZero := false

	for i := 0; i < rowSize; i++ {
		if matrix[i][0] == 0 {
			rowHasZero = true
			break
		}
	}

	for i := 0; i < columnSize; i++ {
		if matrix[0][i] == 0 {
			columnHasZero = true
			break
		}
	}

	for i := 1; i < rowSize; i++ {
		for j := 1; j < columnSize; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < columnSize; i++ {
		if matrix[0][i] == 0 {
			zeroColumn(matrix, i, rowSize)
		}
	}

	for i := 1; i < rowSize; i++ {
		if matrix[i][0] == 0 {
			zeroRow(matrix, i, columnSize)
		}
	}

	if rowHasZero {
		zeroRow(matrix, 0, columnSize)
	}

	if columnHasZero {
		zeroColumn(matrix, 0, rowSize)
	}

	return nil
}

func zeroColumn(matrix [][]int, column int, rowsLength int) {
	for i := 0; i < rowsLength; i++ {
		matrix[i][column] = 0
	}
}

func zeroRow(matrix [][]int, row int, columnsLength int) {
	for i := 0; i < columnsLength; i++ {
		matrix[row][i] = 0
	}
}
