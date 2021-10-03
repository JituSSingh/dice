package diceprinter

import (
	"fmt"

	"github.com/jitussingh/dice"
)

func PrintRoll(sides int, comment string) {
	fmt.Println("%s: %d\n", comment, dice.Roll(sides))
}
