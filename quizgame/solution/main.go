package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	defaultFilename = "problems.csv"
	defaultTimeout  = 30 * time.Second
)

func main() {
	filename := flag.String("file", defaultFilename, "fichero csv con las preguntas y las respuestas")
	timeout := flag.Duration("timeout", defaultTimeout, "tiempo máximo para responder a las preguntas")
	shuffle := flag.Bool("shuffle", false, "orden aleatorio de las preguntas")
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
	if *shuffle {
		data, err = toShuffle(data)
		if err != nil {
			panic(err)
		}
	}

	inReader := bufio.NewReader(os.Stdin)

	fmt.Println("¿Estás listo? Dale a enter cuándo quieras empezar")
	_, _ = inReader.ReadString('\n')

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	var asserts int

loop:
	for _, row := range data {
		if len(row) != 2 {
			panic("liada")
		}
		fmt.Println(row[0])

		ch := make(chan string)

		go func() {
			in, _ := inReader.ReadString('\n')
			ch <- in
		}()

		select {
		case userInput := <-ch:
			if row[1] == strings.Replace(userInput, "\n", "", -1) {
				asserts++
			}
		case <-ctx.Done():
			break loop
		}
	}

	fmt.Printf("Has acertado: %d de %d\n", asserts, len(data))
}

func toShuffle(data [][]string) ([][]string, error) {
	m := make(map[string]string)
	for _, row := range data {
		if len(row) != 2 {
			return nil, errors.New("max length 2")
		}

		m[row[0]] = row[1]
	}

	out := make([][]string, 0, len(data))
	for k, v := range m {
		out = append(out, []string{k, v})
	}

	return out, nil
}
