package bfs

import "fmt"

func numEnclaves(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

    type Pair struct{ r, c int }
    queue := []Pair{}

    // Add all boundary land cells
    for c := 0; c < cols; c++ {
        if grid[0][c] == 1 {
            queue = append(queue, Pair{0, c})
        }
        if grid[rows-1][c] == 1 {
            queue = append(queue, Pair{rows-1, c})
        }
    }

    for r := 0; r < rows; r++ {
        if grid[r][0] == 1 {
            queue = append(queue, Pair{r, 0})
        }
        if grid[r][cols-1] == 1 {
            queue = append(queue, Pair{r, cols-1})
        }
    }

    // BFS directions
    dirs := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}

    // BFS: Mark all reachable land as water
    for len(queue) > 0 {
        cell := queue[0]
        queue = queue[1:]

        r, c := cell.r, cell.c
        if grid[r][c] == 0 {
            continue
        }

        grid[r][c] = 0 // mark as visited

        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]
            if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
                queue = append(queue, Pair{nr, nc})
            }
        }
    }

    // Count remaining land cells
    count := 0
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 1 {
                count++
            }
        }
    }

    return count
}

func main() {
	s := [][]int{
		{1,1,0,0},
		{1,1,1,0},
		{0,1,1,0},
		{0,0,0,0},
	}
	fmt.Println(numEnclaves(s))
}