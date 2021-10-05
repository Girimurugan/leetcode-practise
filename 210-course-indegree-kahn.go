package main

import (
	"fmt"
)

type queue struct {
	items []int
}

//enqueue

func (q *queue) enqueue(item int) {
	q.items = append(q.items, item)
}

//dequeue
func (q *queue) dequeue() int {
	toDequeue := q.items[0]
	q.items = q.items[1:]

	return toDequeue
}

func adjancencyList(edges [][]int) (map[int][]int, map[int]int) {

	adjList := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, edge := range edges {

		fmt.Println(edge)

		if _, exists := adjList[edge[0]]; !exists {
			adjList[edge[0]] = []int{}
		}
		if _, exists := adjList[edge[1]]; !exists {
			adjList[edge[1]] = []int{}
		}
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])

		if _, exists := inDegree[edge[1]]; !exists {
			inDegree[edge[1]] = 0
		}

		if _, exists := inDegree[edge[0]]; !exists {
			inDegree[edge[0]] = 0
		}
		inDegree[edge[1]]++
	}

	return adjList, inDegree
}

func topoligicalSortKahn(prerequisites [][]int) []int {

	/*
	   1. Convert prerequistes as Adj List
	   2. Calculate the indegree in another map
	   3. add all zeros indegree nodes to queue, add into output array
	   4. Dequeue till queue is emprty
	   5. for each dequeue node, find the unvisited neighbours and reduce the degree
	   6. If there are zero degree nodes then add it to queue

	*/

	// Find the adj list
	adjList, inDegree := adjancencyList(prerequisites)
	fmt.Println(adjList)
	fmt.Println(inDegree)


	q := &queue{}
	coursesInOrder := []int{}

	for key,value := range inDegree{
		if value == 0{
			q.enqueue(key)
		}
	}

	for len(q.items) > 0 {
		currentNode := q.dequeue()

		
		coursesInOrder = append(coursesInOrder, currentNode)

		for _, neighbour := range adjList[currentNode]{

			inDegree[neighbour]--
			if inDegree[neighbour] == 0{
				q.enqueue(neighbour)
			}
			
		}
	}

	return coursesInOrder
}

func main() {

	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	fmt.Println(topoligicalSortKahn(prerequisites))

}
