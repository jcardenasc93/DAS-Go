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
	yLimit := len(*maze)
	xLimit := len((*maze)[current.y])
	if current.x < 0 || current.x >= xLimit || current.y < 0 || current.y >= yLimit {
		return true
	}
	return false
}

func checkSeen(current *Point, seen *[][]bool) bool {
	return (*seen)[current.y][current.x]
}

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

// move is the recursive func and evaluates the wrong cases or if exit is found
func move(maze *[]string, current Point, end Point, wall string, seen *[][]bool, path *[]Point) bool {
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
	visited := checkSeen(&current, seen)
	if visited {
		return false
	}
	// Check if exit found
	if current.y == end.y && current.x == end.x {
		*path = append(*path, current)
		return true
	}

	// Recursion here:
	// Pre op -> add current to path & seen
	(*seen)[current.y][current.x] = true

	*path = append(*path, current)
	// Recursion for each direction
	for i := 0; i < len(directions); i++ {
		x, y := directions[i][0], directions[i][1]
		nextCurrent := Point{x: current.x + x, y: current.y + y}
		if move(maze, nextCurrent, end, wall, seen, path) {
			// Stop recursion because end found
			return true
		}
	}
	// Post op -> remove current from path. When all of the dirs are false
	// the position must be poped from path slice
	*path = (*path)[:len(*path)-1]

	return false
}

func solveMaze(maze []string, wall string, start Point, end Point) []Point {
	var seen = make([][]bool, len(maze))
	var path = []Point{}

	// Set seen as false
	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[i]))
	}

	move(&maze, start, end, wall, &seen, &path)
	return path
}
