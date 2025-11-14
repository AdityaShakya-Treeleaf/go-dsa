package utils

import (
	"reflect"
	"slices"
)

// slowers
func AreMatricesEqual(mat1, mat2 [][]int) bool {
	return reflect.DeepEqual(mat1, mat2)
}

func AreMatricesEqualV2(mat1, mat2 [][]int) bool {
	if len(mat1) != len(mat2) {
		return false
	}

	for i := range mat1 {
		if !slices.Equal(mat1[i], mat2[i]) {
			return false
		}
	}

	return true
}
