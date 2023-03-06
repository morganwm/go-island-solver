package traversals

import (
	"github.com/morganwm/go-island-solver/constants"
	"github.com/morganwm/go-island-solver/utils"
)

type IslandTraverser func(
	columnNumber, rowNumber, numberOfRows, numberOfColumns int,
	breakOnDiagonal bool,
	topography [][]int,
	visitedMap *utils.LockableMatrix,
)

var Traversers = map[string]IslandTraverser{
	constants.PARALLEL:    VisitCellAndAllConnectedNeighborsParallel,
	constants.SERIES_IFS:  VisitCellAndAllConnectedNeighborsIfs,
	constants.SERIES_LOOP: VisitCellAndAllConnectedNeighborsLoop,
}

func GetAllowedTraversers() (out []string) {
	for k := range Traversers {
		out = append(out, k)
	}
	return
}
