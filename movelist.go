package main

type Move struct {
	ID    int
	Score int
}

type MoveList []Move

func (m MoveList) GetMove(moveIndex int) Move {
	bestIndex := moveIndex
	bestScore := m[moveIndex].Score
	for i := moveIndex + 1; i < len(m); i++ {
		if m[i].Score > bestScore {
			bestIndex = i
			bestScore = m[i].Score
		}
	}
	m[moveIndex], m[bestIndex] = m[bestIndex], m[moveIndex]
	return m[moveIndex]
}