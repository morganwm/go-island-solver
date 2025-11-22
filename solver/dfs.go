package solver

// VisitCellAndAllConnectedNeighborsLoop recursively visits all connected land cells using a loop.
func VisitCellAndAllConnectedNeighborsLoop(
	columnNumber, rowNumber, numberOfRows, numberOfColumns int,
	breakOnDiagonal bool,
	topography [][]int,
	visitedMap *LockableMatrix,
) {

	// mark it as visited
	visitedMap.Visits(columnNumber, rowNumber)

	// if water (0) then skip
	if topography[rowNumber][columnNumber] == 0 {
		return
	}

	maxColumnNumber := numberOfColumns - 1
	maxRowNumber := numberOfRows - 1

	// check all adjoining squares
	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		rowTarget := rowNumber + rowOffset

		// row is out of bounds
		if rowTarget > maxRowNumber || rowTarget < 0 {
			continue
		}

		for columnOffset := -1; columnOffset <= 1; columnOffset++ {
			columnTarget := columnNumber + columnOffset

			// column is out of bounds
			if columnTarget > maxColumnNumber || columnTarget < 0 {
				continue
			}

			// is the same spot
			if rowOffset == 0 && columnOffset == 0 {
				continue
			}

			// is diagonal
			if breakOnDiagonal && (rowOffset*columnOffset != 0) {
				continue
			}

			// has already been visited
			if visitedMap.HasVisited(columnTarget, rowTarget) {
				continue
			}

			VisitCellAndAllConnectedNeighborsLoop(
				columnTarget,
				rowTarget,
				numberOfRows,
				numberOfColumns,
				breakOnDiagonal,
				topography,
				visitedMap,
			)
		}
	}
}
