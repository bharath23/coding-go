package meta0003

type queue struct {
	elements   []int
	noElements int
	size       int
}

func newQueue(k int) *queue {
	q := &queue{}
	q.elements = []int{}
	q.size = k
	q.noElements = 0
	return q
}

func (q *queue) dequeue() int {
	q.noElements--
	e := q.elements[0]
	q.elements = q.elements[1:]
	return e
}

func (q *queue) enqueue(v int) {
	q.elements = append(q.elements, v)
	q.noElements++
}

func (q *queue) isFull() bool {
	if q.noElements < q.size {
		return false
	}

	return true
}

func getMaximumDishEatenCount(N int, D []int, K int) int {
	q := newQueue(K)
	dishEaten := make(map[int]bool, N)
	count := 0
	for _, d := range D {
		if dishEaten[d] {
			continue
		}

		if q.isFull() {
			e := q.dequeue()
			dishEaten[e] = false
		}

		count++
		dishEaten[d] = true
		q.enqueue(d)
	}

	return count
}
