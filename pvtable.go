package main

type PVTable struct {
	MaxDepth   int
	MoveCounts []int
	Table      []Move
}

type PVTableCursor struct {
	PVTable
	Depth   int
	PVIndex int
}

func NewPVTable(depth int) PVTable {
	size := (depth * depth + depth) >> 1
	return PVTable{
		MaxDepth:   depth,
		MoveCounts: make([]int, depth),
		Table:      make([]Move, size),
	}
}

func (pv PVTable) Cursor() PVTableCursor {
	return PVTableCursor{
		PVTable: pv,
		Depth:   0,
		PVIndex: 0,
	}
}

func (pv PVTable) Line() []Move {
	return pv.Table[:pv.MoveCounts[0]]
}

func (c PVTableCursor) NextDepth() PVTableCursor {
	c.MoveCounts[c.Depth + 1] = 0
	
	return PVTableCursor{
		PVTable: c.PVTable,
		Depth:   c.Depth + 1,
		PVIndex: c.PVIndex + c.MaxDepth - c.Depth,
	}
}

func (c PVTableCursor) Move() Move {
	return c.Table[c.PVIndex]
}

func (c PVTableCursor) Set(move Move) {
	c.Table[c.PVIndex] = move
	
	indexDest := c.PVIndex + 1
	indexSrc  := c.PVIndex + c.MaxDepth - c.Depth
	copySize  := c.MaxDepth - c.Depth - 1

	copy(
		c.Table[indexDest : indexDest + copySize],
		c.Table[indexSrc  : indexSrc  + copySize],
	)

	if c.Depth == c.MaxDepth - 1 {
		c.MoveCounts[c.Depth] = 1
	} else {
		c.MoveCounts[c.Depth] = c.MoveCounts[c.Depth + 1] + 1
	}
}