package main

type Searcher struct {}

func (s *Searcher) search2(depth int, node State) Boundary {
	const Win = 100000000
	return s.PVSearch(BoundaryLower(-Win), BoundaryUpper(Win), depth, node)
}

// Alpha is lower bound
// Beta  is upper bound
// PVSearch returns an exact or lower bound for this nodes best score
// TODO: Move ordering

func (s *Searcher) PVSearch(alpha, beta Boundary, depth int, state State) Boundary {
	// Get possible actions
	moves := state.Moves()
	// If there are no actions,
	//     this is a terminal node, the game is over
	//     return the utility of the node (Exact Value)
	if len(moves) == 0 {
		return BoundaryExact(state.Evaluate())
	}
	// If depth is 0,
	//     return the evaluation of this node (Lower Bound)
	if depth == 0 {
		return BoundaryLower(state.Evaluate())
	}

	// First move
	move     := moves.GetMove(0)
	newState := state.MakeMove(move)
	// Full search
	bestScore := -s.PVSearch(-beta, -alpha, depth-1, newState)

	// If the move is greater than beta,
	//     return the score as a lower bound
	if bestScore > beta {
		return bestScore.ForceLower()
	}

	// exactScore tracks if all the results
	// from the searches from this node were
	// exact bounds
	exactScore := bestScore.IsExact()

	for i := 1; i < len(moves); i++ {
		move     = moves.GetMove(i)
		newState = state.MakeMove(move)
		score   := -s.zwSearch(-alpha, depth-1, newState)

		// If the value is within alpha and beta,
		//     re-search
		if score > alpha && score < beta {
			score      = -s.PVSearch(-beta, -alpha, depth-1, newState)
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
		return bestScore.ForceExact()
	}
	// Otherwise, return a lower bound
	return bestScore.ForceLower()
}

func (s *Searcher) zwSearch(beta Boundary, depth int, state State) Boundary {
	// alpha := beta - 1
	moves := state.Moves()
	if len(moves) == 0 {
		return BoundaryExact(state.Evaluate())
	}
	if depth == 0 {
		return BoundaryLower(state.Evaluate())
	}
	for _, move := range moves {
		newState := state.MakeMove(move)
		score    := -s.zwSearch(1 - beta, depth - 1, newState)
		if score >= beta {
			// fail-hard beta cut-off
			return beta
		}
	}
	// fail-hard return alpha
	return beta - 1
}