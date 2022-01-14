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
	{
		// https://medium.com/javarevisited/day-33-number-of-islands-80ecd0490fe3
		name: "amzn (rect)",
		args: args{
			topography: [][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			breakOnDiagonal: true,
		},
		want:    1,
		wantErr: false,
	},
	{
		// https://medium.com/javarevisited/day-33-number-of-islands-80ecd0490fe3
		name: "amzn (rect - rot)",
		args: args{
			topography: [][]int{
				{0, 1, 1, 1},
				{0, 1, 1, 1},
				{0, 0, 0, 1},
				{0, 0, 1, 1},
				{0, 0, 0, 0},
			},
			breakOnDiagonal: true,
		},
		want:    1,
		wantErr: false,
	},
	{
		name: "custom []",
		args: args{
			topography:      [][]int{},
			breakOnDiagonal: false,
		},
		want:    0,
		wantErr: false,
	},
	{
		name: "custom [[]]",
		args: args{
			topography:      [][]int{{}},
			breakOnDiagonal: false,
		},
		want:    0,
		wantErr: false,
	},
	{
		name: "custom [[0]]",
		args: args{
			topography:      [][]int{{0}},
			breakOnDiagonal: false,
		},
		want:    0,
		wantErr: false,
	},
	{
		name: "custom [[1]]",
		args: args{
			topography:      [][]int{{1}},
			breakOnDiagonal: false,
		},
		want:    1,
		wantErr: false,
	},
	{
		name: "custom long",
		args: args{
			topography:      [][]int{{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
			breakOnDiagonal: false,
		},
		want:    2,
		wantErr: false,
	},
	{
		name: "custom tall",
		args: args{
			topography: [][]int{
				{0},
				{0},
				{0},
				{1},
				{0},
				{0},
				{0},
				{0},
				{0},
				{0},
				{0},
				{0},
				{0},
				{0},
				{1},
			},
			breakOnDiagonal: false,
		},
		want:    2,
		wantErr: false,
	},
	{
		name: "custom all land",
		args: args{
			topography: [][]int{
				{1, 1},
				{1, 1},
			},
			breakOnDiagonal: false,
		},
		want:    1,
		wantErr: false,
	},
}

func runIslandCounterTest(tt test, p bool, t testing.TB) {
	got, _, err := IslandCounter(tt.args.topography, p, tt.args.breakOnDiagonal)
	if (err != nil) != tt.wantErr {
		t.Errorf("IslandCounter() error = %v, wantErr %v", err, tt.wantErr)
		return
	}
	if got != tt.want {
		t.Errorf("IslandCounter() got = %v, want %v", got, tt.want)
	}
}

func runBenchmarkIslandCounter(t test, p bool, b *testing.B) {
	for i := 0; i < b.N; i++ {
		runIslandCounterTest(t, p, b)
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

func BenchmarkIslandCounter(b *testing.B) {
	for _, testcase := range tests {

		b.Run(testcase.name, func(bb *testing.B) {

			// non-parallel
			bb.Run("S", func(bbb *testing.B) {
				runBenchmarkIslandCounter(testcase, false, bbb)
			})

			// parallel
			bb.Run("P", func(bbb *testing.B) {
				runBenchmarkIslandCounter(testcase, true, bbb)
			})

		})
	}
}
