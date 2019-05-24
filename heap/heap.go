package heap

// MinHeap represents a min heap
type MinHeap struct {
	data []int
}

// NewMinHeap creates and returns a new min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (m MinHeap) leftChild(index int) int {
	// if there isn't left child return 0 for now
	if !m.hasLeftChild(index) {
		return 0
	}
	return m.data[(index*2)+1]
}

func (m MinHeap) leftChildIndex(index int) int {
	// TODO: return error if there isn't such index
	return (index * 2) + 1
}

func (m MinHeap) rightChildIndex(index int) int {
	// TODO: return error if there isn't such index
	return (index * 2) + 2
}

func (m MinHeap) rightChild(index int) int {
	// if there isn't right child return 0 for now
	if !m.hasRightChild(index) {
		return 0
	}
	return m.data[(index*2)+2]
}

func (m MinHeap) hasParent(index int) bool {
	return m.parentIndex(index) >= 0
}

func (m MinHeap) hasLeftChild(index int) bool {
	if (2*index)+1 <= len(m.data)-1 {
		return true
	}
	return false
}

func (m *MinHeap) hasRightChild(index int) bool {
	if (index*2)+2 <= len(m.data)-1 {
		return true
	}
	return false
}

func (m MinHeap) hasChild(index int) bool {
	if m.hasLeftChild(index) || m.hasRightChild(index) {
		return true
	}
	return false
}

func (m MinHeap) parent(index int) int {
	return m.data[m.parentIndex(index)]
}

func (m MinHeap) parentIndex(index int) int {
	return (index - 1) / 2
}

// Push insert a new item to heap and ReHeapifies the heap
func (m *MinHeap) Push(item int) {
	m.data = append(m.data, item)
	m.HeapifyUp()
}

// Pop returns and removes the root aka minimum item in heap and ReHeapifies the heap
func (m *MinHeap) Pop() int {
	if len(m.data) == 0 {
		return 0
	}
	p := m.data[0]
	m.data[0] = m.data[len(m.data)-1]
	m.data = m.data[0 : len(m.data)-1]
	m.HeapifyDown()

	return p
}

// HeapifyUp will rebuild heap upward
func (m *MinHeap) HeapifyUp() {
	index := len(m.data) - 1
	parentIndex := m.parentIndex(index)

	child := m.data[index]
	parent := m.data[parentIndex]

	for m.hasParent(index) && parent > child {
		// swap
		m.data[index], m.data[parentIndex] = m.data[parentIndex], m.data[index]

		// move index forward
		index = parentIndex
		parentIndex = m.parentIndex(parentIndex)

		child = m.data[index]
		parent = m.data[parentIndex]
	}
}

// HeapifyDown will rebuild heap down-ward
func (m *MinHeap) HeapifyDown() {
	index := 0
	current := m.data[0]

	for m.hasLeftChild(index) {
		smallestIndex := m.leftChildIndex(index)
		if m.hasRightChild(index) {
			if m.rightChild(index) < m.leftChild(index) {
				smallestIndex = m.rightChildIndex(index)
			}
		}

		if current < m.data[smallestIndex] {
			break
		}

		m.data[index], m.data[smallestIndex] = m.data[smallestIndex], m.data[index]
		index = smallestIndex
		current = m.data[index]

	}
}
