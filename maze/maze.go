package main

import (
	"fmt"
	"os"
)

func ReadMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, err = fmt.Fscanf(file, "%d %d", &row, &col)
	if err != nil {
		panic(err)
	}
	fmt.Println(row, col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, err = fmt.Fscanf(file, "%d", &maze[i][j])
			if err != nil {
				panic(err)
			}
		}
	}
	return maze
}

func main() {
	maze := ReadMaze("maze/maze.in")
	for _, row := range maze {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
