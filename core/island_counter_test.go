package core

import (
	"testing"

	"github.com/morganwm/go-island-solver/constants"
)

type args struct {
	topography [][]int
	options    IslandCounterOptions
}

type test struct {
	name    string
	args    args
	want    int
	wantErr bool
}

type testSuite struct {
	name  string
	cases []test
}

var testEdgeCases = []test{
	{
		name: "[]",
		args: args{
			topography: [][]int{},
			options:    IslandCounterOptions{BreakOnDiagonal: false},
		},
		want:    0,
		wantErr: false,
	},
	{
		name: "[[]]",
		args: args{
			topography: [][]int{{}},
			options:    IslandCounterOptions{BreakOnDiagonal: false},
		},
		want:    0,
		wantErr: false,
	},
	{
		name: "[[0]]",
		args: args{
			topography: [][]int{{0}},
			options:    IslandCounterOptions{BreakOnDiagonal: false},
		},
		want:    0,
		wantErr: false,
	},
	{
		name: "[[1]]",
		args: args{
			topography: [][]int{{1}},
			options:    IslandCounterOptions{BreakOnDiagonal: false},
		},
		want:    1,
		wantErr: false,
	},
	{
		name: "long",
		args: args{
			topography: [][]int{{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
			options:    IslandCounterOptions{BreakOnDiagonal: false},
		},
		want:    2,
		wantErr: false,
	},
	{
		name: "tall",
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
			options: IslandCounterOptions{BreakOnDiagonal: false},
		},
		want:    2,
		wantErr: false,
	},
	{
		name: "all land",
		args: args{
			topography: [][]int{
				{1, 1},
				{1, 1},
			},
			options: IslandCounterOptions{BreakOnDiagonal: false},
		},
		want:    1,
		wantErr: false,
	},
}

var testCasesFunctional = []test{
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
			options: IslandCounterOptions{BreakOnDiagonal: false},
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
			options: IslandCounterOptions{BreakOnDiagonal: true},
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
			options: IslandCounterOptions{BreakOnDiagonal: false},
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
			options: IslandCounterOptions{BreakOnDiagonal: false},
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
			options: IslandCounterOptions{BreakOnDiagonal: true},
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
			options: IslandCounterOptions{BreakOnDiagonal: true},
		},
		want:    1,
		wantErr: false,
	},
}

var settingVariants = map[string]IslandCounterSettings{
	"series_ifs":  {Mode: constants.SERIES_IFS},
	"series_loop": {Mode: constants.SERIES_LOOP},
	"parallel":    {Mode: constants.PARALLEL},
}

var testSuites = map[string][]test{
	"edge":       testEdgeCases,
	"functional": testCasesFunctional,
}

func runIslandCounterTest(tt test, s IslandCounterSettings, t testing.TB) {
	got, _, err := IslandCounter(tt.args.topography, tt.args.options, s)
	if (err != nil) != tt.wantErr {
		t.Errorf("IslandCounter() error = %v, wantErr %v", err, tt.wantErr)
		return
	}
	if got != tt.want {
		t.Errorf("IslandCounter() got = %v, want %v", got, tt.want)
	}
}

func runBenchmarkIslandCounter(t test, s IslandCounterSettings, b *testing.B) {
	for i := 0; i < b.N; i++ {
		runIslandCounterTest(t, s, b)
	}
}

func TestIslandCounter(t *testing.T) {

	for suite, tests := range testSuites {
		t.Run(suite, func(tSuite *testing.T) {

			for _, testcase := range tests {
				tSuite.Run(testcase.name, func(tCase *testing.T) {

					for variant, setting := range settingVariants {
						tCase.Run(variant, func(tVariant *testing.T) {
							runIslandCounterTest(testcase, setting, tVariant)
						})
					}

				})
			}

		})
	}

}

func BenchmarkIslandCounter(b *testing.B) {

	for suite, tests := range testSuites {
		b.Run(suite, func(bSuite *testing.B) {

			for _, testcase := range tests {
				bSuite.Run(testcase.name, func(bCase *testing.B) {

					for variant, settings := range settingVariants {
						bCase.Run(variant, func(bVariant *testing.B) {
							runBenchmarkIslandCounter(testcase, settings, bVariant)
						})
					}

				})
			}

		})
	}

}
