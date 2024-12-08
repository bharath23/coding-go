package internal

type Queue struct {
	values []int
	size   int
}

func NewQueue() *Queue {
	q := &Queue{}
	q.values = make([]int, 0)
	return q
}

func (q *Queue) IsEmpty() bool {
	return (q.size == 0)
}

func (q *Queue) Enqueue(val int) {
	q.values = append(q.values, val)
	q.size++
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	q.size--
	value := q.values[0]
	q.values = q.values[1:]
	return value, true
}
