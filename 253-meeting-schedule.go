package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type meetingIntervals struct{
	startTime int
	endTime int
}


type PriorityQueue []*meetingIntervals

func (pq PriorityQueue)Less(i,j int)bool{
	return pq[i].startTime > pq[j].startTime
}

func (pq PriorityQueue)Len()int{
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int){
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue)Push(x interface{}){
	item := x.(*meetingIntervals)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue)Pop() interface{}{
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func minMeetingRooms(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}


	sort.Slice(intervals, func(i, j int)bool{
		return intervals[i][0] < intervals[j][0]
	})

	pq := PriorityQueue{}

	interval := meetingIntervals{startTime: intervals[0][0], endTime: intervals[0][1]}
	pq = append(pq,&interval)
	
	heap.Init(&pq)

	
	for i := 1; i < len(intervals); i++{
		poppedInterval := heap.Pop(&pq).(*meetingIntervals)
		if intervals[i][0] >= poppedInterval.endTime{
			toPush := meetingIntervals{startTime: intervals[i][0], endTime: intervals[i][1]}
			heap.Push(&pq, toPush)
		}else{
			heap.Push(&pq, poppedInterval)
		}
	}

    
	return len(pq)
}

func main(){
	fmt.Println(minMeetingRooms([][]int{{0,30}, {15,20},{5,10}}))
}