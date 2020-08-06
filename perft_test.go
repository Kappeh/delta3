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
	{NewState(0x000000069e1e1273, 0x827efa7860216c04, White), 5, perft{ 36006,  30780,   9,   0}},
	{NewState(0x0000003814000000, 0x0000000008040000, White), 5, perft{  7376,   6289,   0,   0}},
	{NewState(0x0000162c38303000, 0x00442812044c0000, White), 4, perft{ 50942,  47959,   0,   0}},
	{NewState(0x0000007a271f0804, 0x00101e0418e03020, Black), 3, perft{  1670,   1536,   0,   0}},
	{NewState(0x000002041c000000, 0x0000001800000000, White), 2, perft{    42,     35,   0,   0}},
	{NewState(0x000002083f244000, 0x00040414001a1000, White), 2, perft{   241,    226,   0,   0}},
	{NewState(0x0efe42c0c0c1a000, 0x41013d3f3f3e5cff, Black), 5, perft{   104,     46,   9,   6}},
	{NewState(0x000040a0c1820202, 0x0050301c3c140405, Black), 1, perft{    12,     11,   0,   0}},
	{NewState(0xfe7c380004442000, 0x0083c7fdfbbb9fef, Black), 1, perft{     3,      2,   0,   0}},
	{NewState(0x0000000810080000, 0x0000001008040000, Black), 5, perft{  4394,   3710,   0,   0}},
	{NewState(0x000818084c081402, 0x0000001430500000, White), 3, perft{   867,    782,   0,   0}},
	{NewState(0x00000400043e1f10, 0x0000805f7840400c, Black), 4, perft{ 17091,  15440,   0,   0}},
	{NewState(0x3c282e1020000000, 0x00c05088dcba1038, Black), 2, perft{    92,     80,   0,   0}},
	{NewState(0x0400041c3c274484, 0x20fffbe1c0d82a29, White), 4, perft{  6751,   6042,   0,   0}},
	{NewState(0x30783008040e0050, 0x080488d0b8f0fe00, White), 5, perft{155186, 138395,   0,   0}},
	{NewState(0x9051231d0b572609, 0x040a1c2270081010, White), 4, perft{ 17953,  16379,   0,   0}},
	{NewState(0x0004080814182000, 0x0020301008040000, Black), 3, perft{   914,    820,   0,   0}},
	{NewState(0x101810e476002030, 0x0000611a887c0040, Black), 4, perft{ 28431,  25958,   0,   0}},
	{NewState(0x00002001828b9204, 0x000014583c740900, White), 4, perft{ 13190,  12039,   0,   0}},
	{NewState(0x0000101800080000, 0x0000000038040000, Black), 2, perft{    41,     34,   0,   0}},
	{NewState(0x0f0f09bf118d0307, 0x70f0f6402e72fc10, White), 2, perft{    15,     12,   0,   0}},
	{NewState(0x009050d2262cc820, 0x00080428d8503010, White), 5, perft{371715, 343783,   0,   0}},
	{NewState(0x40603c2800000000, 0x080800143a000000, White), 2, perft{   107,     97,   0,   0}},
	{NewState(0x0088402290502048, 0x000038192e2f1c24, Black), 1, perft{    13,     12,   0,   0}},
	{NewState(0x0000105e26122000, 0x0000602018040000, Black), 3, perft{  1358,   1210,   0,   0}},
	{NewState(0x0000081828180800, 0x0000000410200000, White), 1, perft{    11,     10,   0,   0}},
	{NewState(0x0000081818380008, 0x20e0a0e0e0c42a10, Black), 3, perft{   645,    574,   0,   0}},
	{NewState(0x3c3842e607070e04, 0x83c7bd19d8f8f1fb, Black), 2, perft{     5,      2,   0,   0}},
	{NewState(0x00403eb06134f818, 0x0000000e9e090281, White), 1, perft{    18,     17,   0,   0}},
	{NewState(0x0000000810000000, 0x0000001008000000, Black), 5, perft{  1713,   1396,   0,   0}},
	{NewState(0x1008141210000000, 0x0002080808000000, White), 1, perft{     8,      7,   0,   0}},
	{NewState(0x3ebf2f019b170f0f, 0x000050fe6468f030, Black), 4, perft{   280,    199,   0,   0}},
	{NewState(0x00000070484274c8, 0x0112548816380804, Black), 1, perft{    13,     12,   0,   0}},
	{NewState(0x001008046a04a800, 0x0048203814224080, White), 4, perft{ 21926,  19770,   0,   0}},
	{NewState(0x07ef47fac0ad9000, 0xf01028043e522b1d, White), 3, perft{   493,    419,   0,   0}},
	{NewState(0x022428393b7b7002, 0x08c087c684040601, White), 1, perft{    14,     13,   0,   0}},
	{NewState(0x000000080c000000, 0x0000001010100000, Black), 4, perft{   847,    698,   0,   0}},
	{NewState(0x0000101800000000, 0x0000000038000000, Black), 5, perft{  5510,   4663,   0,   0}},
	{NewState(0x0000000808080000, 0x0000001010100000, Black), 4, perft{   946,    788,   0,   0}},
	{NewState(0x000020101c140400, 0x0000040820000000, White), 1, perft{     8,      7,   0,   0}},
	{NewState(0x0004003810080000, 0x00082e040c060000, Black), 3, perft{  1790,   1635,   0,   0}},
	{NewState(0x1c476305f0c02022, 0x62381c3a0f3c9e48, Black), 1, perft{    12,     11,   0,   0}},
	{NewState(0x007c0f1537033c0c, 0xfd82f0e8c8fcc2f1, White), 4, perft{   123,     57,   4,   0}},
	{NewState(0x00402418040c3221, 0x00200264b8504040, Black), 3, perft{  1949,   1788,   0,   0}},
	{NewState(0x0000002850800400, 0x00003e10287c0800, Black), 2, perft{   117,    106,   0,   0}},
	{NewState(0x0401b2bdae9ea2f8, 0x597e4c4251604000, White), 2, perft{    83,     73,   0,   0}},
	{NewState(0xfefe5c6260f8f800, 0x0001a39d9e0606ff, Black), 5, perft{    45,     12,  11,  12}},
	{NewState(0x00083878f868400c, 0x2356840404143e70, Black), 2, perft{   170,    154,   0,   0}},
	{NewState(0x02020da814a8dd84, 0x9c7c725669562268, White), 1, perft{     6,      5,   0,   0}},
	{NewState(0x080407003e605008, 0x02083838400e0000, White), 2, perft{   153,    141,   0,   0}},
	{NewState(0x00403010101c0000, 0x0000002c28000000, White), 2, perft{    96,     86,   0,   0}},
	{NewState(0x03060e2c2c381800, 0xdcd9f1d3d3c6e63e, Black), 1, perft{     8,      7,   0,   0}},
	{NewState(0x9f01775cf9136f7c, 0x2054880306ec1001, Black), 3, perft{    91,     65,   0,   0}},
	{NewState(0x00040818171f0b04, 0x0080f2666860100a, White), 5, perft{156269, 138615,   0,   0}},
	{NewState(0x11824d1c001d2e49, 0x44283262fe021000, White), 4, perft{ 22394,  20576,   0,   0}},
	{NewState(0x00a0102838380800, 0x8040205040448000, White), 2, perft{   127,    115,   0,   0}},
	{NewState(0x000010f810183040, 0x00f4240408040200, White), 5, perft{115010, 103787,   0,   0}},
	{NewState(0x000d18181c023040, 0x01020406e07c0800, Black), 3, perft{  2563,   2358,   0,   0}},
	{NewState(0x81dfa30f8903078f, 0x22205cf0767cf830, Black), 5, perft{   424,    229,   0,   0}},
	{NewState(0x00002040081e1000, 0x0000083830600000, Black), 1, perft{     8,      7,   0,   0}},
	{NewState(0x000000003c1e2000, 0x0000207c80000000, Black), 1, perft{     9,      8,   0,   0}},
	{NewState(0x804d6f5f6a752345, 0x58b01020140a1c18, White), 3, perft{   574,    504,   0,   0}},
	{NewState(0x002022041c040000, 0x0000007800080000, White), 1, perft{     9,      8,   0,   0}},
	{NewState(0x070f1f162efe1438, 0xf8f0e0e950000b00, White), 3, perft{   171,    136,   0,   0}},
	{NewState(0x1098cce063e68605, 0x0821131f1c195810, Black), 2, perft{   120,    107,   0,   0}},
	{NewState(0x04462f576e54fc00, 0x109850a81128007e, White), 3, perft{   647,    577,   0,   0}},
	{NewState(0x0000203008102000, 0x0010180834020800, White), 2, perft{    71,     64,   0,   0}},
	{NewState(0x0000183438dc0044, 0x203a66cac523fd28, Black), 3, perft{  1482,   1369,   0,   0}},
	{NewState(0x0000381018080000, 0x0020000c04000000, White), 5, perft{ 63205,  56346,   0,   0}},
	{NewState(0x0000000810000000, 0x0000001008000000, Black), 4, perft{   317,    244,   0,   0}},
	{NewState(0x1a1efe787b303a01, 0x2001010604ce043c, White), 5, perft{ 78215,  69036,   0,   0}},
	{NewState(0x0000003c08040000, 0x0000080010200000, White), 2, perft{    55,     47,   0,   0}},
	{NewState(0x2000402030488c02, 0x08103c1848074200, Black), 5, perft{408517, 379266,   0,   0}},
	{NewState(0x080c1838e0000400, 0x905020061c181000, White), 1, perft{     9,      8,   0,   0}},
	{NewState(0x0000201008182800, 0x0000040810000400, White), 2, perft{    38,     32,   0,   0}},
	{NewState(0xe0c06ed66a58b040, 0x0828902915a70f07, White), 2, perft{    84,     73,   0,   0}},
	{NewState(0x00180cdc08000000, 0x20202021321c0c08, Black), 3, perft{  1026,    925,   0,   0}},
	{NewState(0x0101610710280000, 0x00181e180a122000, Black), 1, perft{    14,     13,   0,   0}},
	{NewState(0x4020105c6c548000, 0x02040c0010007000, White), 2, perft{    92,     83,   0,   0}},
	{NewState(0xff383c2ea2120306, 0x00c7c3d15de5fc10, Black), 5, perft{   332,    164,   3,   0}},
	{NewState(0x0101050d1d1d0d04, 0xfefefaf2e2e232fb, Black), 4, perft{     4,      1,   1,   1}},
	{NewState(0x3c322c2e362e4cfe, 0xc0c9d3d1c9d1b301, White), 5, perft{    15,      4,   4,   4}},
	{NewState(0x01203c2a167e2a02, 0x105f425408000000, White), 1, perft{    12,     11,   0,   0}},
	{NewState(0x040c04001c080000, 0x0040301f00742700, Black), 1, perft{    17,     16,   0,   0}},
	{NewState(0x0000003810000000, 0x0000000008000000, White), 3, perft{    79,     61,   0,   0}},
	{NewState(0x020c181808ba8c88, 0x00806022f4402010, White), 4, perft{ 15152,  13739,   0,   0}},
	{NewState(0x0010087420400000, 0x0000440818000000, White), 3, perft{   699,    625,   0,   0}},
	{NewState(0x4444f8f8dc7cc281, 0x18a8040722030400, White), 2, perft{   134,    122,   0,   0}},
	{NewState(0x80f5b87430301000, 0x2000070b4e0e0000, Black), 1, perft{    12,     11,   0,   0}},
	{NewState(0x013b3509f12b78e6, 0xfec4caf60e140201, White), 2, perft{    19,     12,   0,   0}},
	{NewState(0x0000182814204000, 0x0070201408500000, White), 5, perft{125158, 112185,   0,   0}},
	{NewState(0x007020100e0e0100, 0x220c182810100000, Black), 3, perft{  1264,   1136,   0,   0}},
	{NewState(0x0000101810080000, 0x0000000008040000, White), 5, perft{  7376,   6289,   0,   0}},
	{NewState(0xfcc373a996ac58f8, 0x023c0c5468532300, White), 3, perft{   208,    168,   0,   0}},
	{NewState(0x10000409123c3448, 0xe07cf8542c000010, White), 3, perft{  1167,   1056,   0,   0}},
	{NewState(0x0000100070040000, 0x0000081808080000, Black), 1, perft{     7,      6,   0,   0}},
	{NewState(0x0000043830200000, 0x0000080008000000, White), 5, perft{ 12987,  11376,   0,   0}},
	{NewState(0x000000081c000000, 0x00000010000f0800, Black), 5, perft{ 32178,  28502,   0,   0}},
	{NewState(0xfede02040a460202, 0x0121fdfbf5b9f9fc, Black), 1, perft{     2,      1,   0,   0}},
	{NewState(0x0000241c2e060810, 0x0000402011190500, Black), 4, perft{  9948,   9060,   0,   0}},
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

	if moves[0].ID == MoveIDPass {
		p.passes++
	}

	for _, move := range moves {
		next := state.MakeMove(move)
		p.depthFirstSearch(next, depth-1)
	}
}