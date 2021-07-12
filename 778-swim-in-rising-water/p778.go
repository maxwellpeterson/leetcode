package main

import "container/heap"

func main() {}

func swimInWater(grid [][]int) int {
	pq := PriorityQueue{}
	heap.Init(&pq)

	addNeighbors(&Square{ii: 0, jj: 0, priority: grid[0][0]}, grid, &pq)
	grid[0][0] = -1

	for {
		next := heap.Pop(&pq).(*Square)

		if next.ii == len(grid)-1 && next.jj == len(grid)-1 {
			return next.priority
		} else {
			addNeighbors(next, grid, &pq)
		}
	}
}

func addNeighbors(sq *Square, grid [][]int, pq *PriorityQueue) {
	offsets := []int{-1, 1}

	for _, oo := range offsets {
		if sq.ii+oo >= 0 && sq.ii+oo < len(grid) && grid[sq.ii+oo][sq.jj] >= 0 {
			heap.Push(pq, &Square{ii: sq.ii + oo, jj: sq.jj, priority: max(sq.priority, grid[sq.ii+oo][sq.jj])})
			grid[sq.ii+oo][sq.jj] = -1
		}

		if sq.jj+oo >= 0 && sq.jj+oo < len(grid) && grid[sq.ii][sq.jj+oo] >= 0 {
			heap.Push(pq, &Square{ii: sq.ii, jj: sq.jj + oo, priority: max(sq.priority, grid[sq.ii][sq.jj+oo])})
			grid[sq.ii][sq.jj+oo] = -1
		}
	}
}

// Priority queue implementation taken from:
// https://golang.org/pkg/container/heap/

type Square struct {
	ii, jj   int // Position in grid
	priority int // Elevation
	index    int // Index of item in heap
}

type PriorityQueue []*Square

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Square)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func max(aa, bb int) int {
	if aa > bb {
		return aa
	}
	return bb
}
