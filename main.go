package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
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

// Version can be set at link time to override debug.BuildInfo.Main.Version,
// which is "(devel)" when building from within the module. See
// golang.org/issue/29814 and golang.org/issue/29228.
var Version string

func main() {

	var (
		basicOutPut     = flag.Bool("basic-output", false, "if set the UI will only display out the output of the run and not the UI animation, best for use with non-tty shells")
		parallelFlag    = flag.Bool("parallel", false, "if set the program will run in parallel mode")
		breakOnDiagonal = flag.Bool("brak-on-diagonal", false, "if the flag is set the program will run as if diagonal landmasses are not contiguous")
		versionFlag     = flag.Bool("version", false, "")
	)

	flag.Parse()

	if *versionFlag {
		if Version != "" {
			fmt.Println(Version)
			return
		}
		if buildInfo, ok := debug.ReadBuildInfo(); ok {
			fmt.Println(buildInfo.Main.Version)
			return
		}
		fmt.Println("(unknown)")
		return
	}

	started := time.Now()
	islands, routetaken, err := IslandCounter(topo, *parallelFlag, *breakOnDiagonal)
	timeTaken := time.Since(started)
	if err != nil {
		log.Fatalf("[ERROR] could not count islands: %v", err)
	}

	log.Printf("Found %d Islands traversing %d/%d surfaces, took %s", islands, len(routetaken), len(topo)*len(topo), timeTaken)

	if !*basicOutPut {
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
