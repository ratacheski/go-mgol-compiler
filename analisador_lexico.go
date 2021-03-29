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

func Scanner(initialPosition int, fileContent string) (finalPosition int, token structs.Token) {
	currentStatus := 0
	scannedValue := fileContent[initialPosition:]
	lexema := []byte("")
	for index := 0; index <= len(scannedValue); index++ {
		switch currentStatus {
		case 0:
			currentStatus = validaQ0(rune(scannedValue[index]))
			if currentStatus == -1 {
				finalPosition = initialPosition + index
				token.Classe = classe.ERRO
				token.Lexema = "erro1"
				return
			} else if currentStatus != 0 {
				lexema = append(lexema, scannedValue[index])
			}
			break
		case 1:
			if index == len(scannedValue) {
				finalPosition = initialPosition + index
				token.Classe = classe.ERRO
				token.Lexema = "erro2"
				return
			}
			currentStatus = validaQ1(rune(scannedValue[index]))
			lexema = append(lexema, scannedValue[index])
			break
		case 2:
			token.Classe = classe.OPR
			token.Lexema = string(lexema)
			finalPosition = initialPosition + index
			return
		case 3:
			currentStatus = validaQ3(rune(scannedValue[index]))
			if currentStatus == -1 {
				finalPosition = initialPosition + index
				token.Classe = classe.ERRO
				token.Lexema = "erro3"
				return
			}
			lexema = append(lexema, scannedValue[index])
			break
		case 4:
			if index == len(scannedValue) {
				token.Classe = classe.ID
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				token = structs.AddTokenToSymbolTableIfNotExists(token)
				return
			}
			currentStatus = validaQ4(rune(scannedValue[index]))
			if currentStatus == 0 {
				token.Classe = classe.ID
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				token = structs.AddTokenToSymbolTableIfNotExists(token)
				return
			} else if currentStatus == -1 {
				token.Classe = classe.ID
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				token = structs.AddTokenToSymbolTableIfNotExists(token)
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
			currentStatus = validaQ7(rune(scannedValue[index]))
			if currentStatus == 0 {
				token.Classe = classe.OPR
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
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
			currentStatus = validaQ8(rune(scannedValue[index]))
			if currentStatus == 0 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
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
			if index == len(scannedValue) {
				finalPosition = initialPosition + index
				token.Classe = classe.ERRO
				token.Lexema = "erro4"
				return
			}
			currentStatus = validaQ9(rune(scannedValue[index]))
			lexema = append(lexema, scannedValue[index])
			break
		case 10:
			token.Classe = classe.EOF
			token.Lexema = classe.EOF
			finalPosition = initialPosition + index
			return
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
			currentStatus = validaQ16(rune(scannedValue[index]))
			if currentStatus == -1 {
				finalPosition = initialPosition + index
				token.Classe = classe.ERRO
				token.Lexema = "erro5"
				return
			}
			lexema = append(lexema, scannedValue[index])
			break
		case 17:
			currentStatus = validaQ17(rune(scannedValue[index]))
			if currentStatus == -1 {
				finalPosition = initialPosition + index
				token.Classe = classe.ERRO
				token.Lexema = "erro6"
				return
			}
			lexema = append(lexema, scannedValue[index])
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
			currentStatus = validaQ19(rune(scannedValue[index]))
			if currentStatus == 0 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
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
			currentStatus = validaQ20(rune(scannedValue[index]))
			if currentStatus == -1 {
				finalPosition = initialPosition + index
				token.Classe = classe.ERRO
				token.Lexema = "erro7"
				return
			}
			lexema = append(lexema, scannedValue[index])
			break
		case 21:
			if index == len(scannedValue) {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
			}
			currentStatus = validaQ21(rune(scannedValue[index]))
			if currentStatus == 0 {
				token.Classe = classe.NUM
				token.Lexema = string(lexema)
				finalPosition = initialPosition + index
				return
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
	} else if character == 3 {
		return 10
	} else if character == ')' {
		return 11
	} else if character == '(' {
		return 12
	} else if contains(getCaracteresDeOperadorMatematico(), character) {
		return 13
	}
	return -1
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
