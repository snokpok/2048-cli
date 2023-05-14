package main

import (
	"fmt"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type BoardCell struct {
	r int
	c int
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
}

// insert into the board a random five; returns non-nil error and nil cell if
// there's no blank cells left
func insertRandomFive(board [][]string) (*BoardCell, error) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "_" {
				board[i][j] = "5"
				return &BoardCell{i, j}, nil
			}
		}
	}
	return nil, fmt.Errorf("No blank cells left!")
}

func validCell(board [][]string, cell BoardCell) bool {
	return cell.r >= 0 && cell.r < getNumRows(board) && cell.c >= 0 && cell.c < getNumCols(board)
}

func getNumRows(board [][]string) int {
	return len(board)
}

func getNumCols(board [][]string) int {
	return len(board[0])
}

// game lost if all cells have no neighbor with the same value (i.e can't make
// any move on any cell)
func isGameLost(board [][]string) bool {
	for i := 0; i < getNumRows(board); i++ {
		for j := 0; j < getNumCols(board); j++ {
			neighbors := []BoardCell{
				{i, j + 1},
				{i + 1, j},
				{i - 1, j},
				{i, j - 1},
			}
			for _, neighbor := range neighbors {
				if validCell(board, neighbor) && board[i][j] == board[neighbor.r][neighbor.c] {
					return false
				}
			}
		}
	}
	return true
}

// merge paired duplicates; from = "start" -> start merging from the start of
// the slice.
// from = "end" -> start merging from the end of the slice.
func mergePairDuplicates(arr []string, from string) []string {
	// do intermediate merging (e.g 5+5 -> 10)
	switch from {
	case "start":
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] == arr[i+1] {
				first, err := strconv.Atoi(arr[i])
				check(err)
				arr[i] = fmt.Sprintf("%d", first*2)
				// deleting the upper element
				copy(arr[i+1:], arr[i+2:])
				arr = arr[:len(arr)-1]
                i++ // move to the next pair immediately
			}
		}
	case "end":
		for i := len(arr)-2; i >= 0; i-- {
			if arr[i] == arr[i+1] {
				first, err := strconv.Atoi(arr[i])
				check(err)
				arr[i] = fmt.Sprintf("%d", first*2)
				// deleting the upper element
				copy(arr[i+1:], arr[i+2:])
				arr = arr[:len(arr)-1]
                i-- // move to the next pair immediately
			}
		}
	}
    return arr
}

func getNonemptyCells(colOrRow string, board [][]string, index int) []string {
	restAgg := []string{}
	switch colOrRow {
	case "col":
		for i := 0; i < getNumRows(board); i++ {
			if board[i][index] == "_" {
				continue
			}
			restAgg = append(restAgg, board[i][index])
		}
	case "row":
		for i := 0; i < getNumCols(board); i++ {
			if board[index][i] == "_" {
				continue
			}
			restAgg = append(restAgg, board[index][i])
		}
	}
	return restAgg
}

func handleUp(board [][]string) {
	for c := 0; c < getNumCols(board); c++ {
		restAgg := getNonemptyCells("col", board, c)
		// do intermediate merging (e.g 5+5 -> 10)
		restAgg = mergePairDuplicates(restAgg, "start")

		// fill it back
		for i := 0; i < len(restAgg); i++ {
			board[i][c] = restAgg[i]
		}
		for i := len(restAgg); i < getNumRows(board); i++ {
			board[i][c] = "_"
		}
	}
}

func handleDown(board [][]string) {
	for c := 0; c < getNumCols(board); c++ {
		restAgg := getNonemptyCells("col", board, c)
		// do intermediate merging (e.g 5+5 -> 10)
		restAgg = mergePairDuplicates(restAgg, "end")

		// fill it back
		for i := 0; i < len(restAgg); i++ {
			board[getNumRows(board)-len(restAgg)+i][c] = restAgg[i]
		}
		for i := 0; i < getNumRows(board)-len(restAgg); i++ {
			board[i][c] = "_"
		}
	}
}

func handleLeft(board [][]string) {
	for r := 0; r < getNumRows(board); r++ {
		restAgg := getNonemptyCells("row", board, r)
		// do intermediate merging (e.g 5+5 -> 10)
		restAgg = mergePairDuplicates(restAgg, "start")

		// fill it back
		for i := 0; i < len(restAgg); i++ {
			board[r][i] = restAgg[i]
		}
		for i := len(restAgg); i < getNumCols(board); i++ {
			board[r][i] = "_"
		}
	}
}

func handleRight(board [][]string) {
	for r := 0; r < getNumRows(board); r++ {
		restAgg := getNonemptyCells("row", board, r)
		// do intermediate merging (e.g 5+5 -> 10)
		restAgg = mergePairDuplicates(restAgg, "end")

		// fill it back
		for i := 0; i < len(restAgg); i++ {
			board[r][getNumCols(board)-len(restAgg)+i] = restAgg[i]
		}
		for i := 0; i < getNumCols(board)-len(restAgg); i++ {
			board[r][i] = "_"
		}
	}
}

func handleMove(move string, board [][]string) {
	switch move {
	case "up":
		handleUp(board)
	case "down":
		handleDown(board)
	case "left":
		handleLeft(board)
	case "right":
		handleRight(board)
	}
}

func isValidMove(move string) bool {
	return move == "up" || move == "down" || move == "right" || move == "left"
}

func main() {
	board := [][]string{
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
	}
	for {
		insertRandomFive(board)
		printBoard(board)
		if isGameLost(board) {
			fmt.Println("You lost!")
			break
		}
		fmt.Println()
		fmt.Print("> ")
		var move string
		n, err := fmt.Scanf("%s", &move)
		if n == 0 {
			break
		}
		check(err)
		if !isValidMove(move) {
			fmt.Println("Invalid move! Move must be 'up', 'right', 'down', 'left'")
			continue
		}
		handleMove(move, board)
	}
}
