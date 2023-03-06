package utils

import "sync"

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

func (l *LockableMatrix) HasVisitedSafe(column, row int) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.HasVisited(column, row)
}

func (l *LockableMatrix) HasVisited(column, row int) bool {

	cellToCheck := (row * l.width) + column
	if cellToCheck >= len(l.VisitedMap) {
		return false
	}

	return l.VisitedMap[cellToCheck]
}

func (l *LockableMatrix) VisitsSafe(column, row int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Visits(column, row)
}

func (l *LockableMatrix) Visits(column, row int) {

	l.VisitedList = append(l.VisitedList, struct {
		Column int
		Row    int
	}{
		Column: column,
		Row:    row,
	})

	cellToSet := (row * l.width) + (column)

	l.VisitedMap[cellToSet] = true
}
