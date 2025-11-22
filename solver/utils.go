package solver

import (
	"fmt"
	"sync"
)

// LockableMatrix is a thread-safe matrix to track visited cells.
type LockableMatrix struct {
	width       int
	height      int
	mu          sync.RWMutex
	VisitedMap  []bool
	VisitedList []struct {
		Column int
		Row    int
	}
}

// NewLockableMatrix creates a new LockableMatrix with the given dimensions.
func NewLockableMatrix(width, height int) LockableMatrix {
	return LockableMatrix{
		width:      width,
		height:     height,
		VisitedMap: make([]bool, width*height),
		VisitedList: []struct {
			Column int
			Row    int
		}{},
	}
}

// HasVisitedSafe checks if a cell has been visited in a thread-safe manner.
func (l *LockableMatrix) HasVisitedSafe(column, row int) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.HasVisited(column, row)
}

// HasVisited checks if a cell has been visited (not thread-safe).
func (l *LockableMatrix) HasVisited(column, row int) bool {
	cellToCheck := (row * l.width) + column
	if cellToCheck >= len(l.VisitedMap) {
		return false
	}
	return l.VisitedMap[cellToCheck]
}

// VisitsSafe marks a cell as visited in a thread-safe manner.
func (l *LockableMatrix) VisitsSafe(column, row int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Visits(column, row)
}

// Visits marks a cell as visited (not thread-safe).
func (l *LockableMatrix) Visits(column, row int) {
	l.VisitedList = append(l.VisitedList, struct {
		Column int
		Row    int
	}{
		Column: column,
		Row:    row,
	})
	cellToSet := (row * l.width) + column
	l.VisitedMap[cellToSet] = true
}

// RectangleDimensions validates that the topography is a rectangle and returns its dimensions.
func RectangleDimensions(topography [][]int) (rows, columns int, err error) {
	rows = len(topography)
	columns = 0
	for i, row := range topography {
		if i == 0 {
			columns = len(row)
			continue
		}
		if len(row) != columns {
			return 0, 0, fmt.Errorf(
				"topography not rectangular, got %d rows, and row %d had %d columns and expected %d based on previous rows",
				rows, i, len(row), columns,
			)
		}
	}
	return rows, columns, nil
}
