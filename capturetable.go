package main

type CaptureTable struct {
	discsPlayer   Bitboard
	discsOpponent Bitboard
	empty         Bitboard

	capturesN  Bitboard
	capturesE  Bitboard
	capturesS  Bitboard
	capturesW  Bitboard
	capturesNE Bitboard
	capturesSE Bitboard
	capturesSW Bitboard
	capturesNW Bitboard
}

func NewCaptureTable(discsPlayer, discsOpponent Bitboard) CaptureTable {
	ct := CaptureTable{
		discsPlayer:   discsPlayer,
		discsOpponent: discsOpponent,
		empty:       ^(discsPlayer | discsOpponent),

		capturesN:  discsPlayer.ShiftN()  & discsOpponent,
		capturesE:  discsPlayer.ShiftE()  & discsOpponent,
		capturesS:  discsPlayer.ShiftS()  & discsOpponent,
		capturesW:  discsPlayer.ShiftW()  & discsOpponent,
		capturesNE: discsPlayer.ShiftNE() & discsOpponent,
		capturesSE: discsPlayer.ShiftSE() & discsOpponent,
		capturesSW: discsPlayer.ShiftSW() & discsOpponent,
		capturesNW: discsPlayer.ShiftNW() & discsOpponent,
	}

	for i := 0; i < 5; i++ {
		ct.capturesN  |= ct.capturesN.ShiftN()   & discsOpponent
		ct.capturesE  |= ct.capturesE.ShiftE()   & discsOpponent
		ct.capturesS  |= ct.capturesS.ShiftS()   & discsOpponent
		ct.capturesW  |= ct.capturesW.ShiftW()   & discsOpponent
		ct.capturesNE |= ct.capturesNE.ShiftNE() & discsOpponent
		ct.capturesSE |= ct.capturesSE.ShiftSE() & discsOpponent
		ct.capturesSW |= ct.capturesSW.ShiftSW() & discsOpponent
		ct.capturesNW |= ct.capturesNW.ShiftNW() & discsOpponent
	}

	return ct
}

func (ct CaptureTable) Moves() Bitboard {
	return (ct.capturesN.ShiftN()   |
	        ct.capturesE.ShiftE()   |
	        ct.capturesS.ShiftS()   |
	        ct.capturesW.ShiftW()   |
	        ct.capturesNE.ShiftNE() |
	        ct.capturesSE.ShiftSE() |
	        ct.capturesSW.ShiftSW() |
	        ct.capturesNW.ShiftNW() ) & ct.empty
}

func (ct CaptureTable) Captures(disc Bitboard) Bitboard {
	var (
		capturesN  = disc.ShiftN()  & ct.capturesS
		capturesE  = disc.ShiftE()  & ct.capturesW
		capturesS  = disc.ShiftS()  & ct.capturesN
		capturesW  = disc.ShiftW()  & ct.capturesE
		capturesNE = disc.ShiftNE() & ct.capturesSW
		capturesSE = disc.ShiftSE() & ct.capturesNW
		capturesSW = disc.ShiftSW() & ct.capturesNE
		capturesNW = disc.ShiftNW() & ct.capturesSE
	)

	for i := 0; i < 5; i++ {
		capturesN  |= capturesN.ShiftN()   & ct.capturesS
		capturesE  |= capturesE.ShiftE()   & ct.capturesW
		capturesS  |= capturesS.ShiftS()   & ct.capturesN
		capturesW  |= capturesW.ShiftW()   & ct.capturesE
		capturesNE |= capturesNE.ShiftNE() & ct.capturesSW
		capturesSE |= capturesSE.ShiftSE() & ct.capturesNW
		capturesSW |= capturesSW.ShiftSW() & ct.capturesNE
		capturesNW |= capturesNW.ShiftNW() & ct.capturesSE
	}

	return capturesN  | capturesE  | capturesS  | capturesW  |
	       capturesNE | capturesSE | capturesSW | capturesNW
}