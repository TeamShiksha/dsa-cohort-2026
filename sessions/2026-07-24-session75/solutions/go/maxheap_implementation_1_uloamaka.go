package main

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) heapifyUp() {
	curr := len(h.data) - 1

	for curr > 0 {
		parent := (curr - 1) / 2

		if h.data[parent] >= h.data[curr] {
			break
		}

		h.data[parent], h.data[curr] = h.data[curr], h.data[parent]

		curr = parent
	}
}

func (h *MaxHeap) heapifyDown() {
	curr := 0
	for curr < len(h.data)/2 {
		lci := (curr * 2) + 1
		rci := (curr * 2) + 2

		largest := curr

		if lci < len(h.data) && h.data[lci] > h.data[largest] {
			largest = lci
		}
		if rci < len(h.data) && h.data[rci] > h.data[largest] {
			largest = rci
		}
		if largest == curr {
			break
		}

		// swap largest and current
		h.data[curr], h.data[largest] = h.data[largest], h.data[curr]

		curr = largest
	}
}

func (h *MaxHeap) Push(x int) {
	h.data = append(h.data, x)
	h.heapifyUp()
}

func (h *MaxHeap) Pop() int {
	// save root
	root := h.data[0]

	// move last element to root
	h.data[0] = h.data[len(h.data)-1]

	// shrink slice
	h.data = h.data[:len(h.data)-1]

	h.heapifyDown()

	return root
}

func (h *MaxHeap) Peek() int {

	if len(h.data) == 0 {
		return -1
	}

	return h.data[0]
}

func (h *MaxHeap) Size() int {
	return len(h.data)
}
