package structs

import (
	"fmt"
	"github.com/ratacheski/go-mgol-compiler/structs/classe"
	"log"
)

type SymbolTable map[string]Token

var TabelaDeSimbolos SymbolTable

func init() {
	fmt.Print("\033[33m") //Colorização Amarela do log
	log.Println("Inicializando Tabela de Símbolos")
	fmt.Print("\033[0m") // Reset de cor do log
	TabelaDeSimbolos = make(map[string]Token)
	TabelaDeSimbolos = TabelaDeSimbolos.NovaTabelaDeSimbolos()
}

func AddTokenToSymbolTableIfNotExists(token Token) (returnToken Token) {
	if returnToken, present := TabelaDeSimbolos[token.Lexema]; !present {
		TabelaDeSimbolos[token.Lexema] = token
		return token
	} else {
		return returnToken
	}
}

func (tabela SymbolTable) NovaTabelaDeSimbolos() SymbolTable {
	tabela[classe.INICIO] = Token{
		Classe: classe.INICIO,
		Lexema: classe.INICIO,
	}
	tabela[classe.VAR_INICIO] = Token{
		Classe: classe.VAR_INICIO,
		Lexema: classe.VAR_INICIO,
	}
	tabela[classe.VAR_FIM] = Token{
		Classe: classe.VAR_FIM,
		Lexema: classe.VAR_FIM,
	}
	tabela[classe.ESCREVA] = Token{
		Classe: classe.ESCREVA,
		Lexema: classe.ESCREVA,
	}
	tabela[classe.LEIA] = Token{
		Classe: classe.LEIA,
		Lexema: classe.LEIA,
	}
	tabela[classe.SE] = Token{
		Classe: classe.SE,
		Lexema: classe.SE,
	}
	tabela[classe.ENTAO] = Token{
		Classe: classe.ENTAO,
		Lexema: classe.ENTAO,
	}
	tabela[classe.FIM_SE] = Token{
		Classe: classe.FIM_SE,
		Lexema: classe.FIM_SE,
	}
	tabela[classe.FACE_ATE] = Token{
		Classe: classe.FACE_ATE,
		Lexema: classe.FACE_ATE,
	}
	tabela[classe.FIM_FACA] = Token{
		Classe: classe.FIM_FACA,
		Lexema: classe.FIM_FACA,
	}
	tabela[classe.FIM] = Token{
		Classe: classe.FIM,
		Lexema: classe.FIM,
	}
	tabela[classe.INTEIRO] = Token{
		Classe: classe.INTEIRO,
		Lexema: classe.INTEIRO,
	}
	tabela[classe.LIT] = Token{
		Classe: classe.LIT,
		Lexema: classe.LIT,
	}
	tabela[classe.REAL] = Token{
		Classe: classe.REAL,
		Lexema: classe.REAL,
	}
	return tabela
}
