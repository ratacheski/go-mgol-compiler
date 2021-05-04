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

func analiseSintatica() {
	stateStack = append(stateStack, 0)
	columnPosition := 0
	var token structs.Token
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
			columnPosition, token = Scanner(columnPosition, code)
			if token.Classe != "" && token.Classe == classe.ERRO {
				ShowError(token, columnPosition, code)
				columnPosition++
			}
		} else if strings.HasPrefix(action, "R") {
			reduceRule, _ := strconv.Atoi(action[1:])
			NonTerminalReduce := unstackStates(reduceRule)
			goTo := getTransitionField(structs.TransitionTable[stateStack[len(stateStack)-1]], NonTerminalReduce)
			stateStack = append(stateStack, goTo)
		} else {
			fmt.Print("\033[31m") //Colorização Vermelha do log
			fmt.Println(token, action, stateStack, code[columnPosition-1:])
			fmt.Print("\033[0m") // Reset de cor do log
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
		imprimeRegra("P' -> P")
		return "P'"
	case 2: // P -> inicio V A
		stateStack = stateStack[:len(stateStack)-3]
		imprimeRegra("P -> inicio V A")
		return "P"
	case 3: // V -> varinicio LV
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("V -> varinicio LV")
		return "V"
	case 4: // LV -> D LV
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("LV -> D LV")
		return "LV"
	case 5: // LV -> varfim ;
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("LV -> varfim;")
		return "LV"
	case 6: // D -> TIPO L ;
		stateStack = stateStack[:len(stateStack)-3]
		imprimeRegra("D -> TIPO L;")
		return "D"
	case 7: // L -> id , L
		stateStack = stateStack[:len(stateStack)-3]
		imprimeRegra("L -> id , L")
		return "L"
	case 8: // L -> id
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("L -> id")
		return "L"
	case 9: // TIPO -> int
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("TIPO -> int")
		return "TIPO"
	case 10: // TIPO -> real
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("TIPO -> real")
		return "TIPO"
	case 11: // TIPO -> lit
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("TIPO -> lit")
		return "TIPO"
	case 12: // A -> ES A
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("A -> ES A")
		return "A"
	case 13: // ES -> leia id ;
		stateStack = stateStack[:len(stateStack)-3]
		imprimeRegra("ES -> leia id;")
		return "ES"
	case 14: // ES -> escreva ARG ;
		stateStack = stateStack[:len(stateStack)-3]
		imprimeRegra("ES -> escreva ARG;")
		return "ES"
	case 15: // ARG -> literal
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("ARG -> literal")
		return "ARG"
	case 16: // ARG -> num
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("ARG -> num")
		return "ARG"
	case 17: // ARG -> id
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("ARG -> id")
		return "ARG"
	case 18: // A -> CMD A
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("A -> CMD A")
		return "A"
	case 19: // CMD -> id rcb LD ;
		stateStack = stateStack[:len(stateStack)-4]
		imprimeRegra("CMD -> id rcb LD;")
		return "CMD"
	case 20: // LD -> OPRD opm OPRD
		stateStack = stateStack[:len(stateStack)-3]
		imprimeRegra("LD -> OPRD opm OPRD")
		return "LD"
	case 21: // LD -> OPRD
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("LD -> OPRD")
		return "LD"
	case 22: // OPRD -> id
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("OPRD -> id")
		return "OPRD"
	case 23: // OPRD -> num
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("OPRD -> num")
		return "OPRD"
	case 24: // A -> COND A
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("A -> COND A")
		return "A"
	case 25: // COND -> CAB CP
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("COND -> CAB CP")
		return "COND"
	case 26: // CAB -> se (EXP_R) então
		stateStack = stateStack[:len(stateStack)-5]
		imprimeRegra("CAB -> se (EXP_R) então")
		return "CAB"
	case 27: // EXP_R -> OPRD opr OPRD
		stateStack = stateStack[:len(stateStack)-3]
		imprimeRegra("EXP_R -> OPRD opr OPRD")
		return "EXP_R"
	case 28: // CP -> ES CP
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("CP -> ES CP")
		return "CP"
	case 29: // CP -> CMD CP
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("CP -> CMD CP")
		return "CP"
	case 30: // CP -> COND CP
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("CP -> COND CP")
		return "CP"
	case 31: // CP -> fimse
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("CP -> fimse")
		return "CP"
	case 32: // A -> R A
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("A -> R A")
		return "A"
	case 33: // R -> facaAte (EXP_R) CP_R
		stateStack = stateStack[:len(stateStack)-5]
		imprimeRegra("R -> facaAte (EXP_R) CP_R")
		return "R"
	case 34: // CP_R -> ES CP_R
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("CP_R -> ES CP_R")
		return "CP_R"
	case 35: // CP_R -> CMD CP_R
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("CP_R -> CMD CP_R")
		return "CP_R"
	case 36: // CP_R -> COND CP_R
		stateStack = stateStack[:len(stateStack)-2]
		imprimeRegra("CP_R -> COND CP_R")
		return "CP_R"
	case 37: // CP_R -> fimFaca
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("CP_R -> fimFaca")
		return "CP_R"
	case 38: // A -> fim
		stateStack = stateStack[:len(stateStack)-1]
		imprimeRegra("A -> fim")
		return "A"
	}
	return "UNDEFINED"
}

func imprimeRegra(regra string) {
	fmt.Print("\033[32m") //Colorização Verde do log
	fmt.Println(regra)
	fmt.Print("\033[0m") // Reset de cor do log
}
