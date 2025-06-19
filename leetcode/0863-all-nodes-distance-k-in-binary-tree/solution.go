package leetcode0863

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildParentMap(root *TreeNode, parentMap map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		parentMap[root.Left] = root
		buildParentMap(root.Left, parentMap)
	}

	if root.Right != nil {
		parentMap[root.Right] = root
		buildParentMap(root.Right, parentMap)
	}
}

/*
We traverse the tree twice: once to build the parent map, and once to perform
BFS or DFS to find nodes at distance K. Each traversal takes O(n) time, so the
overall time complexity is O(n). We also use additional space for the parent
map and the visited set, each requiring O(n) space, resulting in a total space
complexity of O(n).

Time complexity: O(n)
Space complexity: O
*/
func distanceK(root, target *TreeNode, k int) []int {
	result := []int{}
	if root == nil || target == nil {
		return result
	}

	parentMap := make(map[*TreeNode]*TreeNode)
	buildParentMap(root, parentMap)

	visited := make(map[*TreeNode]bool)
	cur := []*TreeNode{target}
	visited[target] = true
	distance := 0
	for len(cur) != 0 {
		if distance == k {
			break
		}
		next := []*TreeNode{}
		for _, n := range cur {
			for _, node := range []*TreeNode{n.Left, n.Right, parentMap[n]} {
				if node != nil && !visited[node] {
					visited[node] = true
					next = append(next, node)
				}
			}
		}

		cur = next
		distance++
	}

	for _, node := range cur {
		result = append(result, node.Val)
	}

	return result
}
