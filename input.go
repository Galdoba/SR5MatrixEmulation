package main

//import "time"
import (
	"fmt"

	"os"

	"github.com/nsf/termbox-go"
)

var buf chan rune

func readKeybord() string {
	///////////////////////////////////
	buf = make(chan rune, 16)
	var input string
	go iLoop()
	i := 0
	key := rune(0)
	for key != 13 {
		fmt.Printf("%s    \x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08", input)
		if KeyPressed() {
			key = ReadKey()
			input = input + string(key)
			if key == 27 { //при нажатии эскейпа выходим из программы
				fmt.Println("Exit on demand!")
				os.Exit(13)
			}
			if key == 8 { //при нажатии бэкспэйса удаляем одну букву
				if len(input) > 1 {
					input = input[:len(input)-2]
				} else {
				}
				key = 0
			}
			getInputString(input)
			i = 0
		}
		i++
	}
	input = input[:len(input)-1]
	fmt.Println("return string: " + input)
	return input
}

func runeTranslator() rune {

	var key rune
	ev := termbox.PollEvent()
	switch ev.Type {
	case termbox.EventKey:
		key = rune(ev.Key)
		if ev.Ch != 0 {
			key = ev.Ch
		}
	}
	return key
}

func ReadKey() rune {
	return <-buf
}

func KeyPressed() bool {
	return len(buf) > 0
}

func iLoop() {
	for {
		buf <- runeTranslator()
	}
}

func getInputString(input string) string {
	if len(input) > 1 {
		input = input[:len(input)-1]

	}
	//fmt.Println("return string: " + input)
	return input
}
