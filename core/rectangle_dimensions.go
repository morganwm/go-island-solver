package core

import "fmt"

func RectangleDimensions(topography [][]int) (rows, columns int, e error) {
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
