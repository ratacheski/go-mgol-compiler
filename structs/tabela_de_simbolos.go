package structs

type TabelaDeSimbolos []Token

func NovaTabelaDeSimbolos() (tabela TabelaDeSimbolos) {
	inicio := Token{
		Classe: INICIO,
		Lexema: INICIO,
	}
	tabela = append(tabela, inicio)
	varInicio := Token{
		Classe: VAR_INICIO,
		Lexema: VAR_INICIO,
	}
	tabela = append(tabela, varInicio)
	varFim := Token{
		Classe: VAR_FIM,
		Lexema: VAR_FIM,
	}
	tabela = append(tabela, varFim)
	escreva := Token{
		Classe: ESCREVA,
		Lexema: ESCREVA,
	}
	tabela = append(tabela, escreva)
	leia := Token{
		Classe: LEIA,
		Lexema: LEIA,
	}
	tabela = append(tabela, leia)
	se := Token{
		Classe: SE,
		Lexema: SE,
	}
	tabela = append(tabela, se)
	entao := Token{
		Classe: ENTAO,
		Lexema: ENTAO,
	}
	tabela = append(tabela, entao)
	fimSe := Token{
		Classe: FIM_SE,
		Lexema: FIM_SE,
	}
	tabela = append(tabela, fimSe)
	facaAte := Token{
		Classe: FACE_ATE,
		Lexema: FACE_ATE,
	}
	tabela = append(tabela, facaAte)
	fimFaca := Token{
		Classe: FIM_FACA,
		Lexema: FIM_FACA,
	}
	tabela = append(tabela, fimFaca)
	fim := Token{
		Classe: FIM,
		Lexema: FIM,
	}
	tabela = append(tabela, fim)
	inteiro := Token{
		Classe: INTEIRO,
		Lexema: INTEIRO,
	}
	tabela = append(tabela, inteiro)
	lit := Token{
		Classe: LIT,
		Lexema: LIT,
	}
	tabela = append(tabela, lit)
	tokenReal := Token{
		Classe: REAL,
		Lexema: REAL,
	}
	tabela = append(tabela, tokenReal)
	return
}
