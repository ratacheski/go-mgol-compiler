package main

import (
	"fmt"
	"github.com/ratacheski/go-mgol-compiler/structs"
	"github.com/ratacheski/go-mgol-compiler/structs/classe"
	"reflect"
	"strconv"
	"strings"
)

var stateStack []int
var columnPosition = 0
var token structs.Token

func analiseSintatica() {
	stateStack = append(stateStack, 0)
	columnPosition, token = Scanner(columnPosition, code)
	if token.Classe != "" && token.Classe == classe.ERRO {
		ShowError(token, columnPosition, code)
		columnPosition++
	}
	for true {
		action := getActionField(structs.ActionTable[stateStack[len(stateStack)-1]], token.Classe)
		if strings.HasPrefix(action, "ACC") {
			break
		} else if strings.HasPrefix(action, "S") {
			newState, _ := strconv.Atoi(action[1:])
			stateStack = append(stateStack, newState)
			for {
				columnPosition, token = Scanner(columnPosition, code)
				if token.Classe != "" && token.Classe == classe.ERRO {
					ShowError(token, columnPosition, code)
					columnPosition++
				} else if token.Classe != "" && token.Classe != classe.COMENTARIO {
					break
				}
			}
		} else if strings.HasPrefix(action, "R") {
			reduceRule, _ := strconv.Atoi(action[1:])
			NonTerminalReduce := unstackStates(reduceRule)
			goTo := getTransitionField(structs.TransitionTable[stateStack[len(stateStack)-1]], NonTerminalReduce)
			stateStack = append(stateStack, goTo)
		} else {
			errorHandling()
		}
	}

}

func getActionField(v *structs.ActionRow, field classe.Classe) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(classe.Classe.String(field))
	return f.String()
}

func getTransitionField(v *structs.TransitionRow, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if f.IsValid() {
		return int(f.Int())
	}
	return 0
}

func unstackStates(reduceRule int) string {
	switch reduceRule {
	case 1: // P' -> P
		stateStack = stateStack[:len(stateStack)-1]
		printRule("P' -> P")
		return "P'"
	case 2: // P -> inicio V A
		stateStack = stateStack[:len(stateStack)-3]
		printRule("P -> inicio V A")
		return "P"
	case 3: // V -> varinicio LV
		stateStack = stateStack[:len(stateStack)-2]
		printRule("V -> varinicio LV")
		return "V"
	case 4: // LV -> D LV
		stateStack = stateStack[:len(stateStack)-2]
		printRule("LV -> D LV")
		return "LV"
	case 5: // LV -> varfim ;
		stateStack = stateStack[:len(stateStack)-2]
		printRule("LV -> varfim;")
		return "LV"
	case 6: // D -> TIPO L ;
		stateStack = stateStack[:len(stateStack)-3]
		printRule("D -> TIPO L;")
		return "D"
	case 7: // L -> id , L
		stateStack = stateStack[:len(stateStack)-3]
		printRule("L -> id , L")
		return "L"
	case 8: // L -> id
		stateStack = stateStack[:len(stateStack)-1]
		printRule("L -> id")
		return "L"
	case 9: // TIPO -> int
		stateStack = stateStack[:len(stateStack)-1]
		printRule("TIPO -> int")
		return "TIPO"
	case 10: // TIPO -> real
		stateStack = stateStack[:len(stateStack)-1]
		printRule("TIPO -> real")
		return "TIPO"
	case 11: // TIPO -> lit
		stateStack = stateStack[:len(stateStack)-1]
		printRule("TIPO -> lit")
		return "TIPO"
	case 12: // A -> ES A
		stateStack = stateStack[:len(stateStack)-2]
		printRule("A -> ES A")
		return "A"
	case 13: // ES -> leia id ;
		stateStack = stateStack[:len(stateStack)-3]
		printRule("ES -> leia id;")
		return "ES"
	case 14: // ES -> escreva ARG ;
		stateStack = stateStack[:len(stateStack)-3]
		printRule("ES -> escreva ARG;")
		return "ES"
	case 15: // ARG -> literal
		stateStack = stateStack[:len(stateStack)-1]
		printRule("ARG -> literal")
		return "ARG"
	case 16: // ARG -> num
		stateStack = stateStack[:len(stateStack)-1]
		printRule("ARG -> num")
		return "ARG"
	case 17: // ARG -> id
		stateStack = stateStack[:len(stateStack)-1]
		printRule("ARG -> id")
		return "ARG"
	case 18: // A -> CMD A
		stateStack = stateStack[:len(stateStack)-2]
		printRule("A -> CMD A")
		return "A"
	case 19: // CMD -> id rcb LD ;
		stateStack = stateStack[:len(stateStack)-4]
		printRule("CMD -> id rcb LD;")
		return "CMD"
	case 20: // LD -> OPRD opm OPRD
		stateStack = stateStack[:len(stateStack)-3]
		printRule("LD -> OPRD opm OPRD")
		return "LD"
	case 21: // LD -> OPRD
		stateStack = stateStack[:len(stateStack)-1]
		printRule("LD -> OPRD")
		return "LD"
	case 22: // OPRD -> id
		stateStack = stateStack[:len(stateStack)-1]
		printRule("OPRD -> id")
		return "OPRD"
	case 23: // OPRD -> num
		stateStack = stateStack[:len(stateStack)-1]
		printRule("OPRD -> num")
		return "OPRD"
	case 24: // A -> COND A
		stateStack = stateStack[:len(stateStack)-2]
		printRule("A -> COND A")
		return "A"
	case 25: // COND -> CAB CP
		stateStack = stateStack[:len(stateStack)-2]
		printRule("COND -> CAB CP")
		return "COND"
	case 26: // CAB -> se (EXP_R) então
		stateStack = stateStack[:len(stateStack)-5]
		printRule("CAB -> se (EXP_R) então")
		return "CAB"
	case 27: // EXP_R -> OPRD opr OPRD
		stateStack = stateStack[:len(stateStack)-3]
		printRule("EXP_R -> OPRD opr OPRD")
		return "EXP_R"
	case 28: // CP -> ES CP
		stateStack = stateStack[:len(stateStack)-2]
		printRule("CP -> ES CP")
		return "CP"
	case 29: // CP -> CMD CP
		stateStack = stateStack[:len(stateStack)-2]
		printRule("CP -> CMD CP")
		return "CP"
	case 30: // CP -> COND CP
		stateStack = stateStack[:len(stateStack)-2]
		printRule("CP -> COND CP")
		return "CP"
	case 31: // CP -> fimse
		stateStack = stateStack[:len(stateStack)-1]
		printRule("CP -> fimse")
		return "CP"
	case 32: // A -> R A
		stateStack = stateStack[:len(stateStack)-2]
		printRule("A -> R A")
		return "A"
	case 33: // R -> facaAte (EXP_R) CP_R
		stateStack = stateStack[:len(stateStack)-5]
		printRule("R -> facaAte (EXP_R) CP_R")
		return "R"
	case 34: // CP_R -> ES CP_R
		stateStack = stateStack[:len(stateStack)-2]
		printRule("CP_R -> ES CP_R")
		return "CP_R"
	case 35: // CP_R -> CMD CP_R
		stateStack = stateStack[:len(stateStack)-2]
		printRule("CP_R -> CMD CP_R")
		return "CP_R"
	case 36: // CP_R -> COND CP_R
		stateStack = stateStack[:len(stateStack)-2]
		printRule("CP_R -> COND CP_R")
		return "CP_R"
	case 37: // CP_R -> fimFaca
		stateStack = stateStack[:len(stateStack)-1]
		printRule("CP_R -> fimFaca")
		return "CP_R"
	case 38: // A -> fim
		stateStack = stateStack[:len(stateStack)-1]
		printRule("A -> fim")
		return "A"
	}
	return "UNDEFINED"
}

func printRule(rule string) {
	fmt.Print("\033[32m") //Colorização Verde do log
	fmt.Println(rule)
	fmt.Print("\033[0m") // Reset de cor do log
}

func printError(errorMsg string) {
	split := strings.Split(code[:columnPosition+1], "\n")
	columnErrorPosition := len(split[len(split)-1])
	lineErrorPosition := len(split)
	fmt.Print("\033[31m") //Colorização Vermelha do log
	fmt.Println("ERRO SINTÁTICO - ", errorMsg, ", linha", lineErrorPosition, ", coluna", columnErrorPosition)
	fmt.Print("\033[0m") // Reset de cor do log
}

func errorHandling() {
	fmt.Printf("%v", stateStack)
	switch stateStack[len(stateStack)-1] {
	case 0:
		printError("Esperado comando inicio como primeiro comando do algoritmo")
		classes := []classe.Classe{classe.INICIO}
		syncTokens(classes)
		break
	case 2:
		printError("Comando inicio deve ser seguido de declaração de variáveis")
		classes := []classe.Classe{classe.VAR_INICIO}
		syncTokens(classes)
		break
	case 3:
		printError("Instrução deve iniciar com leia, escreva, se, facaAte, ou ser uma atribuição de variável, ou finalização de código")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM}
		syncTokens(classes)
		break
	case 4, 6:
		printError("Bloco de declaração de variáveis deve conter apenas declarações e a finalização do bloco")
		classes := []classe.Classe{classe.VAR_FIM, classe.PT_V, classe.INTEIRO, classe.LIT, classe.REAL}
		syncTokens(classes)
		break
	case 8:
		printError("Declaração múltipla de variável deve ser separada por vírgula e finalizada por ponto e vírgula")
		classes := []classe.Classe{classe.PT_V, classe.INTEIRO, classe.LIT, classe.REAL}
		syncTokens(classes)
		break
	case 9:
		printError("Comando varfim deve ser finalizado com ;")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM}
		syncTokens(classes)
		break
	case 7, 11, 13:
		printError("Declaração de variável deve conter um nome válido")
		classes := []classe.Classe{classe.PT_V, classe.INTEIRO, classe.LIT, classe.REAL, classe.VAR_FIM, classe.ID}
		syncTokens(classes)
		break
	case 14:
		printError("Declaração de inteiro deve conter um nome válido")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM}
		syncTokens(classes)
		break
	case 15:
		printError("Declaração de real deve conter um nome válido")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM}
		syncTokens(classes)
		break
	case 16:
		printError("Declaração de literal deve conter um nome válido")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM}
		syncTokens(classes)
		break
	case 19:
		printError("Comando de leitura deve ser seguido de uma variável ")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM}
		syncTokens(classes)
		break
	case 29:
		printError("Comando de leitura deve ser encerrado com ponto e vírgula")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM, classe.PT_V}
		syncTokens(classes)
		break
	case 20, 32, 33, 34:
		printError("Comando de escrita deve ser seguido de uma variável, número ou literal e finalizar em ponto e vírgula")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM_SE, classe.FIM_FACA, classe.FIM, classe.PT_V}
		syncTokens(classes)
		break
	case 74:
		printError("Comando varfim deve ser prosseguido com leia, escreva, se, facaAte, ou ser uma atribuição de variável, ou finalização de código")
		classes := []classe.Classe{classe.LEIA, classe.ESCREVA, classe.ID, classe.SE, classe.FACE_ATE, classe.FIM}
		syncTokens(classes)
		break
	case 75:
		printError("Declaração de variável deve ser prosseguida de outra declaração ou de finalização de bloco de declaração")
		classes := []classe.Classe{classe.VAR_FIM, classe.PT_V, classe.INTEIRO, classe.LIT, classe.REAL}
		syncTokens(classes)
		break
	}
}

func syncTokens(classes []classe.Classe) {
	for {
		columnPosition, token = Scanner(columnPosition, code)
		if token.Classe != "" && token.Classe == classe.ERRO {
			ShowError(token, columnPosition, code)
			columnPosition++
		} else if token.Classe != "" && containsClasse(classes, token.Classe) {
			break
		}
	}
}

func containsClasse(s []classe.Classe, e classe.Classe) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
