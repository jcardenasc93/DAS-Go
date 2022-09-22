package challenges

/* Recursion

To implement recursion successfuly it's important to know 3 basic rules:
1 -> Must have a pre operation
2 -> After the pre operation the recurse should be call
3 -> Must have a post operation

Additionaly it's important to know in a very clear way the base
(most simple & doesn't depends from recursive return value) case for
the recursive func */

/* Maze Solver challenge

Given an map as arary of strings. Where:
"#" -> Is a wall
" " -> Is a clear path
"S" -> Is start point
"E" -> Is exit

Return an array of coordinates that indicates the path to the exit */

// Point: struct type to represent simple (x,y) coordinates
type Point struct {
	x int
	y int
}

// isOutOfMaze checks if current point is out of maze boundaries
func isOutOfMaze(maze *[]string, current *Point) bool {
	// Assumes maze is square
	limit := len(*maze)
	if current.x < 0 || current.x >= limit || current.y < 0 || current.y >= limit {
		return true
	}
	return false
}

func checkSeen(current *Point, seen *[][]bool) bool {
	if (*seen)[current.x][current.y] {
		return true
	} else {
		(*seen)[current.x][current.y] = true
		return false
	}
}

// move is the recursive func and evaluates the wrong cases or if exit is found
func move(maze *[]string, current Point, end string, wall string, seen [][]bool) bool {
	// Basic case: Evaluates if current point is
	// 1) Current is out of maze boundaries
	// 2) Current is a wall
	// 3) Current is end
	// 4) Current has been visited before

	outOfMaze := isOutOfMaze(maze, &current)
	if outOfMaze {
		return false
	}
	// Look for wall
	if string((*maze)[current.y][current.x]) == wall {
		return false
	}
	// Look for visited
	visited := checkSeen(&current, &seen)
	if visited {
		return false
	}
	return true
}
