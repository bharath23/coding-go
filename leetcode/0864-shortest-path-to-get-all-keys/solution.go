package leetcode0864

type state struct {
	x, y  int
	keys  int
	steps int
}

type queue struct {
	elements []*state
	size     int
}

func (q *queue) dequeue() *state {
	if q.size == 0 {
		return nil
	}

	e := q.elements[0]
	q.elements = q.elements[1:]
	q.size--

	return e
}

func (q *queue) enqueue(s *state) {
	q.elements = append(q.elements, s)
	q.size++
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

func newQueue() *queue {
	q := &queue{
		elements: []*state{},
		size:     0,
	}

	return q
}

/*
The time complexity of BFS in this problem is determined by the total number of
unique states we may encounter. Each state is defined by the current grid
position and the set of collected keys. There are m × n possible positions and
up to 2^k combinations of collected keys (where k is the number of keys),
leading to a total of O(m × n × 2^k) possible states. Each state is visited at
most once, so the overall time complexity is O(m × n × 2^k).

The space complexity is also O(m × n × 2^k), as we need to store the visited
states and potentially queue all of them during BFS traversal.

Time complexity: O(m × n × 2^k)
Space complexity: O(m × n × 2^k)
*/

func shortestPathAllKeys(grid []string) int {
	m := len(grid)
	n := len(grid[0])
	var allKeys int
	var startX, startY int

	// dirs: up, down, left, right
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Find the starting point and number of permutations of the keys
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := grid[i][j]
			if c == '@' {
				startX, startY = i, j
			}
			if c >= 'a' && c <= 'f' {
				allKeys |= 1 << (c - 'a')
			}
		}
	}

	// track visited
	visited := make([][][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = make([]bool, 64) // max keys is 2^6 = 64
		}
	}

	q := newQueue()
	q.enqueue(&state{x: startX, y: startY, keys: 0, steps: 0})
	visited[startX][startY][0] = true
	for !q.isEmpty() {
		s := q.dequeue()
		for _, d := range dirs {
			x := s.x + d[0]
			y := s.y + d[1]
			keys := s.keys
			// make sure we are in the grid
			if (x < 0) || x >= m || y < 0 || y >= n {
				continue
			}

			c := grid[x][y]
			if c == '#' {
				continue // wall
			}

			if c >= 'A' && c <= 'F' && (keys&(1<<(c-'A')) == 0) {
				continue // lock without key
			}

			if c >= 'a' && c <= 'f' {
				keys |= 1 << (c - 'a')
			}

			if keys == allKeys {
				return s.steps + 1 // collected all keys
			}

			if !visited[x][y][keys] {
				visited[x][y][keys] = true
				q.enqueue(&state{x: x, y: y, keys: keys, steps: s.steps + 1})
			}
		}
	}

	return -1
}
