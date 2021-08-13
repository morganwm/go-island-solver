package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var topo = [][]int{
	{1, 1, 0, 0, 0},
	{0, 1, 0, 0, 1},
	{1, 0, 0, 1, 1},
	{0, 0, 0, 0, 0},
	{1, 0, 1, 0, 1},
}

// func not_main() {
// 	log.Printf("[DEBUG] Started")

// 	topo := [][]int{
// 		{1, 1, 0, 0, 0},
// 		{0, 1, 0, 0, 1},
// 		{1, 0, 0, 1, 1},
// 		{0, 0, 0, 0, 0},
// 		{1, 0, 1, 0, 1},
// 	}

// 	islands, routetaken, err := IslandCounter(topo)
// 	if err != nil {
// 		log.Fatalf("[ERROR] could not count islands: %v", err)
// 	}

// 	log.Printf("[DEBUG] found %d islands", islands)

// 	log.Printf("[DEBUG] Done")
// }

func main() {
	islands, routetaken, err := IslandCounter(topo, true)
	if err != nil {
		log.Fatalf("[ERROR] could not count islands: %v", err)
	}

	log.Printf("Found %d Islands traversing %d/%d surfaces", islands, len(routetaken), len(topo)*len(topo))

	displayMap := make([][]string, len(topo))
	for i := range displayMap {
		var row []string
		for _, surfaceTexture := range topo[i] {
			row = append(row, fmt.Sprintf("%d", surfaceTexture))
		}

		displayMap[i] = row
	}

	p := tea.NewProgram(model{
		displayableMap: displayMap,
		topography:     topo,
		routetaken:     routetaken,
		step:           0,
	})
	if err := p.Start(); err != nil {
		log.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}

type tickMsg time.Time

type model struct {
	displayableMap [][]string
	topography     [][]int
	routetaken     []struct {
		Column int
		Row    int
	}
	step int
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tickMsg:
		if m.step == len(m.routetaken) {
			return m, tea.Quit
		}

		m.step++

		// is not 0 then set the previous step location equal to something else
		if m.step != 0 {
			previousStepLocation := m.routetaken[m.step-1]
			valueToSet := "#"
			if topo[previousStepLocation.Row][previousStepLocation.Column] == 0 {
				valueToSet = "_"
			}
			m.displayableMap[previousStepLocation.Row][previousStepLocation.Column] = valueToSet
		}

		if m.step < len(m.routetaken) {
			// set current step location to 'X'
			currentlyStepLocation := m.routetaken[m.step]
			m.displayableMap[currentlyStepLocation.Row][currentlyStepLocation.Column] = "X"
		}

		return m, tickCmd()

	default:
		return m, nil
	}
}

func (m model) View() string {

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

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
