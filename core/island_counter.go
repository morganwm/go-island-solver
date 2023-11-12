package core

import (
	"fmt"

	"github.com/morganwm/go-island-solver/constants"
	"github.com/morganwm/go-island-solver/core/traversals"
	"github.com/morganwm/go-island-solver/utils"
)

type IslandCounterSettings struct {
	Mode string
}

type IslandCounterOptions struct {
	BreakOnDiagonal bool
}

func IslandCounter(topography [][]int, options IslandCounterOptions, settings IslandCounterSettings) (int, []struct {
	Column int
	Row    int
}, error) {

	traverser, ok := traversals.Traversers.GetExact(settings.Mode)
	if !ok {
		return 0, nil,
			fmt.Errorf(
				"unknown mode selected: %s, available modes: %v",
				settings.Mode, traversals.Traversers.GetKeys(),
			)
	}

	rows, columns, err := RectangleDimensions(topography)
	if err != nil {
		return 0, nil, err
	}

	islandCounter := 0
	visitedMap := utils.NewLockableMatrix(columns, rows)

	for rowNumber, row := range topography {
		for columnNumber, surfaceTexture := range row {

			// if we have been here then skip
			if visitedMap.HasVisited(columnNumber, rowNumber) {
				continue
			}

			if surfaceTexture == constants.WATER {
				visitedMap.Visits(columnNumber, rowNumber)
				continue
			}

			// if land then we recursively check all neighbors
			traverser(
				columnNumber,
				rowNumber,
				rows,
				columns,
				options.BreakOnDiagonal,
				topography,
				&visitedMap,
			)

			islandCounter++
		}
	}

	return islandCounter, visitedMap.VisitedList, nil
}
