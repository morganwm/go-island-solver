package main

import (
	"testing"
)

type args struct {
	topography [][]int
}

var tests = []struct {
	name    string
	args    args
	want    int
	wantErr bool
}{
	{
		name: "Matts Test",
		args: args{topography: [][]int{
			{1, 1, 0, 0, 0},
			{0, 1, 0, 0, 1},
			{1, 0, 0, 1, 1},
			{0, 0, 0, 0, 0},
			{1, 0, 1, 0, 1},
		}},
		want:    5,
		wantErr: false,
	},
}

func TestIslandCounter(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := IslandCounter(tt.args.topography)
			if (err != nil) != tt.wantErr {
				t.Errorf("IslandCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IslandCounter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIslandCounter(b *testing.B) {

	testToBenchmark := tests[0]

	for i := 0; i < b.N; i++ {
		got, _, err := IslandCounter(testToBenchmark.args.topography)
		if (err != nil) != testToBenchmark.wantErr {
			b.Errorf("IslandCounter() error = %v, wantErr %v", err, testToBenchmark.wantErr)
			return
		}
		if got != testToBenchmark.want {
			b.Errorf("IslandCounter() got = %v, want %v", got, testToBenchmark.want)
		}
	}
}
