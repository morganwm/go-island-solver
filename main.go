package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/morganwm/go-island-solver/core"
	"github.com/morganwm/go-island-solver/core/traversals"
	"github.com/morganwm/go-island-solver/ui"
)

var (
	/*
		Version can be set at link time to override debug.BuildInfo.Main.Version,
		which is "(devel)" when building from within the module. See
		golang.org/issue/29814 and golang.org/issue/29228.
	*/
	Version string

	basicOutPut     = flag.Bool("basic-output", false, "if set the UI will only display out the output of the run and not the UI animation, best for use with non-tty shells")
	modeFlag        = flag.String("mode", "dfs", fmt.Sprintf("the mode to run the program in: %v", traversals.Traversers.GetKeys()))
	breakOnDiagonal = flag.Bool("break-on-diagonal", false, "if the flag is set the program will run as if diagonal landmasses are not contiguous")
	versionFlag     = flag.Bool("version", false, "show the version information")
	helpFlag        = flag.Bool("help", false, "shows this message")
	speedFlag       = flag.Int("speed", 1000, "the number of seconds per refresh when using the fancy UI")

	/*
		DEFAULT_TOPO is the value used when to topography is provided to the island solver
	*/
	DEFAULT_TOPO = [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
	}
)

func main() {

	topo := DEFAULT_TOPO

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}

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

	if *speedFlag < 1 {
		fmt.Printf("invalid speed: %d, must be greater than 1", *speedFlag)
	}

	started := time.Now()
	islands, routetaken, err := core.IslandCounter(topo,
		core.IslandCounterOptions{
			BreakOnDiagonal: *breakOnDiagonal,
		},
		core.IslandCounterSettings{
			Mode: *modeFlag,
		},
	)
	timeTaken := time.Since(started)
	if err != nil {
		log.Fatalf("[ERROR] could not count islands: %v", err)
	}

	log.Printf("Found %d Islands traversing %d/%d surfaces, took %s", islands, len(routetaken), len(topo)*len(topo), timeTaken)

	// break early for basic-output
	if *basicOutPut {
		return
	}

	p := tea.NewProgram(&ui.IslandSolverModel{
		Speed:      time.Millisecond * time.Duration(*speedFlag),
		Topography: topo,
		Routetaken: routetaken,
	})
	if _, err := p.Run(); err != nil {
		log.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
