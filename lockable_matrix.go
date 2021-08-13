package main

import "sync"

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

func (l *LockableMatrix) HasVisitedSafe(column, row int) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.HasVisited(column, row)
}

func (l *LockableMatrix) HasVisited(column, row int) bool {

	cellToCheck := (row * l.size) + column
	if cellToCheck >= len(l.visitedMap) {
		return false
	}

	return l.visitedMap[cellToCheck]
}

func (l *LockableMatrix) VisitsSafe(column, row int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Visits(column, row)
}

func (l *LockableMatrix) Visits(column, row int) {

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
