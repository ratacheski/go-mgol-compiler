package structs

import (
	"github.com/ratacheski/go-mgol-compiler/structs/classe"
	"log"
)

type SymbolTable []Token

var TabelaDeSimbolos SymbolTable

func init() {
	log.Println("Inicializando Tabela de Símbolos")
	TabelaDeSimbolos = TabelaDeSimbolos.NovaTabelaDeSimbolos()
}

func (tabela SymbolTable) NovaTabelaDeSimbolos() SymbolTable {
	inicio := Token{
		Classe: classe.INICIO,
		Lexema: classe.INICIO,
	}
	tabela = append(tabela, inicio)
	varInicio := Token{
		Classe: classe.VAR_INICIO,
		Lexema: classe.VAR_INICIO,
	}
	tabela = append(tabela, varInicio)
	varFim := Token{
		Classe: classe.VAR_FIM,
		Lexema: classe.VAR_FIM,
	}
	tabela = append(tabela, varFim)
	escreva := Token{
		Classe: classe.ESCREVA,
		Lexema: classe.ESCREVA,
	}
	tabela = append(tabela, escreva)
	leia := Token{
		Classe: classe.LEIA,
		Lexema: classe.LEIA,
	}
	tabela = append(tabela, leia)
	se := Token{
		Classe: classe.SE,
		Lexema: classe.SE,
	}
	tabela = append(tabela, se)
	entao := Token{
		Classe: classe.ENTAO,
		Lexema: classe.ENTAO,
	}
	tabela = append(tabela, entao)
	fimSe := Token{
		Classe: classe.FIM_SE,
		Lexema: classe.FIM_SE,
	}
	tabela = append(tabela, fimSe)
	facaAte := Token{
		Classe: classe.FACE_ATE,
		Lexema: classe.FACE_ATE,
	}
	tabela = append(tabela, facaAte)
	fimFaca := Token{
		Classe: classe.FIM_FACA,
		Lexema: classe.FIM_FACA,
	}
	tabela = append(tabela, fimFaca)
	fim := Token{
		Classe: classe.FIM,
		Lexema: classe.FIM,
	}
	tabela = append(tabela, fim)
	inteiro := Token{
		Classe: classe.INTEIRO,
		Lexema: classe.INTEIRO,
	}
	tabela = append(tabela, inteiro)
	lit := Token{
		Classe: classe.LIT,
		Lexema: classe.LIT,
	}
	tabela = append(tabela, lit)
	tokenReal := Token{
		Classe: classe.REAL,
		Lexema: classe.REAL,
	}
	tabela = append(tabela, tokenReal)
	return tabela
}
