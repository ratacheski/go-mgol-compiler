package structs

import "github.com/ratacheski/go-mgol-compiler/structs/classe"

type Token struct {
	classe.Classe
	Lexema string
	Tipo   *Tipo
}

type Tipo struct {
}
