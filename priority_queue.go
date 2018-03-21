package main

type priorityQueue []*node

// Len returns the len of the queue
func (pq priorityQueue) Len() int { return len(pq) }

// Less returns if the element at index i is or not less than the
// element at index j
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].rank < pq[j].rank
}

// Swap swaps two element at specified indexes
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push pushes an ekement into the queue
func (pq *priorityQueue) Push(x interface{}) {
	l := len(*pq)
	n := x.(*node)
	n.index = l
	*pq = append(*pq, n)
}

// Pop pops an element
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	l := len(old)
	n := old[l-1]
	n.index = -1
	*pq = old[0 : l-1]

	return n
}
