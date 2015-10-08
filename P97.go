// Solve Sudoku puzzles
// Sudoku puzzles go like this:

//    Problem statement                 Solution

//     .  .  4 | 8  .  . | .  1  7	    9  3  4 | 8  2  5 | 6  1  7
//             |         |                      |         |
//     6  7  . | 9  .  . | .  .  .	    6  7  2 | 9  1  4 | 8  5  3
//             |         |                      |         |
//     5  .  8 | .  3  . | .  .  4      5  1  8 | 6  3  7 | 9  2  4
//     --------+---------+--------      --------+---------+--------
//     3  .  . | 7  4  . | 1  .  .      3  2  5 | 7  4  8 | 1  6  9
//             |         |                      |         |
//     .  6  9 | .  .  . | 7  8  .      4  6  9 | 1  5  3 | 7  8  2
//             |         |                      |         |
//     .  .  1 | .  6  9 | .  .  5      7  8  1 | 2  6  9 | 4  3  5
//     --------+---------+--------      --------+---------+--------
//     1  .  . | .  8  . | 3  .  6	    1  9  7 | 5  8  2 | 3  4  6
//             |         |                      |         |
//     .  .  . | .  .  6 | .  9  1	    8  5  3 | 4  7  6 | 2  9  1
//             |         |                      |         |
//     2  4  . | .  .  1 | 5  .  .      2  4  6 | 3  9  1 | 5  7  8

package main

import "fmt"

var p = fmt.Println

var puzzle = [][]int{
	[]int{-1, -1, 4, 8, -1, -1, -1, 1, 7},
	[]int{6, 7, -1, 9, -1, -1, -1, -1, -1},
	[]int{5, -1, 8, -1, 3, -1, -1, -1, 4},
	[]int{3, -1, -1, 7, 4, -1, 1, -1, -1},
	[]int{-1, 6, 9, -1, -1, -1, 7, 8, -1},
	[]int{-1, -1, 1, -1, 6, 9, -1, -1, 5},
	[]int{1, -1, -1, -1, 8, -1, 3, -1, 6},
	[]int{-1, -1, -1, -1, -1, 6, -1, 9, 1},
	[]int{2, 4, -1, -1, -1, 1, 5, -1, -1}}

func getColumn(n int) []int {
	column := []int{}

	for _, v := range puzzle {
		column = append(column, v[n])
	}

	return column
}

func getRow(n int) []int {
	return puzzle[n]
}

func getSquare(n int) []int {
	square := []int{}

	return square
}

func main() {
	p(getRow(3))
}
