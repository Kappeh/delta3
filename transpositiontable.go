package main

import "sync"

type TranspositionInfo struct {
	SearchDepth int
	Score       int
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

func (tt *TranspositionTable) Set(h ZobristHash, i TranspositionInfo) {
	tt.lock.Lock()
	tt.table[h] = i
	tt.lock.Unlock()
}

func (tt *TranspositionTable) Remove(h ZobristHash) {
	tt.lock.Lock()
	delete(tt.table, h)
	tt.lock.Unlock()
}

func (tt *TranspositionTable) Probe(h ZobristHash) (TranspositionInfo, bool) {
	tt.lock.RLock()
	i, ok := tt.table[h]
	tt.lock.RUnlock()
	return i, ok
}