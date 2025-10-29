package main

import (
	"container/heap"
)

//func main() {
//	fmt.Println(topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2))
//}

type MinHeap [][2]int // 构建小顶堆  key-value 需要实现 Len Less Swap Push Pop

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool { // 小顶堆的比较函数
	// 比较出现的频率
	return (*h)[i][1] < (*h)[j][1]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// 使用优先队列

	// 记录出现的次数
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}

	h := &MinHeap{}
	heap.Init(h)

	// 遍历map
	for key, value := range m {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			// 弹出
			heap.Pop(h)
		}
	}

	// 输出结果
	result := make([]int, 0, k)
	for h.Len() > 0 {
		item := heap.Pop(h).([2]int)
		result = append(result, item[0])
	}
	return result
}
