package ds

import (
	"fmt"
	"sync"
)

type Queue interface {
	// Get number of elements in the queue.
	Size() int64

	// Whether the queue is empty or not.
	Empty() bool

	// Clear the queue.
	Clear()

	// Get the first element of queue, and remove it from deque.
	Get() (el interface{}, err error)

	// Insert element at the back of the queue.
	Put(el interface{})
}

// Create a queue.
func NewQueue() (q Queue) {
	return &SimpleQueue{
		length:   0,
		elements: make([]interface{}, 0),
	}
}

// A simple queue implementation.
type SimpleQueue struct {
	sync.RWMutex
	length   int64         // length
	elements []interface{} // elements, stores the data.
}

// Get number of elements in the queue.
func (q *SimpleQueue) Size() int64 {
	return q.length
}

// Whether the queue is empty or not.
func (q *SimpleQueue) Empty() bool {
	return q.length == 0
}

// Clear the queue.
func (q *SimpleQueue) Clear() {
	q.Lock()
	defer q.Unlock()
	q.elements = nil
	q.length = 0
}

// Get the first element of queue, and remove it from deque.
func (q *SimpleQueue) Get() (el interface{}, err error) {
	q.Lock()
	defer q.Unlock()
	if q.length > 0 {
		el, q.elements = q.elements[0], q.elements[1:]
		q.length--
		return el, nil
	} else {
		return nil, fmt.Errorf("empty queue")
	}
}

// Insert element at the back of the queue.
func (q *SimpleQueue) Put(el interface{}) {
	q.Lock()
	defer q.Unlock()
	q.elements = append(q.elements, el)
	q.length++
}
