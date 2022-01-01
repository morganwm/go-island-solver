package main

import (
	"testing"
)

type args struct {
	topography      [][]int
	breakOnDiagonal bool
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
		args: args{
			topography: [][]int{
				{1, 1, 0, 0, 0},
				{0, 1, 0, 0, 1},
				{1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 1, 0, 1},
			},
			breakOnDiagonal: false,
		},
		want:    5,
		wantErr: false,
	},
	{
		// https://dev.to/rattanakchea/amazons-interview-question-count-island-21h6
		name: "Online (dev)",
		args: args{
			topography: [][]int{
				{0, 1, 0, 1, 0},
				{0, 0, 1, 1, 1},
				{1, 0, 0, 1, 0},
				{0, 1, 1, 0, 0},
				{1, 0, 1, 0, 1},
			},
			breakOnDiagonal: true,
		},
		want:    6,
		wantErr: false,
	},
	{
		// https://dev.to/rattanakchea/amazons-interview-question-count-island-21h6
		name: "Online (dev-Mod)",
		args: args{
			topography: [][]int{
				{0, 1, 0, 1, 0},
				{0, 0, 1, 1, 1},
				{1, 0, 0, 1, 0},
				{0, 1, 1, 0, 0},
				{1, 0, 1, 0, 1},
			},
			breakOnDiagonal: false,
		},
		want:    2,
		wantErr: false,
	},
	{
		// https://medium.com/@obiwankenoobi/interview-question-7-find-the-number-of-islands-1216eff9ede9
		name: "Online (medium)",
		args: args{
			topography: [][]int{
				{1, 0, 0, 0, 0},
				{0, 0, 1, 1, 0},
				{0, 1, 1, 0, 0},
				{1, 1, 0, 0, 1},
				{1, 1, 0, 0, 1},
			},
			breakOnDiagonal: false,
		},
		want:    3,
		wantErr: false,
	},
}

func runIslandCounterTest(tt test, p bool, t *testing.T) {
	got, _, err := IslandCounter(tt.args.topography, false, tt.args.breakOnDiagonal)
	if (err != nil) != tt.wantErr {
		t.Errorf("IslandCounter() error = %v, wantErr %v", err, tt.wantErr)
		return
	}
	if got != tt.want {
		t.Errorf("IslandCounter() got = %v, want %v", got, tt.want)
	}
}

func TestIslandCounter(t *testing.T) {

	for _, testcase := range tests {
		t.Run(testcase.name, func(tt *testing.T) {

			// non-parallel
			tt.Run("S", func(ttt *testing.T) {
				runIslandCounterTest(testcase, false, ttt)
			})

			// parallel
			tt.Run("P", func(ttt *testing.T) {
				runIslandCounterTest(testcase, true, ttt)
			})

		})

	}
}

func runBenchmarkIslandCounter(t test, p bool, b *testing.B) {

	for i := 0; i < b.N; i++ {
		got, _, err := IslandCounter(t.args.topography, p, t.args.breakOnDiagonal)
		if (err != nil) != t.wantErr {
			b.Errorf("IslandCounter() error = %v, wantErr %v", err, t.wantErr)
			return
		}
		if got != t.want {
			b.Errorf("IslandCounter() got = %v, want %v", got, t.want)
		}
	}
}

func BenchmarkIslandCounter(b *testing.B) {
	for _, t := range tests {

		b.Run(t.name, func(bb *testing.B) {

			// non-parallel
			bb.Run("S", func(bbb *testing.B) {
				runBenchmarkIslandCounter(t, false, bbb)
			})

			// parallel
			bb.Run("P", func(bbb *testing.B) {
				runBenchmarkIslandCounter(t, true, bbb)
			})

		})
	}
}
