package main

import (
	"bufio"
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
	for index, line := range lines {
		log.Println(index, ": ", line)
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
