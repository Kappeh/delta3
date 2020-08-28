package main

import (
	"math/bits"
)

const (
	MoveBufferSize = 60
	MoveIDPass     = -1

	Black = iota
	White
)

type State struct {
	DiscsBlack     Bitboard
	DiscsWhite     Bitboard
	Player         int
	OpponentPassed bool

	moveBuffer   [MoveBufferSize]Move
	captureTable CaptureTable
	hashes       SymetryHashList
}

func NewState(discsBlack, discsWhite Bitboard, player int) State {
	s := State{
		DiscsBlack:     discsBlack,
		DiscsWhite:     discsWhite,
		Player:         player,
		OpponentPassed: false,

		moveBuffer: [MoveBufferSize]Move{},
		hashes:     NewSymetryHashList(),
	}

	s.hashes = s.hashes.ApplyBitboardBlack(discsBlack)
	s.hashes = s.hashes.ApplyBitboardWhite(discsWhite)

	if player == Black {
		s.captureTable = NewCaptureTable(discsBlack, discsWhite)
	} else {
		s.captureTable = NewCaptureTable(discsWhite, discsBlack)
	}
	
	return s
}

func (s State) Moves() MoveList {
	moves := s.captureTable.Moves()
	
	if moves == 0 {
		if s.OpponentPassed {
			return s.moveBuffer[:0]
		}
		s.moveBuffer[0].ID = MoveIDPass
		return s.moveBuffer[:1]
	}

	var i int
	for ; i < MoveBufferSize && moves != 0; i++ {
		s.moveBuffer[i].ID, moves = moves.Pop()
	}

	return s.moveBuffer[:i]
}

func (s State) MakeMove(move Move) State {
	var disc, captures Bitboard
	
	if move.ID != MoveIDPass {
		disc     = Bitboard(1 << move.ID)
		captures = s.captureTable.Captures(disc)
	} else {
		s.OpponentPassed = true
	}

	if s.Player == Black {
		s.DiscsBlack  |=  captures | disc
		s.DiscsWhite  &= ^captures
		s.Player       =  White
		s.captureTable =  NewCaptureTable(s.DiscsWhite, s.DiscsBlack)

		s.hashes = s.hashes.ApplyBitboardBlack(captures | disc)
		s.hashes = s.hashes.ApplyBitboardWhite(captures)
	} else {
		s.DiscsBlack  &= ^captures
		s.DiscsWhite  |=  captures | disc
		s.Player       =  Black
		s.captureTable =  NewCaptureTable(s.DiscsBlack, s.DiscsWhite)

		s.hashes = s.hashes.ApplyBitboardBlack(captures)
		s.hashes = s.hashes.ApplyBitboardWhite(captures | disc)
	}

	return s
}

func (s State) Evaluate() int {
	// Very crude evaluation function for now
	// Features
	// - DiscCount                  simple implimentation
	// - Mobility                   simple implimentation
	// - DiscPositionScore(static)  simple implimentation
	// - DiscPositionScore(dynamic) not implimented
	// - InteriorDiscCount          not implimented
	
	const (
		discCountScoreWeight    = 1
		mobilityScoreWeight     = 50
		discPositionScoreWeight = 10
	)
	
	var (
		discPositionScoreTable  = [64]int{
			100, -30, 6, 2, 2, 6, -30, 100,
			-30, -50, 0, 0, 0, 0, -50, -30,
			  6,   0, 0, 0, 0, 0,   0,   6,
			  2,   0, 0, 0, 0, 0,   0,   2,
			  2,   0, 0, 0, 0, 0,   0,   2,
			  6,   0, 0, 0, 0, 0,   0,   6,
			-30, -50, 0, 0, 0, 0, -50, -30,
			100, -30, 6, 2, 2, 6, -30, 100,
		}

		discCountScore       int
		mobilityScore        int
		discPositionScore    int
		opponentCaptureTable CaptureTable
	)
	
	discsBlack     := s.DiscsBlack
	discsWhite     := s.DiscsWhite
	discCountBlack := bits.OnesCount64(uint64(s.DiscsBlack))
	discCountWhite := bits.OnesCount64(uint64(s.DiscsWhite))

	var discPosition int
	for discsBlack != 0 {
		discPosition, discsBlack = discsBlack.Pop()
		discPositionScore       += discPositionScoreTable[discPosition]
	}
	for discsWhite != 0 {
		discPosition, discsWhite = discsWhite.Pop()
		discPositionScore       -= discPositionScoreTable[discPosition]
	}

	if s.Player == Black {
		discCountScore       = discCountBlack - discCountWhite
		opponentCaptureTable = NewCaptureTable(s.DiscsWhite, s.DiscsBlack)
	} else {
		discCountScore       = discCountWhite - discCountBlack
		opponentCaptureTable = NewCaptureTable(s.DiscsBlack, s.DiscsWhite)

		discPositionScore = -discPositionScore
	}

	movesPlayer       := s.captureTable.Moves()
	movesOpponent     := opponentCaptureTable.Moves()
	moveCountPlayer   := bits.OnesCount64(uint64(movesPlayer))
	moveCountOpponent := bits.OnesCount64(uint64(movesOpponent))
	mobilityScore      = moveCountPlayer - moveCountOpponent

	return discCountScore    * discCountScoreWeight +
		   mobilityScore     * mobilityScoreWeight  +
		   discPositionScore * discPositionScoreWeight
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