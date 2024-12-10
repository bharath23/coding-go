package leetcode0430

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

/*
Preorder DFS traversal of the list. Time complexity is O(n) as we traverse
each nodes ones. The space complexity is O(1) as no additional storage is
required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func flattenV0[T constraints.Signed](
	root *internal.Node[T],
) *internal.Node[T] {
	var recursiveFlatten func(
		*internal.Node[T],
		*internal.Node[T],
	) *internal.Node[T]
	recursiveFlatten = func(prev, cur *internal.Node[T]) *internal.Node[T] {
		if cur == nil {
			return prev
		}

		cur.Prev = prev
		if prev != nil {
			prev.Next = cur
		}

		next := cur.Next
		tail := recursiveFlatten(cur, cur.Child)
		cur.Child = nil
		return recursiveFlatten(tail, next)
	}

	recursiveFlatten(nil, root)
	return root
}

/*
Preorder DFS traversal using stack. Time complexity is O(n), we traverse each
node ones. The space complexity is O(n) as we need additional storage for the
stack.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func flattenV1[T constraints.Signed](
	root *internal.Node[T],
) *internal.Node[T] {
	stack := internal.NewStack()
	if root == nil {
		return nil
	}

	var prev *internal.Node[T]
	stack.Push(root)
	newRoot := root
	for !stack.IsEmpty() {
		cur := stack.Pop().(*internal.Node[T])
		if cur.Next != nil {
			stack.Push(cur.Next)
			cur.Next = nil
		}

		if cur.Child != nil {
			stack.Push(cur.Child)
			cur.Child = nil
		}

		cur.Prev = prev
		if prev != nil {
			prev.Next = cur
		}

		prev = cur
	}

	return newRoot
}
