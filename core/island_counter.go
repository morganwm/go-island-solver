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

	traverser, ok := traversals.Traversers[settings.Mode]
	if !ok {
		return 0, nil,
			fmt.Errorf(
				"unknown mode selected: %s, available modes: %v",
				settings.Mode, traversals.GetAllowedTraversers(),
			)
	}

	// assert rectangle
	rows := len(topography)
	columns := 0
	for i, row := range topography {
		if i == 0 {
			columns = len(row)
			continue
		}
		if len(row) != columns {
			return 0, nil,
				fmt.Errorf(
					"topography not rectangular, got %d rows, and row %d had %d columns and expected %d based on previous rows",
					rows, i, len(row), columns,
				)
		}
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
