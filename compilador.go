package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	EXTENSAO = "alg"
)

var (
	code string
)

func main() {
	leArquivoFonte()
	//analiseLexica()
	analiseSintatica()
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
