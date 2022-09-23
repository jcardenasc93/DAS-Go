package challenges

import "testing"

var mazeMap = []string{
	"xxxxxxxxxx x",
	"x        x x",
	"x        x x",
	"x xxxxxxxx x",
	"x          x",
	"x xxxxxxxxxx",
}

var successPath = []Point{

	{x: 10, y: 0},
	{x: 10, y: 1},
	{x: 10, y: 2},
	{x: 10, y: 3},
	{x: 10, y: 4},
	{x: 9, y: 4},
	{x: 8, y: 4},
	{x: 7, y: 4},
	{x: 6, y: 4},
	{x: 5, y: 4},
	{x: 4, y: 4},
	{x: 3, y: 4},
	{x: 2, y: 4},
	{x: 1, y: 4},
	{x: 1, y: 5},
}

func drawPath(maze []string, path []Point) []string {
	var finalMaze = maze
	for _, point := range path {
		finalMaze[point.y] = finalMaze[point.y][:point.x] + "o" + finalMaze[point.y][point.x+1:]
	}
	return finalMaze
}

func resultIsOk(result *[]Point, drawResult []string, drawSuccess []string) bool {
	if len(*result) != len(successPath) {
		return false
	}
	for i := 0; i < len(drawResult); i++ {
		if drawResult[i] != drawSuccess[i] {
			return false
		}
	}
	return true
}

func TestMazeSolver(t *testing.T) {
	result := solveMaze(mazeMap, "x", Point{x: 10, y: 0}, Point{x: 1, y: 5})
	if !(resultIsOk(&result, drawPath(mazeMap, result), drawPath(mazeMap, successPath))) {
		t.Errorf("Maze Solver is wrong\nExpected:\n%v\n--------------------Got:\n%v", successPath, result)
	}
}
