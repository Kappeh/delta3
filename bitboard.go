package main

import (
	"math/bits"
)

type Bitboard uint64

func (b Bitboard) Pop() (int, Bitboard) {
	bit := bits.TrailingZeros64(uint64(b))
	return bit, b & ^(1 << bit)
}

func (b Bitboard) ShiftN() Bitboard {
	return (b & 0xFFFFFFFFFFFFFF00) >> 8
}

func (b Bitboard) ShiftE() Bitboard {
	return (b & 0x7F7F7F7F7F7F7F7F) << 1
}

func (b Bitboard) ShiftS() Bitboard {
	return (b & 0x00FFFFFFFFFFFFFF) << 8
}

func (b Bitboard) ShiftW() Bitboard {
	return (b & 0xFEFEFEFEFEFEFEFE) >> 1
}

func (b Bitboard) ShiftNE() Bitboard {
	return (b & 0x7F7F7F7F7F7F7F00) >> 7
}

func (b Bitboard) ShiftSE() Bitboard {
	return (b & 0x007F7F7F7F7F7F7F) << 9
}

func (b Bitboard) ShiftSW() Bitboard {
	return (b & 0x00FEFEFEFEFEFEFE) << 7
}

func (b Bitboard) ShiftNW() Bitboard {
	return (b & 0xFEFEFEFEFEFEFE00) >> 9
}