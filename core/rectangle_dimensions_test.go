package core

import (
	"fmt"
	"testing"
)

func TestRectangleDimensions(t *testing.T) {
	tests := []struct {
		name            string
		input           [][]int
		expectedRows    int
		expectedColumns int
		expectedError   error
	}{
		{
			name:            "empty topography",
			input:           [][]int{},
			expectedRows:    0,
			expectedColumns: 0,
			expectedError:   nil,
		},
		{
			name: "single row",
			input: [][]int{
				{1, 0, 1},
			},
			expectedRows:    1,
			expectedColumns: 3,
			expectedError:   nil,
		},
		{
			name: "single column",
			input: [][]int{
				{1},
				{0},
				{1},
			},
			expectedRows:    3,
			expectedColumns: 1,
			expectedError:   nil,
		},
		{
			name: "rectangular topography",
			input: [][]int{
				{1, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
			},
			expectedRows:    3,
			expectedColumns: 3,
			expectedError:   nil,
		},
		{
			name: "non-rectangular topography",
			input: [][]int{
				{1, 0, 1},
				{0, 1},
				{1, 0, 1},
			},
			expectedRows:    0,
			expectedColumns: 0,
			expectedError:   fmt.Errorf("topography not rectangular, got 3 rows, and row 1 had 2 columns and expected 3 based on previous rows"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows, columns, err := RectangleDimensions(tt.input)
			if rows != tt.expectedRows {
				t.Errorf("expected %d rows, but got %d", tt.expectedRows, rows)
			}
			if columns != tt.expectedColumns {
				t.Errorf("expected %d columns, but got %d", tt.expectedColumns, columns)
			}
			if err != nil && tt.expectedError == nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if err == nil && tt.expectedError != nil {
				t.Errorf("expected error %v, but got no error", tt.expectedError)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, but got %v", tt.expectedError, err)
			}
		})
	}
}
