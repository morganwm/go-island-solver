package parallel

import (
	"github.com/morganwm/go-island-solver/typedefs"
	"github.com/morganwm/go-island-solver/utils"
)

var Traversers utils.ComposableMap[typedefs.IslandTraverser]

func init() {
	Traversers.SetDefault(VisitCellAndAllConnectedNeighborsParallel)
}
