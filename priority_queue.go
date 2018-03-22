package main

import (
	"container/heap"
	"fmt"
	"sync"
)

type priorityQueue struct {
	heap []*node
	c    *sync.Cond
}

// Len returns the len of the queue
func (pq priorityQueue) Len() int { return len(pq.heap) }

// Less returns if the element at index i is or not less than the
// element at index j
func (pq priorityQueue) Less(i, j int) bool {
	return pq.heap[i].rank < pq.heap[j].rank
}

// Swap swaps two element at specified indexes
func (pq priorityQueue) Swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
	pq.heap[i].index = i
	pq.heap[j].index = j
}

// Push pushes an ekement into the queue
func (pq *priorityQueue) Push(x interface{}) {
	l := len(pq.heap)
	n := x.(*node)
	n.index = l
	pq.heap = append(pq.heap, n)
}

// Pop pops an element
func (pq *priorityQueue) Pop() interface{} {
	old := pq.heap
	l := len(old)
	n := old[l-1]
	n.index = -1
	pq.heap = old[0 : l-1]

	return n
}

// heapPopChanMsg represents the data structure for a pop msg
type heapPopChanMsg struct {
	h      *priorityQueue
	result chan interface{}
}

// heapPushChanMsg represents the data structure for a push msg
type heapPushChanMsg struct {
	h    *priorityQueue
	x    interface{}
	done chan bool
}

// heapRemoveChanMsg represents  the data structure for a remove msg
type heapRemoveChanMsg struct {
	h     *priorityQueue
	index int
}

var (
	quitChan       chan bool
	heapPushChan   = make(chan heapPushChanMsg)
	heapPopChan    = make(chan heapPopChanMsg)
	heapRemoveChan = make(chan heapRemoveChanMsg)
)

// HeapPop safely pops an item from a heap interface
func HeapPop(h *priorityQueue) interface{} {
	var result = make(chan interface{}, 1)

	heapPopChan <- heapPopChanMsg{
		h:      h,
		result: result,
	}

	return <-result
}

// HeapPushSync safely pushes an item to an interface and takes a done channel
// to handle synchronicity
func HeapPushSync(h *priorityQueue, x interface{}, done chan bool) {
	heapPushChan <- heapPushChanMsg{
		h:    h,
		x:    x,
		done: done,
	}
}

// HeapPush safely pushes an item to a heap interface
func HeapPush(h *priorityQueue, x interface{}) {
	heapPushChan <- heapPushChanMsg{
		h: h,
		x: x,
	}
}

// HeapRemove safely removes an item from a heap interface
func HeapRemove(h *priorityQueue, index int) {
	if h.Len() != 0 {
		heapRemoveChan <- heapRemoveChanMsg{
			h:     h,
			index: index,
		}

		return
	}
}

//stopWatchHeapOps - stop watching for heap operations
func stopWatchHeapOps() {
	quitChan <- true
}

// watchHeapOps - watch for push/pops to our heap, and serializing the operations
// with channels
func watchHeapOps() chan bool {
	var quit = make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				// do something
				return
			case popMsg := <-heapPopChan:
				popMsg.result <- heap.Pop(popMsg.h)
			case pushMsg := <-heapPushChan:
				{
					fmt.Printf("BEGIN PUSH\n")
					heap.Push(pushMsg.h, pushMsg.x)
					fmt.Printf("FINISHED PUSH\n")
					if pushMsg.done != nil {
						close(pushMsg.done)
					}
				}
			case removeMsg := <-heapRemoveChan:
				heap.Remove(removeMsg.h, removeMsg.index)
			}
		}
	}()

	return quit
}
