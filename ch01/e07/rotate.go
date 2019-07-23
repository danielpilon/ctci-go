// Given an image represented by an NxN matrix, where each pixel in the image is 4
// bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
package e07

import "errors"

func rotateMatrix(matrix [][]int) error {
	if len(matrix) == 0 {
		return errors.New("Invalid matrix of length 0.")
	} else if len(matrix) != len(matrix[0]) {
		return errors.New("Invalid matrix does not have the same length for each dimension.")
	}
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]
			// left -> top
			matrix[first][i] = matrix[last-offset][first]
			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]
			// right -> bottom
			matrix[last][last-offset] = matrix[i][last]
			// top ->  right
			matrix[i][last] = top
		}
	}

	return nil

}
