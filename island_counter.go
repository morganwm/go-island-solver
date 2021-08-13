package main

import (
	"fmt"
	"sync"
)

func IslandCounter(topography [][]int, parallel bool) (int, []struct {
	Column int
	Row    int
}, error) {

	islandCounter := 0

	// assert square
	rows := len(topography)
	for i, row := range topography {
		if len(row) != rows {
			return 0, nil, fmt.Errorf("topography not square, got %d rows, and row %d had %d columns", rows, i, len(row))
		}
	}

	visitedMap := NewLockableMatrix(rows)

	for rowNumber, row := range topography {
		for columnNumber, surfaceTexture := range row {

			// if we have been here then skip
			if visitedMap.HasVisited(columnNumber, rowNumber) {
				continue
			}

			if surfaceTexture == 0 {
				visitedMap.Visits(columnNumber, rowNumber)
				continue
			}

			// if land then we recursively check all neighbors
			if parallel {
				VisitCellAndAllConnectedNeighborsParallel(columnNumber, rowNumber, topography, &visitedMap)
			} else {
				VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber, topography, &visitedMap)
			}

			islandCounter++
		}
	}

	return islandCounter, visitedMap.visitedList, nil
}

func VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber int, topography [][]int, visitedMap *LockableMatrix) {

	// mark it as visited
	visitedMap.Visits(columnNumber, rowNumber)

	// if water then skip
	if topography[rowNumber][columnNumber] == 0 {
		return
	}

	maxDist := len(topography) - 1

	// check all adjoining squares

	// column to the left
	if (columnNumber - 1) >= 0 {

		// upper left
		if rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber-1, rowNumber-1) {
			VisitCellAndAllConnectedNeighbors(columnNumber-1, rowNumber-1, topography, visitedMap)
		}

		// center left
		if !visitedMap.HasVisited(columnNumber-1, rowNumber) {
			VisitCellAndAllConnectedNeighbors(columnNumber-1, rowNumber, topography, visitedMap)
		}

		// bottom left
		if rowNumber+1 <= maxDist && !visitedMap.HasVisited(columnNumber-1, rowNumber+1) {
			VisitCellAndAllConnectedNeighbors(columnNumber-1, rowNumber+1, topography, visitedMap)
		}
	}

	// same column

	// above
	if rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber, rowNumber-1) {
		VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber-1, topography, visitedMap)
	}

	// below
	if rowNumber+1 <= maxDist && !visitedMap.HasVisited(columnNumber, rowNumber+1) {
		VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber+1, topography, visitedMap)
	}

	// column to the right
	if (columnNumber + 1) <= maxDist {

		// upper left
		if rowNumber-1 >= 0 && !visitedMap.HasVisited(columnNumber+1, rowNumber-1) {
			VisitCellAndAllConnectedNeighbors(columnNumber+1, rowNumber-1, topography, visitedMap)
		}

		// center left
		if !visitedMap.HasVisited(columnNumber+1, rowNumber) {
			VisitCellAndAllConnectedNeighbors(columnNumber+1, rowNumber, topography, visitedMap)
		}

		// bottom left
		if rowNumber+1 <= maxDist && !visitedMap.HasVisited(columnNumber+1, rowNumber+1) {
			VisitCellAndAllConnectedNeighbors(columnNumber+1, rowNumber+1, topography, visitedMap)
		}
	}
}

func VisitCellAndAllConnectedNeighborsParallel(columnNumber, rowNumber int, topography [][]int, visitedMap *LockableMatrix) {

	// mark it as visited
	visitedMap.VisitsSafe(columnNumber, rowNumber)

	// if water then skip
	if topography[rowNumber][columnNumber] == 0 {
		return
	}

	maxDist := len(topography) - 1

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
					VisitCellAndAllConnectedNeighborsParallel(cell.column, cell.row, topography, visitedMap)
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
		if rowNumber-1 >= 0 {
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
		if rowNumber+1 <= maxDist {
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
	if rowNumber+1 <= maxDist {
		wg.Add(1)
		cellsToVisit <- struct {
			column int
			row    int
		}{columnNumber, rowNumber + 1}
	}

	// column to the right
	if (columnNumber + 1) <= maxDist {

		// upper right
		if rowNumber-1 >= 0 {
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
		if rowNumber+1 <= maxDist {
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
