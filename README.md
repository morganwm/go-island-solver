# Go Island Solver

[![Test](https://github.com/morganwm/go-island-solver/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/morganwm/go-island-solver/actions/workflows/test.yml)
[![Upload Release Asset](https://github.com/morganwm/go-island-solver/actions/workflows/release.yml/badge.svg)](https://github.com/morganwm/go-island-solver/actions/workflows/release.yml)

A simple Go program that counts the number of islands in a 2D grid (where 1 is land and 0 is water).

It includes a terminal-based UI to visualize the solving process.

A list of releases can be found [here](https://github.com/morganwm/go-island-solver/releases)

## Usage

Run the solver with the default map:
```bash
go run main.go
```

### Options

- `-parallel`: Run in parallel mode (uses goroutines).
- `-speed <ms>`: Set the animation speed in milliseconds (default 1000).
- `-break-on-diagonal`: Treat diagonal landmasses as separate islands.
- `-basic-output`: Disable the UI and just show the result.