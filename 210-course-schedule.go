package main

import (
	"fmt"
)

// using the topoligical sorting
/* Steps
1. Convert the Edges to Adjacency List
2. Range trhough the adjancency listed with Visited and Stack
3. Depth first search on the adjacen neighbours (recursive)
4. If reached the end (At depth) add the node to stack
5. Complete all and then print stack top to bottom

*/

type Stack struct{
	items []int
}

func (s *Stack) push(item int){
	s.items = append(s.items, item)
}

func (s * Stack)pop() int{

	toPop := 0
	if len(s.items) > 0 {
	toPop = s.items[len(s.items)-1]
	s.items = s.items[:(len(s.items)-1)]
	}
	return toPop
}

func adjancencyList(edges [][]int) map[int][]int{

	adjList := make(map[int][]int)
	
	for _, edge := range edges{

		fmt.Println(edge)

		if _, exists := adjList[edge[0]]; !exists{
			adjList[edge[0]] = []int{}
		}
		if _, exists := adjList[edge[1]]; !exists{
			adjList[edge[1]] = []int{}
		}
			adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		

	}

	fmt.Println(adjList)
	return adjList
}


func topoligicalSort(prerequisites [][]int)[]int{

	coursesInOrder := []int{}

	coursesInOrderStack := &Stack{}

	adjList := adjancencyList(prerequisites)

	visited := make(map[int]bool)

	for key, _ := range adjList{


		if visited[key] != true{
		topoligicalSortUtil(key, adjList, visited,coursesInOrderStack)
		}
	}
	

	for len(coursesInOrderStack.items) > 0 {
		coursesInOrder = append(coursesInOrder, coursesInOrderStack.pop())
	}


	return coursesInOrder
}

func topoligicalSortUtil(node int, adjList map[int][]int, visited map[int]bool, coursesInOrderStack *Stack){

	visited[node] = true

	for _, neighbour := range adjList[node]{

		if visited[neighbour] != true{
			topoligicalSortUtil(neighbour, adjList, visited,coursesInOrderStack)
		}
	}
	coursesInOrderStack.push(node)

}

// using Kahns's Algo



func main(){

	prerequisites := [][]int{{1,0},{2,0},{3,1},{3,2}}

	fmt.Println(topoligicalSort(prerequisites))

}