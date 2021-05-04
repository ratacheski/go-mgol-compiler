package classe

type Classe string

const (
	NUM        Classe = "Num"
	LITERAL           = "Literal"
	ID                = "id"
	COMENTARIO        = "comentário"
	EOF               = "EOF"
	OPR               = "OPR"
	RCB               = "RCB"
	OPM               = "OPM"
	AB_P              = "AB_P"
	FC_P              = "FC_P"
	VIR               = "VIR"
	PT_V              = "PT_V"
	ERRO              = "ERRO"
	INICIO            = "inicio"
	VAR_INICIO        = "varinicio"
	VAR_FIM           = "varfim"
	ESCREVA           = "escreva"
	LEIA              = "leia"
	SE                = "se"
	ENTAO             = "entao"
	FIM_SE            = "fimse"
	FACE_ATE          = "faca-ate"
	FIM_FACA          = "fimfaca"
	FIM               = "fim"
	INTEIRO           = "inteiro"
	LIT               = "lit"
	REAL              = "real"
)
func (value Classe) String() string {
	switch value {
	case NUM:
		return "NUM"
	case LITERAL:
		return "LITERAL"
	case ID:
		return "ID"
	case COMENTARIO:
		return "COMENTARIO"
	case EOF:
		return "EOF"
	case OPR:
		return "OPR"
	case RCB:
		return "RCB"
	case OPM:
		return "OPM"
	case AB_P:
		return "AB_P"
	case FC_P:
		return "FC_P"
	case VIR:
		return "VIR"
	case PT_V:
		return "PT_V"
	case ERRO:
		return "ERRO"
	case INICIO:
		return "INICIO"
	case VAR_INICIO:
		return "VAR_INICIO"
	case VAR_FIM:
		return "VAR_FIM"
	case ESCREVA:
		return "ESCREVA"
	case LEIA:
		return "LEIA"
	case SE:
		return "SE"
	case FIM_SE:
		return "FIM_SE"
	case FACE_ATE:
		return "FACE_ATE"
	case FIM_FACA:
		return "FIM_FACA"
	case FIM:
		return "FIM"
	case INTEIRO:
		return "INTEIRO"
	case LIT:
		return "LIT"
	case REAL:
		return "REAL"
	case ENTAO:
		return "ENTAO"
	}
	return "UNDEFINED"
}
