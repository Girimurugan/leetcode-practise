package main

import (
	"fmt"
	"container/heap"
)

/* Sample Data analysis


 */

/*Logic for the brute force
sorting - which is O(NLogN)

*/

/*Optimal Logic

Using Quick Select which is better in best case
but in this example we will see using Heap
*/

type IntHeap []int

func (h IntHeap)Len()int{
	return len(h)
}

func (h IntHeap) Less (i, j int) bool{
	return h[i] < h[j] // min heap
}

func (h IntHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push (input interface{}){
	*h = append(*h, input.(int))
	fmt.Println("Pushed -", input.(int))
}

func (h *IntHeap) Pop()interface{}{

	// toPop := h.items[0]

	// //length := len(h.items) - 1
	
	// // move the last element first
	// //h.items[0] = h.items[length]

	// // shorten the slice
	// //h.items = h.items[:length]

	// h.items = h.items[1:]
	// fmt.Println("Poped -", toPop)
	// return toPop

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	fmt.Println("Poped -", x)
   return x

}

// Main Function
func findKthLargest(nums []int, k int) int {

	h := &IntHeap{}
	heap.Init(h)

	for i:=0; i< len(nums); i++{

		heap.Push(h, nums[i])

		if h.Len() > k{
			heap.Pop(h)
		}
	}



	return heap.Pop(h).(int)
    
}

// Base condition if is recursive

// return conditions

// Dry run with tests

// Time Complexity

// Space Complexity

// Driver function
func main(){

	fmt.Println(findKthLargest([]int{3,2,1,5,6,4},2))

	fmt.Println("-------")
	//fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6},4))

}