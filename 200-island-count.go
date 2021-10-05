package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {

	if grid == nil{
		return 0
	}
	rowCount := len(grid)
	colCount := len(grid[0])

	if rowCount == 0 || colCount == 0{
		return 0
	}

	islandCount := 0

	for i:= 0 ; i < rowCount ; i++{
		for j:=0 ; j < colCount ; j++{

			if grid[i][j] == '1'{

				islandCount++
				dfs(grid,i,j,rowCount,colCount)

			}


		}
	}
    return islandCount
}

func dfs(grid [][]byte, r int, c int, rowCount int, colCount int){

	// base condition for recursion

	if r < 0 || c < 0 || r >= rowCount || c >= colCount || grid[r][c]== '0'{
		return
	}
	grid[r][c] = '0'
	dfs(grid,r-1,c,rowCount, colCount)
	dfs(grid,r,c-1,rowCount, colCount)
	dfs(grid,r+1,c,rowCount, colCount)
	dfs(grid,r,c+1,rowCount, colCount)


}


// Time Complexity : O(m x n)
// space complexity : O( m x n)

func main(){

grid := [][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','0','1'}}

fmt.Println(numIslands(grid))

}