package main

import (
	"github.com/ratacheski/go-mgol-compiler/structs"
	"github.com/ratacheski/go-mgol-compiler/structs/classe"
	"unicode"
)

func getCaracteresDeLoopInicial() []rune {
	return []rune(" \n\t")
}

func getCaracteresDeOperadorMatematico() []rune {
	return []rune("+-*/")
}

func contains(s []rune, str rune) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func Scanner(initialPosition int, line string) (finalPosition int, token structs.Token) {
	currentStatus := 0
	scannedValue := []rune(line[initialPosition:])
	lexema := []rune("")
	for index := 0; index <= len(scannedValue); index++ {
		switch currentStatus {
		case 0:
			currentStatus = validaQ0(scannedValue[index])
			if currentStatus != 0 && currentStatus != -1 {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 1:
			currentStatus = validaQ1(scannedValue[index])
			lexema = append(lexema, scannedValue[index])
			break
		case 2:
			token.Classe = classe.OPR
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 3:
			currentStatus = validaQ3(scannedValue[index])
			if currentStatus != -1 {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 4:
			if index == len(scannedValue) {
				token.Classe = classe.ID
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				//TODO adicionar na tabela de símbolos caso não exista
				return
			}
			currentStatus = validaQ4(scannedValue[index])
			if currentStatus == 0 {
				token.Classe = classe.ID
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				//TODO adicionar na tabela de símbolos caso não exista
				return initialPosition + index, token
			} else if currentStatus == -1 {
				token.Classe = classe.ID
				token.Lexema = string(scannedValue[:initialPosition+index-1])
				finalPosition = initialPosition + index
				//TODO adicionar na tabela de símbolos caso não exista
				return
			} else {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 5:
			token.Classe = classe.VIR
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 6:
			token.Classe = classe.PT_V
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 7:
			if index == len(scannedValue) {
				token.Classe = classe.OPR
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			}
			currentStatus = validaQ7(scannedValue[index])
			if currentStatus == 0 {
				token.Classe = classe.OPR
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return initialPosition + index, token
			} else if currentStatus == -1 {
				token.Classe = classe.OPR
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			} else {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 8:
			if index == len(scannedValue) {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			}
			currentStatus = validaQ8(scannedValue[index])
			if currentStatus == 0 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return initialPosition + index, token
			} else if currentStatus == -1 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			} else {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 9:
			currentStatus = validaQ9(scannedValue[index])
			lexema = append(lexema, scannedValue[index])
			break
		case 11:
			token.Classe = classe.FC_P
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 12:
			token.Classe = classe.AB_P
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 13:
			token.Classe = classe.OPM
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 14:
			token.Classe = classe.COMENTARIO
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 15:
			token.Classe = classe.RCB
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 16:
			currentStatus = validaQ16(scannedValue[index])
			if currentStatus != -1 {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 17:
			currentStatus = validaQ17(scannedValue[index])
			if currentStatus != -1 {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 18:
			token.Classe = classe.LITERAL
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 19:
			if index == len(scannedValue) {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			}
			currentStatus = validaQ19(scannedValue[index])
			if currentStatus == 0 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return initialPosition + index, token
			} else if currentStatus == -1 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			} else {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 20:
			currentStatus = validaQ20(scannedValue[index])
			if currentStatus != -1 {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 21:
			if index == len(scannedValue) {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			}
			currentStatus = validaQ21(scannedValue[index])
			if currentStatus == 0 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return initialPosition + index, token
			} else if currentStatus == -1 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			} else {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 22:
			token.Classe = classe.OPR
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		}

	}
	return
}

func validaQ0(character rune) (targetStatus int) {
	if contains(getCaracteresDeLoopInicial(), character) {
		return 0
	} else if character == '{' {
		return 1
	} else if character == '=' {
		return 2
	} else if character == '<' {
		return 3
	} else if unicode.IsLetter(character) {
		return 4
	} else if character == ',' {
		return 5
	} else if character == ';' {
		return 6
	} else if character == '>' {
		return 7
	} else if unicode.IsNumber(character) {
		return 8
	} else if character == '"' {
		return 9
	} else if character == ')' {
		return 11
	} else if character == '(' {
		return 12
	} else if contains(getCaracteresDeOperadorMatematico(), character) {
		return 13
	} else {
		return -1
	}
}

func validaQ1(character rune) (targetStatus int) {
	if character == '}' {
		return 14
	}
	return 1
}

func validaQ3(character rune) (targetStatus int) {
	if character == '-' {
		return 15
	} else if character == '>' || character == '=' {
		return 22
	}
	return -1
}
func validaQ4(character rune) (targetStatus int) {
	if unicode.IsLetter(character) || unicode.IsNumber(character) || character == '_' {
		return 4
	} else if contains(getCaracteresDeLoopInicial(), character) {
		return 0
	}
	return -1
}

func validaQ7(character rune) (targetStatus int) {
	if character == '=' {
		return 7
	} else if contains(getCaracteresDeLoopInicial(), character) {
		return 0
	}
	return -1
}

func validaQ8(character rune) (targetStatus int) {
	if unicode.IsNumber(character) {
		return 8
	} else if character == '.' {
		return 16
	} else if character == 'E' || character == 'e' {
		return 17
	} else if contains(getCaracteresDeLoopInicial(), character) {
		return 0
	}
	return -1
}

func validaQ9(character rune) (targetStatus int) {
	if character == '"' {
		return 18
	}
	return 9
}

func validaQ16(character rune) (targetStatus int) {
	if unicode.IsNumber(character) {
		return 19
	}
	return -1
}

func validaQ17(character rune) (targetStatus int) {
	if unicode.IsNumber(character) {
		return 21
	} else if character == '+' || character == '-' {
		return 20
	}
	return -1
}

func validaQ19(character rune) (targetStatus int) {
	if unicode.IsNumber(character) {
		return 19
	} else if character == 'E' || character == 'e' {
		return 17
	} else if contains(getCaracteresDeLoopInicial(), character) {
		return 0
	}
	return -1
}

func validaQ20(character rune) (targetStatus int) {
	if unicode.IsNumber(character) {
		return 21
	}
	return -1
}

func validaQ21(character rune) (targetStatus int) {
	if unicode.IsNumber(character) {
		return 21
	} else if contains(getCaracteresDeLoopInicial(), character) {
		return 0
	}
	return -1
}
