package main

import (
	"bufio"
	"fmt"
	"github.com/ratacheski/go-mgol-compiler/structs"
	"github.com/ratacheski/go-mgol-compiler/structs/classe"
	"log"
	"os"
)

const (
	EXTENSAO = "alg"
)

var (
	lines []string
)

func main() {
	leArquivoFonte()
	for linePosition, line := range lines {
		columnPosition := 0
		for columnPosition < len(line) {
			var token structs.Token
			columnPosition, token = Scanner(columnPosition, line)
			if token.Classe != "" && token.Classe != classe.ERRO {
				fmt.Print("\033[32m") //Colorização Verde do log
				log.Println("Classe:", token.Classe, ", Lexema:", token.Lexema, ", Tipo:", token.Tipo)
				fmt.Print("\033[0m") // Reset de cor do log
			} else {
				ShowError(token, columnPosition+1, linePosition+1)
				columnPosition++
			}
		}
	}
}

func leArquivoFonte() {
	if len(os.Args) < 2 {
		log.Fatal("Forneça o caminho do arquivo para compilar")
	}
	path := os.Args[1]
	if path[len(path)-3:] != EXTENSAO {
		log.Fatal("O arquivo deve possuir a extensão " + EXTENSAO + " para ser compilado")
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Erro ao ler arquivo fonte")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
