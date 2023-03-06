package dfs

import (
	"github.com/morganwm/go-island-solver/core/traversals/dfs/parallel"
	"github.com/morganwm/go-island-solver/core/traversals/dfs/series"
	"github.com/morganwm/go-island-solver/typedefs"
	"github.com/morganwm/go-island-solver/utils"
)

var Traversers utils.ComposableMap[typedefs.IslandTraverser]

func init() {
	Traversers.SetDefault(series.Traversers.GetDefault())
	Traversers.Add("series", series.Traversers)
	Traversers.Add("parallel", parallel.Traversers)
}
