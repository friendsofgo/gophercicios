package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

const defaultFilename = "problems.csv"

func main() {
	filename := flag.String("file", defaultFilename, "fichero csv con las preguntas y las respuestas")
	flag.Parse()

	csvFile, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(csvFile)

	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var asserts int
	for _, row := range data {
		if len(row) != 2 {
			panic("liada")
		}
		fmt.Println(row[0])
		inReader := bufio.NewReader(os.Stdin)
		in, _ := inReader.ReadString('\n')

		if row[1] == strings.Replace(in, "\n", "", -1) {
			asserts++
		}

	}
	fmt.Printf("Has acertado: %d de %d\n", asserts, len(data))
}
