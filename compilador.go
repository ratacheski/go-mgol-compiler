package main

import (
	"fmt"
	"github.com/ratacheski/go-mgol-compiler/structs"
	"github.com/ratacheski/go-mgol-compiler/structs/classe"
	"io/ioutil"
	"os"
)

const (
	EXTENSAO = "alg"
)

var (
	code  string
	lines []string
)

func main() {
	leArquivoFonte()
	columnPosition := 0
	for columnPosition < len(code) {
		var token structs.Token
		columnPosition, token = Scanner(columnPosition, code)
		if token.Classe != "" && token.Classe != classe.ERRO {
			fmt.Print("\033[32m") //Colorização Verde do log
			fmt.Printf("Classe: %s, Lexema: %s, Tipo: %+v\n", token.Classe, token.Lexema, token.Tipo)
			fmt.Print("\033[0m") // Reset de cor do log
		} else {
			ShowError(token, columnPosition, code)
			columnPosition++
		}
	}
}

func leArquivoFonte() {
	if len(os.Args) < 2 {
		fmt.Print("\033[31m") //Colorização Vermelha do log
		fmt.Println("Forneça o caminho do arquivo para compilar")
		fmt.Print("\033[0m") // Reset de cor do log
		os.Exit(1)
	}
	path := os.Args[1]
	if path[len(path)-3:] != EXTENSAO {
		fmt.Print("\033[31m") //Colorização Vermelha do log
		fmt.Println("O arquivo deve possuir a extensão " + EXTENSAO + " para ser compilado")
		fmt.Print("\033[0m") // Reset de cor do log
		os.Exit(1)
	}
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print("\033[31m") //Colorização Vermelha do log
		fmt.Println("Erro ao ler arquivo fonte")
		fmt.Print("\033[0m") // Reset de cor do log
		os.Exit(1)
	}
	//Adição do símbolo ASCII de fim de texto
	fileContent = append(fileContent, 3)
	code = string(fileContent)
}
