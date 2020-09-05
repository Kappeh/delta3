package main

import "sync"

type TranspositionInfo struct {
	NodeInfo
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
			NodeInfo: v.NodeInfo,
			Ancient:  true,
		}
	}
	tt.lock.Unlock()
}

func (tt *TranspositionTable) MarkNotAncient(h ZobristHash) {
	tt.lock.Lock()
	v, ok := tt.table[h]
	if ok {
		tt.table[h] = TranspositionInfo{
			NodeInfo: v.NodeInfo,
			Ancient:  false,
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

func (tt *TranspositionTable) ProbeSymetry(s SymetryHashList) (ZobristHash, bool) {
	var hash ZobristHash
	var ok   bool

	tt.lock.RLock()
	if _, ok = tt.table[s.Symetry1]; ok {
		hash = s.Symetry1
		goto Return
	}
	if _, ok = tt.table[s.Symetry2]; ok {
		hash = s.Symetry2
		goto Return
	}
	if _, ok = tt.table[s.Symetry3]; ok {
		hash = s.Symetry3
		goto Return
	}
	if _, ok = tt.table[s.Symetry4]; ok {
		hash = s.Symetry4
		goto Return
	}
	if _, ok = tt.table[s.Symetry5]; ok {
		hash = s.Symetry5
		goto Return
	}
	if _, ok = tt.table[s.Symetry6]; ok {
		hash = s.Symetry6
		goto Return
	}
	if _, ok = tt.table[s.Symetry7]; ok {
		hash = s.Symetry7
		goto Return
	}
	if _, ok = tt.table[s.Symetry8]; ok {
		hash = s.Symetry8
		goto Return
	}
	
Return:
	tt.lock.RUnlock()
	return hash, ok
}