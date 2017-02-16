package main

//import "time"
import (
	"fmt"

	"os"

	"github.com/nsf/termbox-go"
)

var buf chan rune

func readKeybord() string {
	/*err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()*/
	///////////////////////////////////
	buf = make(chan rune, 16)
	var input string
	lenth := len(input)
	go iLoop()
	i := 0
	key := rune(0)
	for key != 13 {
		if KeyPressed() {
			key = ReadKey()
			if key == 27 { //при нажатии эскейпа выходим из программы
				fmt.Println("Exit on demand!")
				os.Exit(13)
			}
			if key == 8 { //при нажатии бэкспэйса удаляем одну букву
				input = input[:lenth]
				fmt.Print(input, "\x08")
			}
			input = input + string(key)
			fmt.Print(string(key))
			//fmt.Printf("%c:%d globKey - %c \n counter= %d", key, key, key, i)
			//fmt.Println("")
			getInputString(input)
			i = 0
		}
		//fmt.Printf("%d\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08", i)
		i++
	}
	fmt.Println("stop")
	//lenth := len(input)
	//input = input[:lenth-1]
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
	lenth := len(input)
	input = input[:lenth-1]
	//fmt.Println("return string: " + input)
	return input
}
