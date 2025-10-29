package main

//func main() {
//	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
//}

type MyQueue struct { // 单调递减队列
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

// 实现 Front Empty Push Back Pop

func (q *MyQueue) Front() int { // 输出当前窗口的最大值
	return q.queue[0]
}

func (q *MyQueue) Back() int { // 输出队列的最小值
	return q.queue[len(q.queue)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
}

func (q *MyQueue) Push(v int) {
	for len(q.queue) > 0 && v > q.queue[len(q.queue)-1] {
		// 截断
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, v)
}
func (q *MyQueue) Pop(v int) { // 只需移除最大值
	if !q.Empty() && v == q.queue[0] {
		q.queue = q.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	// 新建单调队列实例
	q := NewMyQueue()
	result := make([]int, 0)
	m := 0
	// 将k个放入单调队列中
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}

	// 目前最大值
	m = q.Front()
	result = append(result, m)
	// 移动
	for i := k; i < len(nums); i++ {
		// 先将前面的弹出（如果是最大值的话）
		q.Pop(nums[i-k])
		// 推进去当前的
		q.Push(nums[i])
		// 最大致
		result = append(result, q.Front())
	}
	return result
}
