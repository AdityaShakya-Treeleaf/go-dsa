package dsa

import (
	"log"
)

//Increment Sub matrices by One
//You are given a positive integer n, indicating that we initially have an n x n 0-indexed integer matrix mat filled with zeroes.

//You are also given a 2D integer array query. For each query[i] = [row1i, col1i, row2i, col2i], you should do the following operation:

//Add 1 to every element in the submatrix with the top left corner (row1i, col1i) and the bottom right corner (row2i, col2i). That is, add 1 to mat[x][y] for all row1i <= x <= row2i and col1i <= y <= col2i.
//Return the matrix mat after performing every query.

//Example 1:

//Input: n = 3, queries = [[1,1,2,2],[0,0,1,1]]
//Output: [[1,1,0],[1,2,1],[0,1,1]]
//Explanation: The diagram above shows the initial matrix, the matrix after the first query, and the matrix after the second query.
//- In the first query, we add 1 to every element in the sub matrix with the top left corner (1, 1) and bottom right corner (2, 2).
//- In the second query, we add 1 to every element in the sub matrix with the top left corner (0, 0) and bottom right corner (1, 1).

type IncrementSubMatricesOne struct{}

func (i IncrementSubMatricesOne) RangeAddQueries(n int, queries [][]int) [][]int {
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("Panic: %s", rec)
		}
	}()
	if n == 0 {
		return [][]int{}
	}

	baseMat := make([][]int, n)
	for i := 0; i < n; i++ {
		baseMat[i] = make([]int, n)
	}

	for i := 0; i < len(queries); i++ {
		topLeftX := queries[i][0]
		topLeftY := queries[i][1]
		botRightX := queries[i][2]
		botRightY := queries[i][3]
		for j := topLeftX; j <= botRightX; j++ {
			for k := topLeftY; k <= botRightY; k++ {
				baseMat[j][k] += 1
			}
		}
	}

	return baseMat
}
