package leetcode0215

import (
	"math/rand"
	"sort"
)

type MinHeap struct {
	data []int
	size int
}

func newMinHeap() *MinHeap {
	return &MinHeap{
		data: []int{},
		size: 0,
	}
}

func (h *MinHeap) siftdown() {
	n := h.size
	i := 0
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i
		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func (h *MinHeap) siftup() {
	i := h.size - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[i] >= h.data[parent] {
			break
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *MinHeap) Pop() int {
	val := h.data[0]
	h.size--
	h.data[0] = h.data[h.size]
	h.data = h.data[:h.size]
	h.siftdown()

	return val
}

func (h *MinHeap) Push(val int) {
	h.data = append(h.data, val)
	h.size++
	h.siftup()
}

func (h *MinHeap) Size() int {
	return h.size
}

func findKthLargestv0(nums []int, k int) int {
	sort.Sort(sort.IntSlice(nums))
	return nums[len(nums)-k]
}

func findKthLargestv1(nums []int, k int) int {
	h := newMinHeap()
	for _, n := range nums {
		h.Push(n)
		if h.Size() > k {
			h.Pop()
		}
	}

	return h.Pop()
}

func partition(nums []int, left, right int) int {
	pivotIndex := left + rand.Intn(right-left+1)
	pivot := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	i := left
	for j := left; j < right; j++ {
		if nums[j] > pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	pivotIndex := partition(nums, left, right)
	switch {
	case pivotIndex == k:
		return nums[k]
	case pivotIndex < k:
		return quickSelect(nums, pivotIndex+1, right, k)
	default:
		return quickSelect(nums, left, pivotIndex-1, k)
	}
}

func findKthLargestv2(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}
