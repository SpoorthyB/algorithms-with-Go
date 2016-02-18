package main

func CountIslands(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		      return 0
		    }
		    numRow := len(grid)
		    numCol := len(grid[0])

		    visited := [][]bool{}
		    one := []bool{false,false,false,false,false}
		    two	:= []bool{false,false,false,false,false}
		    three:=[]bool{false,false,false,false,false}
		    four:= []bool{false,false,false,false,false}
		    visited = append(visited,one)
		    visited = append(visited,two)
		    visited = append(visited,three)
		    visited = append(visited,four)
		    numIsland := 0
		    for i := 0; i < numRow; i++ {
		      for j := 0; j < numCol; j++ {
		        if grid[i][j] == 1 && !visited[i][j] {
		          numIsland++;
		          search(grid, visited, i, j);
		        }
		      }
		    }

		    return numIsland;
}

func search(grid [][] int, visited [][] bool,row int,col int) {
 	xDir := []int{1, -1, 0, 0}
 	yDir := []int{0, 0, 1, -1}
    if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
      return;
    } else if visited[row][col] || grid[row][col] == 0 {
      return;
    }

    visited[row][col] = true;
    for i:=0; i < 4; i++ {
      search(grid, visited, row + xDir[i], col + yDir[i]);
    }
}
