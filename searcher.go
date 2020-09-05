package main

const MaxDepth = 100

type Searcher struct {
	RootNode State

	TranspositionTable *TranspositionTable
	Killers             [MaxDepth * 2]Move
	History             [MaxDepth][64 + 1]int
	PVTable             PVTable
}

type NodeInfo struct {
	// Score is an alpha value when a  lower bound
	// Score is a  beta  value when an upper bound
	SearchDepth int
	PVMove      Move
}

// Alpha is lower bound
// Beta  is upper bound
// PVSearch returns an exact or lower bound for this nodes best score
// TODO: Move ordering
func (s *Searcher) pvSearch(alpha, beta Boundary, depth int, state State) (Boundary, bool) {
	// Probe hash table
	
	// Get possible actions
	moves := state.Moves()
	// If there are no actions,
	//     this is a terminal node, the game is over
	//     return the utility of the node (Exact Value)
	if len(moves) == 0 {
		return BoundaryExact(state.Evaluate()), true
	}
	// If depth is 0,
	//     return the evaluation of this node (Lower Bound)
	if depth == 0 {
		return BoundaryLower(state.Evaluate()), true
	}

	// First move
	move     := moves.GetMove(0)
	newState := state.MakeMove(move)
	// Full search
	bestScore, _ := s.pvSearch(-beta, -alpha, depth-1, newState)
	bestScore     = -bestScore

	// If the move is greater than beta,
	//     return the score as a lower bound
	if bestScore > beta {
		return bestScore.ForceLower(), true
	}

	// exactScore tracks if all the results
	// from the searches from this node were
	// exact bounds
	exactScore := bestScore.IsExact()

	for i := 1; i < len(moves); i++ {
		move      = moves.GetMove(i)
		newState  = state.MakeMove(move)
		score, _ := s.zwSearch(-alpha, depth-1, newState)
		score     = -score

		// If the value is within alpha and beta,
		//     re-search
		if score > alpha && score < beta {
			score, _   = s.pvSearch(-beta, -alpha, depth-1, newState)
			score      = -score
			exactScore = exactScore && score.IsExact()
		}

		// If the score if better than alpha,
		//     update alpha
		if score > alpha {
			alpha = score
		}

		// If the score is better than the best score,
		//     update the besst score
		if score > bestScore {
			bestScore = score
		}
		
		// If score (or alpha) is greater than beta,
		//     beta-cutoff
		if score >= beta {
			break
		}
	}

	// If all of the searches performed were exact searches,
	//     return the score as an exact score
	//
	//     note: in the case that there are unsearched
	//     moves, we can return an exact bound because it's
	//     greater than beta anyway thus will not be considered
	//     in the parent recurive call as it will be less than alpha
	//     (Not technically true but it will get removes while
	//     backtracking up the tree)
	if exactScore {
		return bestScore.ForceExact(), false
	}
	// Otherwise, return a lower bound
	return bestScore.ForceLower(), false
}

func (s *Searcher) zwSearch(beta Boundary, depth int, state State) (Boundary, bool) {
	// alpha := beta - 1
	moves := state.Moves()
	if len(moves) == 0 {
		return BoundaryExact(state.Evaluate()), true
	}
	if depth == 0 {
		return BoundaryLower(state.Evaluate()), true
	}
	for _, move := range moves {
		newState := state.MakeMove(move)
		score, _ := s.zwSearch(1 - beta, depth - 1, newState)
		score     = -score
		if score >= beta {
			// fail-hard beta cut-off
			return beta, true
		}
	}
	// fail-hard return alpha
	return beta - 1, true
}