package main

import (
	"fmt"
)


type Queue struct{
	items []int
}

func (q *Queue)enqueue(rowCol int){
	q.items = append(q.items, rowCol)
}

func (q *Queue)dequeue()int{

	toDequeue := q.items[0]
	q.items = q.items[1:]
	return toDequeue
}

// this will be using the BFS algo with a queue
func numIslands(grid [][]byte)int{

	q := &Queue{}

	if grid == nil {
		return 0
	}

	rowCount := len(grid)
	colCount := len(grid[0])

	if rowCount == 0 || colCount == 0 {
		return 0
	}

	islandCount := 0

	for r:=0 ; r < rowCount ; r++ {

		for c:= 0; c<colCount; c++{

			if grid[r][c] == '1'{

				islandCount++

				rcId := (r*colCount)+c
				q.enqueue(rcId)

				grid[r][c] = '0'

				for len(q.items) > 0{

					current := q.dequeue()

					currR := current/colCount
					currC := current%colCount

					//fmt.Println(currR,"-",currC)

					if currR-1 >= 0 && grid[currR-1][currC] == '1' {
						q.enqueue((currR-1)*colCount+currC)
						grid[currR-1][currC] = '0'
					}

					if  currR+1 < rowCount && grid[currR+1][currC] == '1' {
						q.enqueue((currR+1)*colCount+currC)
						grid[currR+1][currC] = '0'
					}

					if currC-1 >= 0 && grid[currR][currC-1] == '1' {
						q.enqueue(currR*colCount+(currC-1))
						grid[currR][currC-1] = '0'
					}

					if currC+1 <  colCount && grid[currR][currC+1] == '1' {
						q.enqueue(currR*colCount+(currC+1))
						grid[currR][currC+1] = '0'
					}

				}


			}

		}

	}

	return islandCount

}



func main(){

	grid := [][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','0','1'}}
	
	fmt.Println(numIslands(grid))
	
	}