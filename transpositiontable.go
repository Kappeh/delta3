package main

import "sync"

type TranspositionInfo struct {
	// Score is an alpha value when a  lower bound
	// Score is a  beta  value when an upper bound
	Score       Boundary
	SearchDepth int
	PVMoveID    int
	Ancient     bool
}

type TranspositionTable struct {
	lock  *sync.RWMutex
	table map[ZobristHash]TranspositionInfo
}

func NewTranspositionTable() *TranspositionTable {
	return &TranspositionTable{
		lock:  &sync.RWMutex{},
		table: make(map[ZobristHash]TranspositionInfo),
	}
}

func (tt *TranspositionTable) Clear() {
	tt.lock.Lock()
	tt.table = make(map[ZobristHash]TranspositionInfo)
	tt.lock.Unlock()
}

func (tt *TranspositionTable) MarkAllAncient() {
	tt.lock.Lock()
	for k, v := range tt.table {
		tt.table[k] = TranspositionInfo{
			Score:       v.Score,
			SearchDepth: v.SearchDepth,
			Ancient:     true,
		}
	}
	tt.lock.Unlock()
}

func (tt *TranspositionTable) MarkNotAncient(h ZobristHash) {
	tt.lock.Lock()
	v, ok := tt.table[h]
	if ok {
		tt.table[h] = TranspositionInfo{
			Score:       v.Score,
			SearchDepth: v.SearchDepth,
			Ancient:     false,
		}
	}
	tt.lock.Unlock()
}

func (tt *TranspositionTable) Set(h ZobristHash, i TranspositionInfo) {
	tt.lock.RLock()
	v, ok := tt.table[h]
	tt.lock.RUnlock()
	
	if ok && v.Ancient == false {
		return
	}

	tt.lock.Lock()
	tt.table[h] = i
	tt.lock.Unlock()
}

func (tt *TranspositionTable) Probe(h ZobristHash) (TranspositionInfo, bool) {
	tt.lock.RLock()
	i, ok := tt.table[h]
	tt.lock.RUnlock()
	return i, ok
}

func (tt *TranspositionTable) ProbeSymetry(s SymetryHashList) (TranspositionInfo, bool) {
	var v  TranspositionInfo
	var ok bool

	tt.lock.RLock()
	if v, ok = tt.table[s.Symetry1]; ok {
		goto Return
	}
	if v, ok = tt.table[s.Symetry2]; ok {
		goto Return
	}
	if v, ok = tt.table[s.Symetry3]; ok {
		goto Return
	}
	if v, ok = tt.table[s.Symetry4]; ok {
		goto Return
	}
	if v, ok = tt.table[s.Symetry5]; ok {
		goto Return
	}
	if v, ok = tt.table[s.Symetry6]; ok {
		goto Return
	}
	if v, ok = tt.table[s.Symetry7]; ok {
		goto Return
	}
	if v, ok = tt.table[s.Symetry8]; ok {
		goto Return
	}
Return:
	tt.lock.RUnlock()
	return v, ok
}