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

func contains(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}

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

	fistRow := (n / 3) * 3
	fistColumn := (n % 3) * 3

	for i := 0; i < 3; i++ {
		square = append(square,
			puzzle[fistRow+i][fistColumn:fistColumn+3]...)
	}

	return square
}

func checkRow(r, n int) bool {
	row := getRow(r)
	return contains(row, n)
}

func checkColumn(c, n int) bool {
	col := getColumn(c)
	return contains(col, n)
}

func checkSquare(s, n int) bool {
	sq := getSquare(s)
	return contains(sq, n)
}

func findSq(row, col int) int {
	return (row/3)*3 + (col / 3)
}

func availables(r, c int) []int {
	av := []int{}
	sq := findSq(r, c)
	for i := 1; i <= 9; i++ {
		if !checkColumn(c, i) && !checkRow(r, i) && !checkSquare(sq, i) {
			av = append(av, i)
		}
	}
	return av
}

func next(r, c int) (int, int) {
	c += 1
	for i := r; i < 9; i++ {
		for j := c; j < 9; j++ {
			if puzzle[i][j] == -1 {
				return i, j
			}
		}
		c = 0
	}
	return -1, -1
}

type cell struct {
	row    int
	column int
}

func printBoard(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == -1 {
				fmt.Print("x", "  ")
				continue
			}
			fmt.Print(board[i][j], "  ")
		}
		p()
	}
}

func solve(board [][]int) {
	r, c := -1, -1
	if puzzle[0][0] == -1 {
		r, c = 0, 0
	} else {
		r, c = next(0, 0)
	}

	prev := []cell{}

	avMap := make(map[cell][]int)

	for {

		av := []int{}
		// availables for the current cell
		if a, ok := avMap[cell{r, c}]; ok {
			av = a
		} else {
			av = availables(r, c)
		}

		if len(av) == 0 {
			if len(prev) == 0 {
				p("game over")
				break
			}
			delete(avMap, cell{r, c})
			board[r][c] = -1
			// backtrack
			r, c = prev[len(prev)-1].row, prev[len(prev)-1].column
			prev = prev[:len(prev)-1]
			continue
		}

		prev = append(prev, cell{r, c})

		if len(av) > 1 {
			board[r][c] = av[0]
			avMap[cell{r, c}] = av[1:]
			r, c = next(r, c)
		}

		if len(av) == 1 {
			board[r][c] = av[0]
			avMap[cell{r, c}] = []int{}
			r, c = next(r, c)
		}

		if r == -1 && c == -1 {
			printBoard(board)
			// p("done")
			break
		}

	}

}

func main() {
	board := make([][]int, len(puzzle))
	copy(board, puzzle)
	solve(board)

}
