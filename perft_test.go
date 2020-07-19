package main

import "testing"

type perft struct {
	states    int
	leaves    int
	passes    int
	gameOvers int
}

type perftResult struct {
	state    State
	depth    int
	expected perft
}

var perftResults = []perftResult{
	{NewState(0x000002001c000000, 0x0000001e00000000, Black), 5, perft{ 17205,  14914,   1,   1}},
	{NewState(0x0000000002000001, 0x0024141c1c1f0200, Black), 4, perft{  1361,   1158,   3,   3}},
	{NewState(0x0000080030000000, 0x0040203f0c040000, Black), 5, perft{ 42012,  37833,   3,   3}},
	{NewState(0x000010191a08080e, 0x0000000004040701, Black), 4, perft{   762,    655,   2,   0}},
	{NewState(0x203674ea58684000, 0x4c08081020002000, White), 4, perft{ 19195,  17105,   1,   0}},
	{NewState(0x0004181010000000, 0x0000000808000000, White), 5, perft{ 12143,  10427,   1,   1}},
	{NewState(0x022c28f630240000, 0x0000000808080000, White), 5, perft{ 44871,  40512,   1,   1}},
	{NewState(0x00000040203e0000, 0x0020181818000800, White), 4, perft{ 10892,   9844,   1,   0}},
	{NewState(0x07263c383030000a, 0x000900478c0c1c00, White), 5, perft{114387, 101937,   2,   0}},
	{NewState(0x000002021a080000, 0x0000001c00000000, White), 4, perft{  2003,   1692,   1,   1}},
	{NewState(0x414070f04cfe2c08, 0x0080840830000100, White), 5, perft{ 92927,  85309,   1,   0}},
	{NewState(0x000808f808040000, 0x0000000010080000, White), 5, perft{ 14653,  12768,   3,   3}},
	{NewState(0x0000000860000000, 0x0010101018100800, Black), 5, perft{ 23492,  20666,   3,   3}},
	{NewState(0x0000000000080800, 0x0000001c1c140000, Black), 5, perft{ 32887,  29152,   1,   1}},
	{NewState(0x0008141010000000, 0x0010086c080c0a00, Black), 5, perft{137715, 125597,   3,   3}},
	{NewState(0x103c3f34181a0200, 0x4800000800040007, White), 5, perft{ 42270,  38326,   2,   2}},
	{NewState(0x4040500810080000, 0x8000201008040000, Black), 5, perft{  6572,   5584,   1,   1}},
	{NewState(0x0000100878080000, 0x0000201000100800, White), 5, perft{ 28860,  25381,   1,   1}},
	{NewState(0x0000000000060810, 0x0000001c1c180008, White), 5, perft{ 24692,  21998,   2,   0}},
	{NewState(0x0000001432000000, 0x00080808081c0000, Black), 4, perft{ 10012,   8998,   1,   0}},
	{NewState(0x0000003810001000, 0x00000000083c0000, Black), 4, perft{  3422,   3025,   1,   1}},
	{NewState(0x103e063878000000, 0x04407844041c0000, White), 5, perft{292354, 267585,   2,   0}},
	{NewState(0x000000001c0c0000, 0x1020787820101810, Black), 5, perft{135940, 125392,  23,  23}},
	{NewState(0x00007c0810000000, 0x0000001008000000, White), 5, perft{ 18539,  16253,   3,   3}},
	{NewState(0x1140645000000000, 0x0008882c7eb00000, White), 5, perft{162203, 147723,   1,   0}},
	{NewState(0x0000003800040200, 0x000000003e000000, Black), 5, perft{ 30214,  26546,   3,   3}},
	{NewState(0x0020602830200000, 0x0000005008000000, White), 5, perft{ 36727,  32737,   1,   1}},
	{NewState(0x0000101820400000, 0x0000000018000000, White), 5, perft{  8495,   7250,   1,   0}},
	{NewState(0x0000101010040200, 0x0000080808080800, Black), 5, perft{ 30103,  26445,   3,   3}},
	{NewState(0x0008102020400000, 0x0000041818000000, Black), 4, perft{  3264,   2873,   1,   1}},
	{NewState(0x00001c0d04860000, 0x007460707b704000, White), 5, perft{187118, 167499,  21,   0}},
	{NewState(0x040c546472000000, 0x0090081808080800, White), 5, perft{198221, 182379,   3,   1}},
	{NewState(0x0000402000042000, 0x00080e5d7fb05010, Black), 5, perft{158851, 146071,   2,   2}},
	{NewState(0x0018101810000000, 0x0e0408040a000000, Black), 5, perft{ 25210,  22119,   1,   0}},
	{NewState(0x0000020418380000, 0x0000001800000000, White), 5, perft{ 15823,  13835,   1,   0}},
	{NewState(0x0000783818080000, 0x0002040000100000, White), 5, perft{ 36577,  32822,   1,   1}},
	{NewState(0x0040203800000000, 0x000000001e100000, Black), 5, perft{ 12439,  10903,   1,   1}},
	{NewState(0x0000001808080000, 0x00003c0030202000, Black), 5, perft{ 60982,  55081,   4,   4}},
	{NewState(0x00202018b82c0200, 0x000000e000000000, White), 5, perft{ 20963,  18780,   2,   2}},
	{NewState(0x0000104030100000, 0x0000081808000000, White), 5, perft{ 22019,  19466,   3,   3}},
	{NewState(0x0000206028000000, 0x1030501c140c0400, Black), 5, perft{180402, 166162,   2,   2}},
	{NewState(0x000808f818102400, 0x00000000000e0000, White), 5, perft{ 36118,  32393,   5,   5}},
	{NewState(0x000000001c3c0200, 0x0000203c20000000, Black), 5, perft{ 91971,  83240,   1,   1}},
	{NewState(0x00000038000e0000, 0x0000000038100404, Black), 5, perft{ 30101,  26346,   1,   0}},
	{NewState(0x0044283c100c1400, 0x8000000008102002, Black), 5, perft{  6706,   5728,   2,   2}},
	{NewState(0x0000000808162808, 0x0000001010080400, Black), 4, perft{  3929,   3474,   1,   1}},
	{NewState(0x0000003e10000000, 0x000008000a080800, White), 5, perft{ 17178,  14917,   1,   1}},
	{NewState(0x00001c1818080000, 0x001c020400000000, White), 5, perft{ 56273,  49908,   1,   0}},
	{NewState(0x0000220418080000, 0x0000003800000000, White), 5, perft{ 23998,  20948,   1,   0}},
	{NewState(0x0004083010000000, 0x0000000808000000, White), 5, perft{  8476,   7233,   1,   1}},
	{NewState(0x0000000006000000, 0x0000001c181c0000, Black), 5, perft{ 33592,  29822,   1,   0}},
	{NewState(0x0000003048ff5220, 0x0000048810000010, White), 5, perft{ 36380,  32098,   1,   0}},
	{NewState(0x0040200818080000, 0x0000181020000000, Black), 5, perft{ 20016,  17435,   1,   1}},
	{NewState(0x00000400100e0000, 0x0000081c08000000, Black), 5, perft{ 15696,  13666,   1,   1}},
	{NewState(0x007c7c5cec0e0800, 0x1000002010901004, Black), 5, perft{ 27771,  23925,   1,   0}},
	{NewState(0x0000701810141000, 0x0000000408000000, White), 5, perft{ 15497,  13688,   1,   1}},
	{NewState(0x0008000afe3c6080, 0x00003e3000000000, Black), 5, perft{ 97610,  87426,   6,   0}},
	{NewState(0x000004181c000000, 0x0008380000000000, Black), 5, perft{ 27509,  24318,   1,   0}},
	{NewState(0x000000101c181000, 0x0020180c02040800, White), 5, perft{ 80876,  71742,   1,   0}},
	{NewState(0x0000020018080000, 0x0000001e04040000, Black), 5, perft{ 48935,  43452,   1,   0}},
	{NewState(0x0008180808000000, 0x0000001036283000, Black), 5, perft{ 78247,  70795,   1,   1}},
	{NewState(0x800000080c102000, 0x00402010100c0c08, White), 5, perft{ 31935,  28160,   1,   0}},
	{NewState(0x0020303814000000, 0x0000080008040000, White), 5, perft{ 26806,  23903,   1,   1}},
	{NewState(0x0000700810200000, 0x0000001028000000, White), 5, perft{ 15461,  13530,   1,   1}},
	{NewState(0x040600091d008000, 0x00080e1602000000, Black), 5, perft{ 94404,  83552,   1,   0}},
	{NewState(0x0040321c1c000100, 0x0000000000070604, White), 5, perft{ 12790,  11412,   1,   1}},
	{NewState(0x0080000c00002000, 0x000060301e3c0000, Black), 5, perft{ 62612,  56541,   6,   0}},
	{NewState(0x000020080c102000, 0x0000101010000000, White), 5, perft{ 23966,  20918,   1,   1}},
	{NewState(0x0000000e04080101, 0x000408103a060608, White), 5, perft{ 35697,  30985,   2,   0}},
	{NewState(0x0008082818240000, 0x0000001020400000, White), 5, perft{ 31948,  28347,   1,   0}},
	{NewState(0x101010080c002000, 0x00600050303c0000, White), 5, perft{ 87250,  77747,  12,   0}},
	{NewState(0x0040307028180800, 0x0000000c10000004, White), 5, perft{ 66816,  60701,   5,   2}},
	{NewState(0x0000101818000200, 0x0000000006040400, Black), 5, perft{  5583,   4765,   6,   2}},
	{NewState(0x0040405818080000, 0x0000380000000000, White), 4, perft{  3510,   3028,   1,   1}},
	{NewState(0x001010181f10103a, 0x00080d662068ac00, White), 5, perft{125816, 110853,   3,   0}},
	{NewState(0x00002014040e0000, 0x0000040838000000, Black), 5, perft{ 50003,  44531,   1,   1}},
	{NewState(0x0000100810101004, 0x0000201028000800, White), 5, perft{ 11443,   9827,   1,   1}},
	{NewState(0x00081a3e7868fc20, 0x0101010080140044, White), 4, perft{  4714,   4093,   4,   0}},
	{NewState(0x0000201010040200, 0x0000000c0c100000, Black), 5, perft{ 11572,   9889,   1,   1}},
	{NewState(0x0000101818000a00, 0x0000000006040404, Black), 3, perft{   125,     97,   1,   0}},
	{NewState(0x8000109800000000, 0x206020207c281000, Black), 5, perft{ 72956,  66089,   4,   0}},
	{NewState(0x0000501810204000, 0x0020202028002000, White), 5, perft{ 99136,  88213,   1,   0}},
	{NewState(0x1c00001818380000, 0x0010380000000000, Black), 5, perft{ 17565,  15524,   1,   0}},
	{NewState(0x1031564cdc949010, 0x0084083020400000, White), 5, perft{200274, 184257,   3,   3}},
	{NewState(0x0004181010000000, 0x0000000808000000, White), 5, perft{ 12143,  10427,   1,   1}},
	{NewState(0x0040d07e7f060400, 0xa838280000000000, Black), 5, perft{ 28583,  24789,   3,   0}},
	{NewState(0x000004053e040200, 0x0000001800180001, White), 5, perft{ 75954,  68760,   4,   2}},
	{NewState(0x004002bc1c000000, 0x0010204080000000, Black), 4, perft{   308,    260,   1,   1}},
	{NewState(0x1020603014101000, 0x0000180c08000000, White), 5, perft{133814, 122223,   6,   6}},
	{NewState(0x0000000814000400, 0x0000001008040200, Black), 5, perft{  6890,   5842,   1,   0}},
	{NewState(0x002020383c3af848, 0x0000004000000000, White), 5, perft{  7346,   6557,   4,   2}},
	{NewState(0x00000038302c0200, 0x000000000c000000, White), 5, perft{ 31282,  27848,  34,  34}},
	{NewState(0x0000040400000000, 0x0000011a1c1c0000, Black), 3, perft{   338,    293,   1,   0}},
	{NewState(0x3010181810000102, 0x040822070d0e1400, White), 5, perft{ 20108,  17236,   1,   0}},
	{NewState(0x0000000000162004, 0x0040201c1c080800, White), 4, perft{  3129,   2731,   1,   1}},
	{NewState(0x0408101010000000, 0x00020c0808000000, Black), 5, perft{ 24263,  21166,   1,   1}},
	{NewState(0x0004181010000000, 0x0000000808000000, White), 5, perft{ 12143,  10427,   1,   1}},
	{NewState(0x0020202818080000, 0x0000001020000000, White), 5, perft{ 15433,  13487,   3,   3}},
	{NewState(0x008002f400000002, 0x000048085c3c1c10, White), 5, perft{ 28392,  24925,   1,   1}},
	{NewState(0x0000502018080000, 0x0000001c00000000, White), 5, perft{ 19129,  16727,   3,   3}},
}

func TestPerft(t *testing.T) {
	for _, test := range perftResults {
		result := runperftTest(test.state, test.depth)
		if result != test.expected {
			t.Fatal("Expected result not given")
		}
	}
}

func runperftTest(state State, depth int) perft {
	var p perft
	p.depthFirstSearch(state, depth)
	return p
}

func (p *perft) depthFirstSearch(state State, depth int) {
	p.states++
	
	moves := state.Moves()
	if len(moves) == 0 {
		p.leaves++
		p.gameOvers++
		return
	}
	
	if depth == 0 {
		p.leaves++
		return
	}

	if moves[0] == MovePass {
		p.passes++
	}

	for _, move := range moves {
		next := state.MakeMove(move)
		p.depthFirstSearch(next, depth-1)
	}
}