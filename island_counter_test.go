package main

import (
	"testing"
)

type args struct {
	topography [][]int
}

type test struct {
	name    string
	args    args
	want    int
	wantErr bool
}

var tests = []test{
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
			got, _, err := IslandCounter(tt.args.topography, false)
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

func TestIslandCounterParallel(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := IslandCounter(tt.args.topography, true)
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

func runBenchmarkIslandCounter(t test, b *testing.B) {

	for i := 0; i < b.N; i++ {
		got, _, err := IslandCounter(t.args.topography, false)
		if (err != nil) != t.wantErr {
			b.Errorf("IslandCounter() error = %v, wantErr %v", err, t.wantErr)
			return
		}
		if got != t.want {
			b.Errorf("IslandCounter() got = %v, want %v", got, t.want)
		}
	}
}

func runBenchmarkIslandCounterParallel(t test, b *testing.B) {

	for i := 0; i < b.N; i++ {
		got, _, err := IslandCounter(t.args.topography, true)
		if (err != nil) != t.wantErr {
			b.Errorf("IslandCounter() error = %v, wantErr %v", err, t.wantErr)
			return
		}
		if got != t.want {
			b.Errorf("IslandCounter() got = %v, want %v", got, t.want)
		}
	}
}

func BenchmarkIslandCounterTest1(b *testing.B) {
	runBenchmarkIslandCounter(tests[0], b)
}

func BenchmarkIslandCounterParallelTest1(b *testing.B) {
	runBenchmarkIslandCounterParallel(tests[0], b)
}
