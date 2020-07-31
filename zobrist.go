package main

import "math/rand"

type ZobristHasher struct {
	HashTableBlack [64]uint64
	HashTableWhite [64]uint64
}

func NewZobristHasher(r *rand.Rand) ZobristHasher {
	hasher := ZobristHasher{}
	for i := 0; i < 64; i++ {
		hasher.HashTableBlack[i] = r.Uint64()
		hasher.HashTableWhite[i] = r.Uint64()
	}
	return hasher
}

type ZobristHash uint64

func (h ZobristHasher) HashOfTileBlack(tile int) ZobristHash {
	return ZobristHash(h.HashTableBlack[tile])
}

func (h ZobristHasher) HashOfTileWhite(tile int) ZobristHash {
	return ZobristHash(h.HashTableWhite[tile])
}

func (h1 ZobristHash) Compose(h2 ZobristHash) ZobristHash {
	return h1 ^ h2
}

type SymetryLookUpTable struct {
	Symetry1 [64]int
	Symetry2 [64]int
	Symetry3 [64]int
	Symetry4 [64]int
	Symetry5 [64]int
	Symetry6 [64]int
	Symetry7 [64]int
	Symetry8 [64]int
}

func NewSymetryLookUpTable() SymetryLookUpTable {
	table := SymetryLookUpTable{}
	for i := 0; i < 64; i++ {
		row, col := i/8, i%8
		table.Symetry1[i] = (col    ) + 8 * (row    )
		table.Symetry2[i] = (col    ) + 8 * (7 - row)
		table.Symetry3[i] = (7 - col) + 8 * (row    )
		table.Symetry4[i] = (7 - col) + 8 * (7 - row)
		table.Symetry5[i] = (row    ) + 8 * (col    )
		table.Symetry6[i] = (row    ) + 8 * (7 - col)
		table.Symetry7[i] = (7 - row) + 8 * (col    )
		table.Symetry8[i] = (7 - row) + 8 * (7 - col)
	}
	return table
}

type SymetryHashList struct {
	hasher *ZobristHasher
	lookUp *SymetryLookUpTable

	Symetry1 ZobristHash
	Symetry2 ZobristHash
	Symetry3 ZobristHash
	Symetry4 ZobristHash
	Symetry5 ZobristHash
	Symetry6 ZobristHash
	Symetry7 ZobristHash
	Symetry8 ZobristHash
}

func NewSymetryHashList() SymetryHashList {
	hasher := NewZobristHasher(rand.New(rand.NewSource(0)))
	lookup := NewSymetryLookUpTable()
	return SymetryHashList{
		hasher: &hasher,
		lookUp: &lookup,
	}
}

func (h SymetryHashList) ApplyBitboardBlack(b Bitboard) SymetryHashList {
	var moveID int
	for b != 0 {
		moveID, b = b.Pop()
		h.Symetry1 ^= h.hasher.HashOfTileBlack(h.lookUp.Symetry1[moveID])
		h.Symetry2 ^= h.hasher.HashOfTileBlack(h.lookUp.Symetry2[moveID])
		h.Symetry3 ^= h.hasher.HashOfTileBlack(h.lookUp.Symetry3[moveID])
		h.Symetry4 ^= h.hasher.HashOfTileBlack(h.lookUp.Symetry4[moveID])
		h.Symetry5 ^= h.hasher.HashOfTileBlack(h.lookUp.Symetry5[moveID])
		h.Symetry6 ^= h.hasher.HashOfTileBlack(h.lookUp.Symetry6[moveID])
		h.Symetry7 ^= h.hasher.HashOfTileBlack(h.lookUp.Symetry7[moveID])
		h.Symetry8 ^= h.hasher.HashOfTileBlack(h.lookUp.Symetry8[moveID])
	}
	return h
}

func (h SymetryHashList) ApplyBitboardWhite(b Bitboard) SymetryHashList {
	var moveID int
	for b != 0 {
		moveID, b = b.Pop()
		h.Symetry1 ^= h.hasher.HashOfTileWhite(h.lookUp.Symetry1[moveID])
		h.Symetry2 ^= h.hasher.HashOfTileWhite(h.lookUp.Symetry2[moveID])
		h.Symetry3 ^= h.hasher.HashOfTileWhite(h.lookUp.Symetry3[moveID])
		h.Symetry4 ^= h.hasher.HashOfTileWhite(h.lookUp.Symetry4[moveID])
		h.Symetry5 ^= h.hasher.HashOfTileWhite(h.lookUp.Symetry5[moveID])
		h.Symetry6 ^= h.hasher.HashOfTileWhite(h.lookUp.Symetry6[moveID])
		h.Symetry7 ^= h.hasher.HashOfTileWhite(h.lookUp.Symetry7[moveID])
		h.Symetry8 ^= h.hasher.HashOfTileWhite(h.lookUp.Symetry8[moveID])
	}
	return h
}