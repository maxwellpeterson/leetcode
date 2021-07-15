package main

func main() {}

type HeadTail struct {
	head, tail *Node
}

type Node struct {
	index, value int
	prev, next   *Node
}

func shortestSubarray(input []int, kk int) int {
	initial := &Node{-1, 0, nil, nil}
	monoDeque := HeadTail{initial, initial}

	acc := 0
	shortest := len(input) + 1

	for ii, vv := range input {
		acc += vv
		if vv > 0 {
			// Add new node to head of queue, guaranteed to preserve monotonic
			// ordering as accumulator increases at this iteration
			newHead := &Node{ii, acc, nil, monoDeque.head}
			monoDeque.head.prev = newHead
			monoDeque.head = newHead

			// Check if subarray that ends at current element sums to at least kk,
			// trim left if possible and compare to current minimum
			var nn *Node
			for nn = monoDeque.tail; nn.prev != nil && acc-nn.value >= kk; nn = nn.prev {
			}

			if nn != monoDeque.tail {
				shortest = min(shortest, ii-nn.next.index)
				nn.next = nil
				monoDeque.tail = nn
			}

		} else {
			// Add new node to head of queue, but preserve monotonic ordering by
			// popping elements greater than or equal to current accumulated value
			var nn *Node
			for nn = monoDeque.head; nn.next != nil && nn.next.value >= acc; nn = nn.next {
			}

			nn.index = ii
			nn.value = acc
			nn.prev = nil
			monoDeque.head = nn
		}

	}

	if shortest > len(input) {
		return -1
	} else {
		return shortest
	}
}

func min(aa, bb int) int {
	if aa < bb {
		return aa
	}
	return bb
}
