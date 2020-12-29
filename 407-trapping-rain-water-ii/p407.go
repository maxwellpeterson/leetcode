package main

import (
	"container/heap"
)

func main() {}

// Running time:
// Memory usage:
func trapRainWater(heights [][]int) int {
	if len(heights) <= 2 || len(heights[0]) <= 2 {
		return 0
	}

	rows, cols := len(heights), len(heights[0])
	discovered, pQueue := initializeSearch(rows, cols, heights)

	// Main search phase:
	// - While priority queue is not empty, run modified BFS, noting height of root cell
	//   (Assume neighbor is not discovered, and mark discovered)
	//   - If neighbor taller than root, do not add to local queue, add to priority queue
	//   - If neighbor same height as root, treat normally
	//   - If neighbor shorter than root, treat normally, but add height difference to water counter

	totalWater := 0

	for pQueue.Len() > 0 {
		root := heap.Pop(&pQueue).(*Cell)
		neighbors := getNeighbors(root.row, root.col, rows, cols, heights, discovered)
		for _, neighbor := range neighbors {
			if neighbor.height < root.height {
				totalWater += root.height - neighbor.height
				neighbor.height = root.height
			}
			heap.Push(&pQueue, neighbor)
		}
	}

	return totalWater
}

func getNeighbors(row, col, rows, cols int, heights [][]int, discovered [][]bool) []*Cell {
	neighbors := []*Cell{}
	for _, offset := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		offRow, offCol := row+offset[0], col+offset[1]
		if offRow >= 0 && offRow < rows && offCol >= 0 && offCol < cols && !discovered[offRow][offCol] {
			neighbors = append(neighbors, &Cell{row: offRow, col: offCol, height: heights[offRow][offCol]})
		}
	}
	return neighbors
}

// Initialization:
// - Iterate over all nodes along border of input array
//   - Mark as discovered
//   - Add to priority queue based on height

func initializeSearch(rows, cols int, heights [][]int) ([][]bool, PriorityQueue) {
	discovered := make([][]bool, rows)
	pQueue := PriorityQueue{}
	heap.Init(&pQueue)

	// Initialize top and bottom rows
	for _, row := range []int{0, rows - 1} {
		discovered[row] = make([]bool, cols)
		for col := range discovered[row] {
			discovered[row][col] = true
			heap.Push(&pQueue, &Cell{row: row, col: col, height: heights[row][col]})
		}
	}

	// Initialize left and right columns, between top and bottom rows
	for row := 1; row < rows-1; row++ {
		discovered[row] = make([]bool, cols)
		for col := range []int{0, cols - 1} {
			discovered[row][col] = true
			heap.Push(&pQueue, &Cell{row: row, col: col, height: heights[row][col]})
		}
	}

	return discovered, pQueue
}

type Cell struct {
	row, col, height, index int
}

// Implements heap.Interface and holds Cells
type PriorityQueue []*Cell

func (pQueue PriorityQueue) Len() int {
	return len(pQueue)
}

func (pQueue PriorityQueue) Less(ii, jj int) bool {
	return pQueue[ii].height < pQueue[jj].height
}

func (pQueue PriorityQueue) Swap(ii, jj int) {
	pQueue[ii], pQueue[jj] = pQueue[jj], pQueue[ii]
	pQueue[ii].index = ii
	pQueue[jj].index = jj
}

func (pQueue *PriorityQueue) Push(x interface{}) {
	queueLength := len(*pQueue)
	cell := x.(*Cell)
	cell.index = queueLength
	*pQueue = append(*pQueue, cell)
}

func (pQueue *PriorityQueue) Pop() interface{} {
	oldQueue := *pQueue
	oldLength := len(oldQueue)
	cell := oldQueue[oldLength-1]
	oldQueue[oldLength-1] = nil
	cell.index = -1
	*pQueue = oldQueue[:oldLength-1]
	return cell
}
