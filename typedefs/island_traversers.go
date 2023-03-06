package typedefs

import "github.com/morganwm/go-island-solver/utils"

type IslandTraverser func(
	columnNumber, rowNumber, numberOfRows, numberOfColumns int,
	breakOnDiagonal bool,
	topography [][]int,
	visitedMap *utils.LockableMatrix,
)
