package algorithms

// Queue represents a circular buffer queue implementation
type Queue struct {
	items    []int
	capacity int
	front    int
	rear     int
	size     int
}

// NewQueue creates a new queue with the specified capacity
func NewQueue(capacity int) *Queue {
	return &Queue{
		items:    make([]int, capacity),
		capacity: capacity,
		front:    0,
		rear:     0,
		size:     0,
	}
}

// Push adds an item to the queue
func (q *Queue) Push(item int) {
	if q.size == q.capacity {
		// Queue is full, remove oldest item
		q.Pop()
	}
	q.items[q.rear] = item
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

// Pop removes and returns the oldest item from the queue
func (q *Queue) Pop() int {
	if q.size == 0 {
		return -1
	}
	item := q.items[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return item
}

// Peek returns the oldest item without removing it
func (q *Queue) Peek() int {
	if q.size == 0 {
		return -1
	}
	return q.items[q.front]
}

// ToSlice returns all items in the queue as a slice
func (q *Queue) ToSlice() []int {
	if q.size == 0 {
		return []int{}
	}

	result := make([]int, q.size)
	if q.front < q.rear {
		// Simple case: front to rear is sequential
		copy(result, q.items[q.front:q.rear])
	} else {
		// Wrapped around case
		firstPart := q.capacity - q.front
		copy(result, q.items[q.front:])
		copy(result[firstPart:], q.items[:q.rear])
	}
	return result
}
