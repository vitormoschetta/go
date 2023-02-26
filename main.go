package main

import (
	"fmt"
	"github.com/vitormoschetta/go/calc"
	"github.com/vitormoschetta/go/services"
)

func main() {
	calc.Execute()
	services.Execute()
	fmt.Println("Fim!")
}
