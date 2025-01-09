package educative0016

import (
	"fmt"
	"strconv"

	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

type stack[T constraints.Integer] struct {
	elements []*internal.TreeNode[T]
	size     int
}

func (s *stack[T]) pop() (*internal.TreeNode[T], bool) {
	if s.size == 0 {
		return nil, false
	}

	e := s.elements[s.size-1]
	s.elements = s.elements[:s.size-1]
	s.size--
	return e, true
}

func (s *stack[T]) push(n *internal.TreeNode[T]) {
	s.elements = append(s.elements, n)
	s.size++
}

func (s *stack[T]) isEmpty() bool {
	return s.size == 0
}

func newStack[T constraints.Integer]() *stack[T] {
	s := &stack[T]{}
	s.elements = make([]*internal.TreeNode[T], 0)
	return s
}

func serialize[T constraints.Integer](root *internal.TreeNode[T]) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}
	s := newStack[T]()
	s.push(root)
	for !s.isEmpty() {
		cur, _ := s.pop()
		if cur != nil {
			res = append(res, fmt.Sprintf("%v", cur.Val))
			s.push(cur.Right)
			s.push(cur.Left)
		} else {
			res = append(res, "")
		}
	}

	return res
}

func deserialize[T constraints.Integer](
	stream *[]string,
) *internal.TreeNode[T] {
	if stream == nil || len(*stream) == 0 {
		return nil
	}

	val, err := strconv.Atoi((*stream)[0])
	if err != nil {
		return nil
	}

	root := &internal.TreeNode[T]{Val: T(val)}
	s := newStack[T]()
	s.push(root)
	cur, ok := s.pop()
	if !ok {
		return root
	}
	i := 1
	for i < len(*stream) {
		if len((*stream)[i]) > 0 {
			val, err := strconv.Atoi((*stream)[i])
			if err != nil {
				return nil
			}
			node := &internal.TreeNode[T]{Val: T(val)}
			if cur.Left == nil {
				cur.Left = node
				s.push(cur)
				cur = node
			} else {
				cur.Right = node
				cur = node
			}
		} else {
			if cur.Left == nil {
				i++
			}
			if len((*stream)[i]) > 0 {
				val, err := strconv.Atoi((*stream)[i])
				if err != nil {
					return nil
				}
				node := &internal.TreeNode[T]{Val: T(val)}
				cur.Right = node
				s.push(node)
			}
			cur, ok = s.pop()
			if !ok {
				return root
			}
		}
		i++
	}

	return root
}
