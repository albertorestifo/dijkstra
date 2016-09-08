package dijkstra

import "sort"

// Queue is a priority queue implementation
type Queue struct {
	keys  []string
	nodes map[string]int
}

// Len is part of sort.Interface
func (q *Queue) Len() int {
	return len(q.keys)
}

// Swap is part of sort.Interface
func (q *Queue) Swap(i, j int) {
	q.keys[i], q.keys[j] = q.keys[j], q.keys[i]
}

// Less is part of sort.Interface
func (q *Queue) Less(i, j int) bool {
	a := q.keys[i]
	b := q.keys[j]

	return q.nodes[a] < q.nodes[b]
}

// Set updates or inserts a new key in the priority queue
func (q *Queue) Set(key string, priority int) {
	// Initialize the map if not done so already

	// Inserts a new key if we don't have it already
	if _, ok := q.nodes[key]; !ok {
		q.keys = append(q.keys, key)
	}

	// Set the priority for the key
	q.nodes[key] = priority

	// Sort the keys array
	sort.Sort(q)
}

// Next removes the first element from the queue and retuns it's key and priority
func (q *Queue) Next() (key string, priority int) {
	// shift the key form the queue
	key, keys := q.keys[0], q.keys[1:]
	q.keys = keys

	priority = q.nodes[key]

	delete(q.nodes, key)

	return key, priority
}

// IsEmpty returns true when the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.keys) == 0
}

// Get returns the priority of a passed key
func (q *Queue) Get(key string) (priority int, ok bool) {
	priority, ok = q.nodes[key]
	return
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	var q Queue
	q.nodes = make(map[string]int)
	return &q
}
