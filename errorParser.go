package main

import (
	"fmt"
	"github.com/ratacheski/go-mgol-compiler/structs"
	"log"
	"strings"
)

func ShowError(token structs.Token, columnPosition int, linePosition int) {
	errorMsg := ""
	switch token.Lexema {
	case "erro1":
		errorMsg = "Caractere inválido na linguagem"
		break
	case "erro2":
		errorMsg = "Um comentário deve ser fechado com chaves (})"
		break
	case "erro3":
		errorMsg = "Sinal < deve compor um recebimento ou um comparador de menor, menor igual ou diferente"
		break
	case "erro4":
		errorMsg = "Um valor literal deve ser fechado com aspas duplas (\")"
		break
	case "erro5":
		errorMsg = "Um número com ponto flutuante deve possuir algarismo após o símbolo decimal"
		break
	case "erro6":
		errorMsg = "Um número com notação científica deve possuir algarismos ou um sinal após o símbolo de exponenciação"
		break
	case "erro7":
		errorMsg = "Um número com notação científica deve possuir algarismos após o sinal do expoente"
		break
	}
	fmt.Print("\033[31m") //Colorização Vermelha do log
	log.Println(strings.ToUpper(token.Lexema), "-", errorMsg, ", linha", linePosition, ", coluna", columnPosition)
	fmt.Print("\033[0m") // Reset de cor do log
}
