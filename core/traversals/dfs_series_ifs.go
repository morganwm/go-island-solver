package traversals

import (
	"github.com/morganwm/go-island-solver/constants"
	"github.com/morganwm/go-island-solver/utils"
)

func VisitCellAndAllConnectedNeighborsIfs(
	columnNumber, rowNumber, numberOfRows, numberOfColumns int,
	breakOnDiagonal bool,
	topography [][]int,
	visitedMap *utils.LockableMatrix,
) {

	// mark it as visited
	visitedMap.Visits(columnNumber, rowNumber)

	// if water then skip
	if topography[rowNumber][columnNumber] == constants.WATER {
		return
	}

	maxColumnNumber := numberOfColumns - 1
	maxRowNumber := numberOfRows - 1

	// check all adjoining squares

	// column to the left
	if (columnNumber - 1) >= 0 {

		// upper left
		if !breakOnDiagonal && rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber-1, rowNumber-1) {
			VisitCellAndAllConnectedNeighborsIfs(columnNumber-1, rowNumber-1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}

		// center left
		if !visitedMap.HasVisited(columnNumber-1, rowNumber) {
			VisitCellAndAllConnectedNeighborsIfs(columnNumber-1, rowNumber, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}

		// bottom left
		if !breakOnDiagonal && rowNumber+1 <= maxRowNumber && !visitedMap.HasVisited(columnNumber-1, rowNumber+1) {
			VisitCellAndAllConnectedNeighborsIfs(columnNumber-1, rowNumber+1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}
	}

	// same column

	// above
	if rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber, rowNumber-1) {
		VisitCellAndAllConnectedNeighborsIfs(columnNumber, rowNumber-1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
	}

	// below
	if rowNumber+1 <= maxRowNumber && !visitedMap.HasVisited(columnNumber, rowNumber+1) {
		VisitCellAndAllConnectedNeighborsIfs(columnNumber, rowNumber+1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
	}

	// column to the right
	if (columnNumber + 1) <= maxColumnNumber {

		// upper left
		if !breakOnDiagonal && rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber+1, rowNumber-1) {
			VisitCellAndAllConnectedNeighborsIfs(columnNumber+1, rowNumber-1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}

		// center left
		if !visitedMap.HasVisited(columnNumber+1, rowNumber) {
			VisitCellAndAllConnectedNeighborsIfs(columnNumber+1, rowNumber, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}

		// bottom left
		if !breakOnDiagonal && rowNumber+1 <= maxRowNumber && !visitedMap.HasVisited(columnNumber+1, rowNumber+1) {
			VisitCellAndAllConnectedNeighborsIfs(columnNumber+1, rowNumber+1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}
	}
}
