package main

import (
	"fmt"
	"sync"
)

type LockableMatrix struct {
	size        int
	mu          sync.RWMutex
	visitedMap  []bool
	visitedList []struct {
		Column int
		Row    int
	}
}

func NewLockableMatrix(size int) LockableMatrix {
	return LockableMatrix{
		size:       size,
		visitedMap: make([]bool, size*size),
		visitedList: []struct {
			Column int
			Row    int
		}{},
	}
}

func (l *LockableMatrix) HasVisited(column, row int) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	cellToCheck := (row * l.size) + column
	if cellToCheck >= len(l.visitedMap) {
		return false
	}

	return l.visitedMap[cellToCheck]
}

func (l *LockableMatrix) Visits(column, row int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.visitedList = append(l.visitedList, struct {
		Column int
		Row    int
	}{
		Column: column,
		Row:    row,
	})

	cellToSet := (row * l.size) + (column)

	l.visitedMap[cellToSet] = true
}

func IslandCounter(topography [][]int) (int, []struct {
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
			VisitCellAndAllConnectedNeighbors(columnNumber, rowNumber, topography, &visitedMap)

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
