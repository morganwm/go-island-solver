package solver

// IslandCounter counts the number of islands in the given topography.
// It returns the count, the list of visited cells in order, and any error encountered.
func IslandCounter(topography [][]int, breakOnDiagonal bool, parallel bool) (int, []struct {
	Column int
	Row    int
}, error) {

	rows, columns, err := RectangleDimensions(topography)
	if err != nil {
		return 0, nil, err
	}

	islandCounter := 0
	visitedMap := NewLockableMatrix(columns, rows)

	for rowNumber, row := range topography {
		for columnNumber, surfaceTexture := range row {

			// if we have been here then skip
			if visitedMap.HasVisited(columnNumber, rowNumber) {
				continue
			}

			// if water (0) then mark visited and skip
			if surfaceTexture == 0 {
				visitedMap.Visits(columnNumber, rowNumber)
				continue
			}

			// if land then we recursively check all neighbors
			if parallel {
				VisitCellAndAllConnectedNeighborsParallel(
					columnNumber,
					rowNumber,
					rows,
					columns,
					breakOnDiagonal,
					topography,
					&visitedMap,
				)
			} else {
				VisitCellAndAllConnectedNeighborsLoop(
					columnNumber,
					rowNumber,
					rows,
					columns,
					breakOnDiagonal,
					topography,
					&visitedMap,
				)
			}

			islandCounter++
		}
	}

	return islandCounter, visitedMap.VisitedList, nil
}
