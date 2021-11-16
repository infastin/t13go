package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/infastin/gul/polynom"
)

func main() {
	file, _ := os.Open("polynoms.txt")
	defer file.Close()

	var polynoms []*polynom.Polynomial

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		var monoms []polynom.Monomial

		scanner := bufio.NewScanner(bytes.NewBuffer(line))
		scanner.Split(bufio.ScanWords)

		i := 0
		for scanner.Scan() {
			var coef float64

			fmt.Sscanf(scanner.Text(), "%f", &coef)
			if coef != 0 {
				monoms = append(monoms, polynom.Monomial{
					Coef:   coef,
					Degree: i,
				})
			}

			i++
		}

		polynoms = append(polynoms, polynom.New(monoms...))
	}

	a := polynoms[0]
	b := polynoms[1]
	c := polynoms[2]

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)

	fmt.Printf("a + b = %v\n", a.Add(b))
	fmt.Printf("a - b = %v\n", a.Sub(b))
	fmt.Printf("a * b = %v\n", a.Mul(b))
	fmt.Printf("a / b = %v\n", a.Div(b))
	fmt.Printf("a %% b = %v\n", a.Mod(b))

	fmt.Println()

	fmt.Printf("euqlid(a, b) = %v\n", a.Euclidean(b))
	fmt.Printf("euqlid(a, c) = %v\n", a.Euclidean(c))

	fmt.Println()

	fmt.Printf("c / euqlid(a, c) = %v\n", c.Div(a.Euclidean(c)))
	fmt.Printf("a / euqlid(a, c) = %v\n", a.Div(a.Euclidean(c)))
}
