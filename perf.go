package main

import "fmt"

type Perf struct {
	Depth     int
	States    int
	Leaves    int
	Passes    int
	GameOvers int
}

func PerfTest(state State, depth int) Perf {
	p := Perf{ Depth: depth }
	p.DepthFirstSearch(state, depth)
	return p
}

func (p *Perf) DepthFirstSearch(state State, depth int) {
	p.States++
	
	moves := state.Moves()
	if len(moves) == 0 {
		p.Leaves++
		p.GameOvers++
		return
	}
	
	if depth == 0 {
		p.Leaves++
		return
	}

	if moves[0] == MovePass {
		p.Passes++
	}

	for _, move := range moves {
		next := state.MakeMove(move)
		p.DepthFirstSearch(next, depth-1)
	}
}

func (p Perf) String() string {
	return fmt.Sprintf(
		"Depth: %2d, States: %10d, Leaves: %10d, Passes: %5d, Game overs: %5d",
		p.Depth, p.States, p.Leaves, p.Passes, p.GameOvers,
	)
}