package core

import (
	"fmt"
	"sync"
)

const (
	WATER = 0
	LAND  = 1
)

type IslandCounterSettings struct {
	Parallel bool
}

type IslandCounterOptions struct {
	BreakOnDiagonal bool
}

func IslandCounter(topography [][]int, options IslandCounterOptions, settings IslandCounterSettings) (int, []struct {
	Column int
	Row    int
}, error) {

	islandCounter := 0

	// assert rectangle
	rows := len(topography)
	columns := 0
	for i, row := range topography {
		if i == 0 {
			columns = len(row)
			continue
		}
		if len(row) != columns {
			return 0, nil, fmt.Errorf("topography not restangular, got %d rows, and row %d had %d columns and expected %d based on previous rows", rows, i, len(row), columns)
		}
	}

	visitedMap := NewLockableMatrix(columns, rows)

	for rowNumber, row := range topography {
		for columnNumber, surfaceTexture := range row {

			// if we have been here then skip
			if visitedMap.HasVisited(columnNumber, rowNumber) {
				continue
			}

			if surfaceTexture == WATER {
				visitedMap.Visits(columnNumber, rowNumber)
				continue
			}

			// if land then we recursively check all neighbors
			if settings.Parallel {
				VisitCellAndAllConnectedNeighborsParallel(columnNumber, rowNumber, rows, columns, options.BreakOnDiagonal, topography, &visitedMap)
			} else {
				VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber, rows, columns, options.BreakOnDiagonal, topography, &visitedMap)
			}

			islandCounter++
		}
	}

	return islandCounter, visitedMap.visitedList, nil
}

func VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber, numberOfRows, numberOfColumns int, breakOnDiagonal bool, topography [][]int, visitedMap *LockableMatrix) {

	// mark it as visited
	visitedMap.Visits(columnNumber, rowNumber)

	// if water then skip
	if topography[rowNumber][columnNumber] == WATER {
		return
	}

	maxColumnNumber := numberOfColumns - 1
	maxRowNumber := numberOfRows - 1

	// check all adjoining squares

	// column to the left
	if (columnNumber - 1) >= 0 {

		// upper left
		if !breakOnDiagonal && rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber-1, rowNumber-1) {
			VisitCellAndAllConnectedNeighbors(columnNumber-1, rowNumber-1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}

		// center left
		if !visitedMap.HasVisited(columnNumber-1, rowNumber) {
			VisitCellAndAllConnectedNeighbors(columnNumber-1, rowNumber, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}

		// bottom left
		if !breakOnDiagonal && rowNumber+1 <= maxRowNumber && !visitedMap.HasVisited(columnNumber-1, rowNumber+1) {
			VisitCellAndAllConnectedNeighbors(columnNumber-1, rowNumber+1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}
	}

	// same column

	// above
	if rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber, rowNumber-1) {
		VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber-1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
	}

	// below
	if rowNumber+1 <= maxRowNumber && !visitedMap.HasVisited(columnNumber, rowNumber+1) {
		VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber+1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
	}

	// column to the right
	if (columnNumber + 1) <= maxColumnNumber {

		// upper left
		if !breakOnDiagonal && rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber+1, rowNumber-1) {
			VisitCellAndAllConnectedNeighbors(columnNumber+1, rowNumber-1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}

		// center left
		if !visitedMap.HasVisited(columnNumber+1, rowNumber) {
			VisitCellAndAllConnectedNeighbors(columnNumber+1, rowNumber, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}

		// bottom left
		if !breakOnDiagonal && rowNumber+1 <= maxRowNumber && !visitedMap.HasVisited(columnNumber+1, rowNumber+1) {
			VisitCellAndAllConnectedNeighbors(columnNumber+1, rowNumber+1, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
		}
	}
}

func VisitCellAndAllConnectedNeighborsParallel(columnNumber, rowNumber, numberOfRows, numberOfColumns int, breakOnDiagonal bool, topography [][]int, visitedMap *LockableMatrix) {

	// mark it as visited
	visitedMap.VisitsSafe(columnNumber, rowNumber)

	// if water then skip
	if topography[rowNumber][columnNumber] == WATER {
		return
	}

	maxColumnNumber := numberOfColumns - 1
	maxRowNumber := numberOfRows - 1

	cellsToVisit := make(chan struct {
		column int
		row    int
	})
	done := make(chan bool)

	var wg sync.WaitGroup

	go func() {
		for {
			select {

			case cell := <-cellsToVisit:
				if visitedMap.HasVisitedSafe(cell.column, cell.row) {
					wg.Done()
					continue
				}
				go func() {
					VisitCellAndAllConnectedNeighborsParallel(cell.column, cell.row, numberOfRows, numberOfColumns, breakOnDiagonal, topography, visitedMap)
					wg.Done()
				}()

			case <-done:
				return
			}
		}
	}()

	// check all adjoining squares

	// column to the left
	if (columnNumber - 1) >= 0 {

		// upper left
		if !breakOnDiagonal && rowNumber-1 >= 0 {
			wg.Add(1)
			cellsToVisit <- struct {
				column int
				row    int
			}{columnNumber - 1, rowNumber - 1}
		}

		// center left
		wg.Add(1)
		cellsToVisit <- struct {
			column int
			row    int
		}{columnNumber - 1, rowNumber}

		// bottom left
		if !breakOnDiagonal && rowNumber+1 <= maxRowNumber {
			wg.Add(1)
			cellsToVisit <- struct {
				column int
				row    int
			}{columnNumber - 1, rowNumber + 1}
		}
	}

	// same column

	// above
	if rowNumber-1 >= 0 {
		wg.Add(1)
		cellsToVisit <- struct {
			column int
			row    int
		}{columnNumber, rowNumber - 1}
	}

	// below
	if rowNumber+1 <= maxRowNumber {
		wg.Add(1)
		cellsToVisit <- struct {
			column int
			row    int
		}{columnNumber, rowNumber + 1}
	}

	// column to the right
	if (columnNumber + 1) <= maxColumnNumber {

		// upper right
		if !breakOnDiagonal && rowNumber-1 >= 0 {
			wg.Add(1)
			cellsToVisit <- struct {
				column int
				row    int
			}{columnNumber + 1, rowNumber - 1}
		}

		// center right
		wg.Add(1)
		cellsToVisit <- struct {
			column int
			row    int
		}{columnNumber + 1, rowNumber}

		// bottom right
		if !breakOnDiagonal && rowNumber+1 <= maxRowNumber {
			wg.Add(1)
			cellsToVisit <- struct {
				column int
				row    int
			}{columnNumber + 1, rowNumber + 1}
		}
	}

	wg.Wait()
	done <- true

}
