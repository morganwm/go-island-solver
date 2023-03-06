package parallel

import (
	"sync"

	"github.com/morganwm/go-island-solver/constants"
	"github.com/morganwm/go-island-solver/utils"
)

func VisitCellAndAllConnectedNeighborsParallel(
	columnNumber, rowNumber, numberOfRows, numberOfColumns int,
	breakOnDiagonal bool,
	topography [][]int,
	visitedMap *utils.LockableMatrix,
) {

	// mark it as visited
	visitedMap.VisitsSafe(columnNumber, rowNumber)

	// if water then skip
	if topography[rowNumber][columnNumber] == constants.WATER {
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
