package main

import (
	"container/heap"
	"fmt"
)

type FrequentNum struct {
	num       int
	frequency int
}

type PriorityQueueNums []*FrequentNum

func (pq PriorityQueueNums) Len() int {
	return len(pq)
}

func (pq PriorityQueueNums) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency // start with the minimum frequency, so what is left will be the max frequency
}

func (pq PriorityQueueNums) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueNums) Push(x interface{}) {
	item := x.(*FrequentNum)
	*pq = append(*pq, item)
}

func (pq *PriorityQueueNums) Pop() interface{} {

	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return x

}

func topKFrequent(nums []int, k int) []int {

	// get the frequency counted in the map
	countMap := make(map[int]int)
	pq := &PriorityQueueNums{}

	heap.Init(pq)

	for _, num := range nums {

		if _, exists := countMap[num]; !exists {
			countMap[num] = 0
		}
		countMap[num]++
	}

	for key, value := range countMap {

		frequencyItem := &FrequentNum{num: key, frequency: value}
		heap.Push(pq, frequencyItem)

		if pq.Len() > k {
			heap.Pop(pq)
		}

	}
	result := make([]int,k)

	for i := k - 1; i >=0; i-- {
		result[i] = heap.Pop(pq).(*FrequentNum).num
	}

	return result
}

func main() {

	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
