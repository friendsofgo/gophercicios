package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	quizzes := []byte(`5+5,10
7+3,10
1+1,2
`)

	r := bytes.NewReader(quizzes)
	csvReader := csv.NewReader(r)

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
