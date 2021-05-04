package structs

type TransitionRow struct {
	P     int
	V     int
	A     int
	LV    int
	D     int
	TIPO  int
	L     int
	ES    int
	ARG   int
	CMD   int
	LD    int
	OPRD  int
	COND  int
	CAB   int
	CP    int
	EXP_R int
	R     int
	CP_R  int
}

var TransitionTable = [76]*TransitionRow{
	{P: 1}, // ESTADO 0
	{},     // ESTADO 1
	{V: 3}, // ESTADO 2
	{A: 17, ES: 18, CMD: 21, COND: 23, CAB: 24, R: 26}, // ESTADO 3
	{LV: 5, D: 6, TIPO: 7},                             // ESTADO 4
	{},                                                 // ESTADO 5
	{LV: 10, D: 6, TIPO: 7},                            // ESTADO 6
	{L: 13},                                            // ESTADO 7
	{},                                                 // ESTADO 8
	{},                                                 // ESTADO 9
	{},                                                 // ESTADO 10
	{L: 12},                                            // ESTADO 11
	{},                                                 // ESTADO 12
	{},                                                 // ESTADO 13
	{},                                                 // ESTADO 14
	{},                                                 // ESTADO 15
	{},                                                 // ESTADO 16
	{},                                                 // ESTADO 17
	{A: 54, ES: 18, CMD: 21, COND: 23, CAB: 24, R: 26}, // ESTADO 18
	{},        // ESTADO 19
	{ARG: 31}, // ESTADO 20
	{A: 55, ES: 18, CMD: 21, COND: 23, CAB: 24, R: 26}, // ESTADO 21
	{}, // ESTADO 22
	{A: 56, ES: 18, CMD: 21, COND: 23, CAB: 24, R: 26}, // ESTADO 23
	{ES: 58, CMD: 59, COND: 60, CAB: 24, CP: 57},       // ESTADO 24
	{}, // ESTADO 25
	{A: 62, ES: 18, CMD: 21, COND: 23, CAB: 24, R: 26}, // ESTADO 26
	{},                    // ESTADO 27
	{},                    // ESTADO 28
	{},                    // ESTADO 29
	{},                    // ESTADO 30
	{},                    // ESTADO 31
	{},                    // ESTADO 32
	{},                    // ESTADO 33
	{},                    // ESTADO 34
	{},                    // ESTADO 35
	{LD: 37, OPRD: 38},    // ESTADO 36
	{},                    // ESTADO 37
	{},                    // ESTADO 38
	{},                    // ESTADO 39
	{},                    // ESTADO 40
	{},                    // ESTADO 41
	{OPRD: 43},            // ESTADO 42
	{},                    // ESTADO 43
	{OPRD: 48, EXP_R: 45}, // ESTADO 44
	{},                    // ESTADO 45
	{},                    // ESTADO 46
	{},                    // ESTADO 47
	{},                    // ESTADO 48
	{OPRD: 50},            // ESTADO 49
	{},                    // ESTADO 50
	{OPRD: 48, EXP_R: 52}, // ESTADO 51
	{},                    // ESTADO 52
	{ES: 64, CMD: 65, COND: 66, CAB: 24, CP_R: 63}, // ESTADO 53
	{}, // ESTADO 54
	{}, // ESTADO 55
	{}, // ESTADO 56
	{}, // ESTADO 57
	{ES: 58, CMD: 59, COND: 60, CAB: 24, CP: 69}, // ESTADO 58
	{ES: 58, CMD: 59, COND: 60, CAB: 24, CP: 68}, // ESTADO 59
	{ES: 58, CMD: 59, COND: 60, CAB: 24, CP: 70}, // ESTADO 60
	{}, // ESTADO 61
	{}, // ESTADO 62
	{}, // ESTADO 63
	{ES: 64, CMD: 65, COND: 66, CAB: 24, CP_R: 71}, // ESTADO 64
	{ES: 64, CMD: 65, COND: 66, CAB: 24, CP_R: 72}, // ESTADO 65
	{ES: 64, CMD: 65, COND: 66, CAB: 24, CP_R: 73}, // ESTADO 66
	{}, // ESTADO 67
	{}, // ESTADO 68
	{}, // ESTADO 69
	{}, // ESTADO 70
	{}, // ESTADO 71
	{}, // ESTADO 72
	{}, // ESTADO 73
	{}, // ESTADO 74
	{}, // ESTADO 75

}
