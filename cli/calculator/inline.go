package main

import (
	"fmt"
	"os"
	"strings"

	shuntingyard "github.com/VixsTy/calculator/pkg/calculator/shunting-yard"
)

func inline() {
	input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
	calc := shuntingyard.NewShuntingYard()
	fmt.Println(calc.Calc(input))
}
