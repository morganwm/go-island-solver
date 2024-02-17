package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
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
	DisplayableMap [][]string
	Topography     [][]int
	Routetaken     []struct {
		Column int
		Row    int
	}
	Step int
}

// Init implements tea.Model.
func (m *IslandSolverModel) Init() tea.Cmd {
	return tickCmd(m.Speed)
}

// Update implements tea.Model.
func (m *IslandSolverModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case time.Time:
		if m.Step == len(m.Routetaken) {
			return m, tea.Quit
		}

		m.Step++

		// is not 0 then set the previous step location equal to something else
		if m.Step != 0 {
			previousStepLocation := m.Routetaken[m.Step-1]
			valueToSet := "#"
			if m.Topography[previousStepLocation.Row][previousStepLocation.Column] == 0 {
				valueToSet = "_"
			}
			m.DisplayableMap[previousStepLocation.Row][previousStepLocation.Column] = valueToSet
		}

		if m.Step < len(m.Routetaken) {
			// set current step location to 'X'
			currentlyStepLocation := m.Routetaken[m.Step]
			m.DisplayableMap[currentlyStepLocation.Row][currentlyStepLocation.Column] = "X"
		}

		return m, tickCmd(m.Speed)

	default:
		return m, nil
	}
}

// View implements tea.Model.
func (m *IslandSolverModel) View() string {
	viewString := "\n"

	for _, row := range m.DisplayableMap {
		for _, surfaceTexture := range row {
			viewString += "\t"
			viewString += surfaceTexture
		}
		viewString += "\n"
	}

	return viewString
}
