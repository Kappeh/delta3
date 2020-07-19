package main

import (
	"errors"
)

const (
	MoveBufferSize = 60
	MovePass       = -1

	Black = iota
	White
)

type State struct {
	DiscsBlack     Bitboard
	DiscsWhite     Bitboard
	Player         int
	OpponentPassed bool

	moveBuffer   [MoveBufferSize]int
	captureTable CaptureTable

	// Hash values (including symmetry)
}

func NewState(discsBlack, discsWhite Bitboard, player int) (State, error) {
	s := State{
		DiscsBlack:     discsBlack,
		DiscsWhite:     discsWhite,
		Player:         player,
		OpponentPassed: false,

		moveBuffer: [MoveBufferSize]int{},
	}

	switch player {
	case Black:
		s.captureTable = NewCaptureTable(discsBlack, discsWhite)
	case White:
		s.captureTable = NewCaptureTable(discsWhite, discsBlack)
	default:
		return s, errors.New("invalid player")
	}

	return s, nil
}

func (s State) Moves() []int {
	moves := s.captureTable.Moves()
	
	if moves == 0 {
		if s.OpponentPassed {
			return s.moveBuffer[:0]
		}
		s.moveBuffer[0] = MovePass
		return s.moveBuffer[:1]
	}

	var i int
	for ; i < MoveBufferSize && moves != 0; i++ {
		s.moveBuffer[i], moves = moves.Pop()
	}

	return s.moveBuffer[:i]
}

func (s State) MakeMove(move int) State {
	var disc, captures Bitboard
	
	if move != MovePass {
		disc     = Bitboard(1 << move)
		captures = s.captureTable.Captures(disc)
	} else {
		s.OpponentPassed = true
	}

	if s.Player == Black {
		s.DiscsBlack  |=  captures | disc
		s.DiscsWhite  &= ^captures
		s.Player       =  White
		s.captureTable =  NewCaptureTable(s.DiscsWhite, s.DiscsBlack)
	} else {
		s.DiscsBlack  &= ^captures
		s.DiscsWhite  |=  captures | disc
		s.Player       =  Black
		s.captureTable =  NewCaptureTable(s.DiscsBlack, s.DiscsWhite)
	}

	return s
}

func (s State) Evaluate() int {
	return 0
}

func (s State) String() string {
	const black, white, empty = "B", "W", "-"
	result := ""

	for i := 0; i < 64; i++ {
		disc := Bitboard(1 << i)
		switch {
		case s.DiscsBlack & disc != 0:
			result += black
		case s.DiscsWhite & disc != 0:
			result += white
		default:
			result += empty
		}
		result += " "

		if i%8 == 7 {
			result += "\n"
		}
	}

	if s.Player == Black {
		result += "Black to play"
	} else {
		result += "White to play"
	}

	return result
}