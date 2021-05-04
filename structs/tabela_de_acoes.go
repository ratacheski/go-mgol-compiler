package structs

type ActionRow struct {
	INICIO     string
	VAR_INICIO string
	VAR_FIM    string
	PT_V       string
	ID         string
	VIR        string
	INTEIRO    string
	REAL       string
	LIT        string
	LEIA       string
	ESCREVA    string
	LITERAL    string
	NUM        string
	RCB        string
	OPM        string
	SE         string
	AB_P       string
	FC_P       string
	ENTAO      string
	OPR        string
	FIM_SE     string
	FACE_ATE   string
	FIM_FACA   string
	FIM        string
	EOF        string
}

var ActionTable = [76]*ActionRow{
	{INICIO: "S2"},     //ESTADO 0
	{EOF: "ACC"},       //ESTADO 1
	{VAR_INICIO: "S4"}, //ESTADO 2
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FACE_ATE: "S27", FIM: "S28"}, //ESTADO 3
	{VAR_FIM: "S9", INTEIRO: "S14", REAL: "S15", LIT: "S16"},                         //ESTADO 4
	{ID: "R3", LEIA: "R3", ESCREVA: "R3", SE: "R3", FACE_ATE: "R3", FIM: "R3"},       //ESTADO 5
	{VAR_FIM: "S9", INTEIRO: "S14", REAL: "S15", LIT: "S16"},                         //ESTADO 6
	{ID: "S8"},               //ESTADO 7
	{PT_V: "R8", VIR: "S11"}, //ESTADO 8
	{PT_V: "S74"},            //ESTADO 9
	{ID: "R4", LEIA: "R4", ESCREVA: "R4", SE: "R4", FACE_ATE: "R4", FIM: "R4"}, //ESTADO 10
	{ID: "S8"},    //ESTADO 11
	{PT_V: "R7"},  //ESTADO 12
	{PT_V: "S75"}, //ESTADO 13
	{ID: "R9"},    //ESTADO 14
	{ID: "R10"},   //ESTADO 15
	{ID: "R11"},   //ESTADO 16
	{EOF: "R2"},   //ESTADO 17
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FACE_ATE: "S27", FIM: "S28"}, //ESTADO 18
	{ID: "S29"},                             //ESTADO 19
	{ID: "S34", LITERAL: "S32", NUM: "S33"}, //ESTADO 20
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FACE_ATE: "S27", FIM: "S28"}, //ESTADO 21
	{RCB: "S36"}, //ESTADO 22
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FACE_ATE: "S27", FIM: "S28"}, //ESTADO 23
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FIM_SE: "S61"},               //ESTADO 24
	{AB_P: "S44"}, //ESTADO 25
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FACE_ATE: "S27", FIM: "S28"}, //ESTADO 26
	{AB_P: "S51"}, //ESTADO 27
	{EOF: "R38"},  //ESTADO 28
	{PT_V: "S30"}, //ESTADO 29
	{ID: "R13", LEIA: "R13", ESCREVA: "R13", SE: "R13", FIM_SE: "R13", FACE_ATE: "R13", FIM_FACA: "R13", FIM: "R13"}, //ESTADO 30
	{PT_V: "S35"}, //ESTADO 31
	{PT_V: "R15"}, //ESTADO 32
	{PT_V: "R16"}, //ESTADO 33
	{PT_V: "R17"}, //ESTADO 34
	{ID: "R14", LEIA: "R14", ESCREVA: "R14", SE: "R14", FIM_SE: "R14", FACE_ATE: "R14", FIM_FACA: "R14", FIM: "R14"}, //ESTADO 35
	{ID: "S39", NUM: "S40"},   //ESTADO 36
	{PT_V: "S41"},             //ESTADO 37
	{PT_V: "R21", OPM: "S42"}, //ESTADO 38
	{PT_V: "R22", OPM: "R22", FC_P: "R22", OPR: "R22"},                                                               //ESTADO 39
	{PT_V: "R23", OPM: "R23", FC_P: "R23", OPR: "R23"},                                                               //ESTADO 40
	{ID: "R19", LEIA: "R19", ESCREVA: "R19", SE: "R19", FIM_SE: "R19", FACE_ATE: "R19", FIM_FACA: "R19", FIM: "R19"}, //ESTADO 41
	{ID: "S39", NUM: "S40"}, //ESTADO 42
	{PT_V: "R20"},           //ESTADO 43
	{ID: "S39", NUM: "S40"}, //ESTADO 44
	{FC_P: "S46"},           //ESTADO 45
	{ENTAO: "S47"},          //ESTADO 46
	{ID: "R26", LEIA: "R26", ESCREVA: "R26", SE: "R26", FIM_SE: "R26"}, //ESTADO 47
	{OPR: "S49"},            //ESTADO 48
	{ID: "S39", NUM: "S40"}, //ESTADO 49
	{FC_P: "R27"},           //ESTADO 50
	{ID: "S39", NUM: "S40"}, //ESTADO 51
	{FC_P: "S53"},           //ESTADO 52
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FIM_FACA: "S67"}, //ESTADO 53
	{EOF: "R12"}, //ESTADO 54
	{EOF: "R18"}, //ESTADO 55
	{EOF: "R24"}, //ESTADO 56
	{ID: "R25", LEIA: "R25", ESCREVA: "R25", SE: "R25", FIM_SE: "R25", FACE_ATE: "R25", FIM_FACA: "R25", FIM: "R25"}, //ESTADO 57
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FIM_SE: "S61"},                                               //ESTADO 58
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FIM_SE: "S61"},                                               //ESTADO 59
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FIM_SE: "S61"},                                               //ESTADO 60
	{ID: "R31", LEIA: "R31", ESCREVA: "R31", SE: "R31", FIM_SE: "R31", FACE_ATE: "R31", FIM_FACA: "R31", FIM: "R31"}, //ESTADO 61
	{EOF: "R32"}, //ESTADO 62
	{ID: "R33", LEIA: "R33", ESCREVA: "R33", SE: "R33", FACE_ATE: "R33", FIM: "R33"},                                 //ESTADO 63
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FIM_FACA: "S67"},                                             //ESTADO 64
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FIM_FACA: "S67"},                                             //ESTADO 65
	{ID: "S22", LEIA: "S19", ESCREVA: "S20", SE: "S25", FIM_FACA: "S67"},                                             //ESTADO 66
	{ID: "R37", LEIA: "R37", ESCREVA: "R37", SE: "R37", FACE_ATE: "R37", FIM: "R37"},                                 //ESTADO 67
	{ID: "R29", LEIA: "R29", ESCREVA: "R29", SE: "R29", FIM_SE: "R29", FACE_ATE: "R29", FIM_FACA: "R29", FIM: "R29"}, //ESTADO 68
	{ID: "R28", LEIA: "R28", ESCREVA: "R28", SE: "R28", FIM_SE: "R28", FACE_ATE: "R28", FIM_FACA: "R28", FIM: "R28"}, //ESTADO 69
	{ID: "R30", LEIA: "R30", ESCREVA: "R30", SE: "R30", FIM_SE: "R30", FACE_ATE: "R30", FIM_FACA: "R30", FIM: "R30"}, //ESTADO 70
	{ID: "R34", LEIA: "R34", ESCREVA: "R34", SE: "R34", FACE_ATE: "R34", FIM: "R34"},                                 //ESTADO 71
	{ID: "R35", LEIA: "R35", ESCREVA: "R35", SE: "R35", FACE_ATE: "R35", FIM: "R35"},                                 //ESTADO 72
	{ID: "R36", LEIA: "R36", ESCREVA: "R36", SE: "R36", FACE_ATE: "R36", FIM: "R36"},                                 //ESTADO 73
	{ID: "R5", LEIA: "R5", ESCREVA: "R5", SE: "R5", FACE_ATE: "R5", FIM: "R5"},                                       //ESTADO 74
	{VAR_FIM: "R6", INTEIRO: "R6", REAL: "R6", LIT: "R6"},                                                            //ESTADO 75
}
