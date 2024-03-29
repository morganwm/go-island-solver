package traversals

import (
	"github.com/morganwm/go-island-solver/core/traversals/dfs"
	"github.com/morganwm/go-island-solver/typedefs"
	"github.com/morganwm/go-island-solver/utils"
)

var Traversers utils.ComposableMap[typedefs.IslandTraverser]

func init() {
	Traversers.SetDefault(dfs.Traversers.GetDefault())
	Traversers.Add("dfs", dfs.Traversers)
}
