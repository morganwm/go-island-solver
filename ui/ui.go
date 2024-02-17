package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/morganwm/go-island-solver/constants"
)

var _ tea.Model = (*IslandSolverModel)(nil)
var _ tea.Model = &IslandSolverModel{}

func tickCmd(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return time.Time(t)
	})
}

// Model is the application's internal state. It holds the current step and the route taken
type IslandSolverModel struct {
	Speed          time.Duration
	displayableMap [][]string
	Topography     [][]int
	Routetaken     []struct {
		Column int
		Row    int
	}
	step int
}

// Init implements tea.Model.
func (m *IslandSolverModel) Init() tea.Cmd {

	// initialize step to the beginning
	m.step = 0

	// convert topo to a map we can display
	m.displayableMap = make([][]string, len(m.Topography))
	for i := range m.displayableMap {
		var row []string
		for _, surfaceTexture := range m.Topography[i] {
			row = append(row, fmt.Sprintf("%d", surfaceTexture))
		}

		m.displayableMap[i] = row
	}

	// kick off the call back event
	return tickCmd(m.Speed)
}

// Update implements tea.Model.
func (m *IslandSolverModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {

	case tea.KeyMsg:
		return m, tea.Quit

	case time.Time:
		if m.step == len(m.Routetaken) {
			return m, tea.Quit
		}

		m.step++

		// is not 0 then set the previous step location equal to something else
		if m.step != 0 {
			previousStepLocation := m.Routetaken[m.step-1]

			switch m.Topography[previousStepLocation.Row][previousStepLocation.Column] {

			case constants.WATER:
				m.displayableMap[previousStepLocation.Row][previousStepLocation.Column] = "_"

			case constants.LAND:
				m.displayableMap[previousStepLocation.Row][previousStepLocation.Column] = "#"
			}
		}

		if m.step < len(m.Routetaken) {
			// set current step location to 'X'
			currentlyStepLocation := m.Routetaken[m.step]
			// m.DisplayableMap[currentlyStepLocation.Row][currentlyStepLocation.Column] = "X"
			m.displayableMap[currentlyStepLocation.Row][currentlyStepLocation.Column] = fmt.Sprintf(
				"[%d]",
				m.Topography[currentlyStepLocation.Row][currentlyStepLocation.Column],
			)
		}

		return m, tickCmd(m.Speed)

	default:
		return m, nil
	}
}

// View implements tea.Model.
func (m *IslandSolverModel) View() string {
	viewString := "\n"

	for _, row := range m.displayableMap {
		for _, surfaceTexture := range row {
			viewString += "\t"
			viewString += surfaceTexture
		}
		viewString += "\n"
	}

	return viewString
}
