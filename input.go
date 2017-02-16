package main

//import "time"
import (
    "fmt"
    "github.com/nsf/termbox-go"
)

var buf chan rune

func readKeybord() string {
     err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()
///////////////////////////////////
    buf = make(chan rune, 16)
    var input string
    go iLoop()
    i := 0
    key := rune(0)
    for key != 13 {
        if KeyPressed() {
            key = ReadKey()
            input = input + string(key)
            fmt.Printf("%c:%d globKey - %c \n counter= %d", key, key, key, i)
            fmt.Println("")
            i = 0
        }
        //fmt.Printf("%d\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08", i)
        i++
    }
    //fmt.Println("stop")
    lenth := len(input)
    input = input[:lenth - 1]
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