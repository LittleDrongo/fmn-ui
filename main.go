package main

import (
	"fmt"

	"github.com/LittleDrongo/fmn-ui/lib/ui"
	"github.com/LittleDrongo/go_libs/console/color"
)

func main() {
	QueteSample()

}

func QueteSample() {
	str := `Настоящая свобода имеется только там, где уничтожена эксплуатация, где нет угнетения одних людей другими, где нет безработицы и нищенства, где человек не дрожит за то, что завтра может потерять работу, жилище, хлеб.
	
	—  Иосиф Виссарионович Сталин`

	fmt.Println()
	ui.PrintQuote(str, 120, color.GREEN, color.RESET)
	fmt.Println()

}
