package leetcode0138

import (
	"github.com/bharath23/coding-go/internal"
)

/*
Two pass solution. First pass copy the list and create a mapping of old nodes
to new nodes. Second pass fill the random pointer. Time complexity is O(n) as
each node is visited twice. Space complexity is O(n) to store the map.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func copyRandomListV0(head *internal.RandomList) *internal.RandomList {
	var copyHead, copyTail *internal.RandomList
	nodeMap := make(map[*internal.RandomList]*internal.RandomList)
	for cur := head; cur != nil; cur = cur.Next {
		n := &internal.RandomList{Val: cur.Val}
		nodeMap[cur] = n
		if copyHead == nil {
			copyHead = n
		}

		if copyTail != nil {
			copyTail.Next = n
		}

		copyTail = n
	}

	copyCur := copyHead
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Random != nil {
			n := nodeMap[cur.Random]
			copyCur.Random = n
		}

		copyCur = copyCur.Next
	}

	return copyHead
}

/*
Recursive solution. We copy the nodes recursively, first copying next followed
by random. Time complexity is O(n) as all nodes are just visited once. Space
complexity is O(1) as no additional storage is required, can be considered
O(n) to account for recursion.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func copyRandomListV1(head *internal.RandomList) *internal.RandomList {
	nodeMap := make(map[*internal.RandomList]*internal.RandomList)
	var recursiveCopy func(*internal.RandomList) *internal.RandomList
	recursiveCopy = func(node *internal.RandomList) *internal.RandomList {
		if node == nil {
			return nil
		}

		if n := nodeMap[node]; n != nil {
			return n
		}

		n := &internal.RandomList{Val: node.Val}
		nodeMap[node] = n
		n.Next = recursiveCopy(node.Next)
		n.Random = recursiveCopy(node.Random)
		return n
	}

	return recursiveCopy(head)
}

/*
One pass solution. We copy the next and random nodes as we traverse the list.
We maintain a map of nodes so that if it has been visited we do not recreate
it. Time complexity O(n) because we visit all the nodes just once. Space
complexity is O(n) as we need to store the map.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/
func copyRandomListV2(head *internal.RandomList) *internal.RandomList {
	nodeMap := make(map[*internal.RandomList]*internal.RandomList)
	getNodeCopy := func(node *internal.RandomList) *internal.RandomList {
		if node == nil {
			return nil
		}

		if n := nodeMap[node]; n != nil {
			return n
		}

		n := &internal.RandomList{Val: node.Val}
		nodeMap[node] = n
		return n
	}

	if head == nil {
		return nil
	}

	newCur := getNodeCopy(head)
	for cur := head; cur != nil; cur = cur.Next {
		newCur.Next = getNodeCopy(cur.Next)
		newCur.Random = getNodeCopy(cur.Random)
		newCur = newCur.Next
	}

	return nodeMap[head]
}

/*
Traversal solution. We traverse the list creating a copy of the node and
inserting it right after it. We then traverse the list again fixing the Random
pointers for the newly copied node to point to the appropriate nodes. Then we
decouple the list. Time complexity is O(n) as we traverse the list thrice.
Space complexity is O(1) as no additional storage is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/
func copyRandomListV3(head *internal.RandomList) *internal.RandomList {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		n := &internal.RandomList{Val: cur.Val}
		n.Next = cur.Next
		cur.Next = n
		cur = cur.Next.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	var copyHead, copyTail *internal.RandomList
	cur = head
	for cur != nil {
		n := cur.Next
		cur.Next = n.Next
		if copyHead == nil {
			copyHead = n
		}

		if copyTail != nil {
			copyTail.Next = n
		}

		copyTail = n
		cur = cur.Next
	}

	return copyHead
}
