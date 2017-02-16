package main

import (
	"fmt"
    "strings"
    "github.com/nsf/termbox-go" 
	//"math/rand"
	//"time"
    //"bufio"
    //"os"
	"time"
)

/*var comm1 string
var comm2 string
var comm3 string
var comm4 string*/

//var matrixActionList = createMatrixActionList()



func formCommand(iconSource Icon, iconTarget Icon) ([]string) {
    //fmt.Println("Формируем команду")
    var comm []string
    comm1 := iconSource.getIconName()
    comm2 := "HACK" //рандомное из допущеных действий
    comm3 := iconTarget.getIconName()
    command := comm1 + ">" + comm2 + ">" + comm3
    //command = strings.ToUpper(command)
    //outputRed(command)
    //comm := strings.Split(command, ">")
    comm = strings.SplitN(command, ">",4)
    checkCommand(command)
    return comm
}

func userInput () ([]string) {
    var comm []string
    var command string
    inputLoop := false
    for inputLoop == false {
        var input string
        fmt.Println(masterIconList.iconArray)
        input = readKeybord()
        //fmt.Scanln(&input)
        command = parseInput(input)
        //command = readInput()
        command = strings.Replace(command, " ", "_", -1)
        //outputRed(command)
        comm = strings.SplitN(command, ">",4)
        inputLoop = checkCommand(command)




		/*reader := bufio.NewReader(os.Stdin)
	    fmt.Println("Enter text: ")
		command, _ := reader.ReadString('\n')
		fmt.Println(command)
        comm = strings.SplitN(command, ">",4)
        fmt.Println(comm)
        comm[2] = strings.TrimRight(comm[2], "\n")
        inputLoop = checkCommand(comm[0], comm[1], comm[2])*/
    }
    
    return comm
    
}

func checkCommand(command string) bool {
    sourceOK := false
    actionOK := false
    targetOK := false
    comm := strings.SplitN(command, ">",4)
    if len(comm) < 3 {
        outputRed("ERROR: NOT ENOUGH DATA TO EXECUTE COMMAND...")
        outputRed("PLEASE USE NEXT FORM: [sourceIconName]>[action]>[targetIconName]>[action parameter(optional)]")
        return false
    }
    comm[0] = strings.Replace(comm[0], " ", "_", -1)
    comm[0] = strings.ToUpper(comm[0])
    comm[1] = strings.Replace(comm[1], " ", "_", -1)
    comm[1] = strings.ToUpper(comm[1])
    comm[2] = strings.Replace(comm[2], " ", "_", -1)
    comm[2] = strings.ToUpper(comm[2])
    commandOK := false
    //проверка 1 ********************************************
    sourceOK = checkSourceSpelling(comm[0])
    //проверка 2
    targetOK = false
    if comm[1] != "MATRIX_SEARCH" {
    actionOK = checkActionSpelling(comm[1])
    //проверка 3  если действие не MATRIX_SEARCH иначе все валидно*******************************************
    targetOK = checkTargetSpelling(comm[2])
    } else if comm[1] == "MATRIX_SEARCH" {
        actionOK = true
        targetOK = true
    }

    //суммиривание проверок
    if (sourceOK && actionOK && targetOK) == true {
        commandOK = true
    } 
    if sourceOK != true {
        outputRed("ERROR: SOURCE INPUT INCORRECT...")
    }
    if actionOK != true {
        outputRed("ERROR: COMMAND UNKNOWN...")
    }
    if targetOK != true {
        outputRed("ERROR: TARGET INPUT INCORRECT...")
    }
    if len(comm) == 4 {
        fmt.Println("есть четвертый элемент = ", comm[3])
    }
    if (sourceOK && actionOK && targetOK) == true {
        commandOK = true
    } 
    return commandOK
}

func checkSourceSpelling (comm1 string) bool {
    comm1check := false
    var sourceName string
    comm1 = strings.ToUpper(comm1)
    for i := range masterIconList.iconArray {
        sourceName = strings.ToUpper(masterIconList.iconArray[i].getIconName())
        if (sourceName == comm1) {
            comm1check = true
            //fmt.Println("проверка comm1check успешна")        
        }
    }
    return comm1check
}

func checkActionSpelling(comm2 string) bool {
    comm2check := false
    comm2 = strings.Replace(comm2, " ", "_", -1)
    comm2 = strings.ToUpper(comm2)
    for i := range matrixActionList {
        if matrixActionList[i] == comm2 {
            comm2check = true
            //fmt.Println("проверка comm2check успешна")        
        }
    }
    return comm2check
}

func checkTargetSpelling (comm3 string) bool {
     comm3check := false
     var targetName string
     comm3 = strings.ToUpper(comm3)
     comm3 = strings.TrimRight(comm3, "\n")
     comm3 = strings.Replace(comm3, "\r\n", "", -1)
     //fmt.Println("Начинаем проверку")
     if len(targetList.iconArray) == 0 {
         comm3check = true
         //outputRed("NO TARGETS SPOTTED")
     }
    for i := range targetList.iconArray {
       // fmt.Println("Шаг", i)
        targetName = strings.ToUpper(targetList.iconArray[i].getIconName())
        targetName = strings.TrimRight(targetName, "\n")
        targetName = strings.Replace(targetName, "\r\n", "", -1)
        //fmt.Println(targetName, comm3)
        //strings.TrimRight(input, "\n")
        if strings.Compare(targetName, comm3) == 0 {
        //if targetName == comm3 {
            comm3check = true
          //  fmt.Println("проверка comm3check успешна")
                    
        }
    }
    return comm3check
}

func parseInput(s string) string { 
    var value string
    if len(s) == 0 { 
        value = "-1" 
        // value = 10 
        // fmt.Printf("Default option [%d] selected\n\n", value) 
        return value 
    } 
    for _, ch := range s { 
        if ch >= 19 && ch <= 122 { 
            value = value + string(ch)
        } else { 
            value = "-1 "
            break 
        } 
    } 
    return value 
}

func outputRed (s string) {
    var letter []string 
    letter = strings.Split(s, "")
    for i := range letter {
    output := "\033[31m" + strings.ToUpper(letter[i]) + "\033[0m"
    fmt.Print(output)    
    time.Sleep(time.Millisecond * 18)
    }
    fmt.Println("")
    
}

func readInput () string {
    var input string
    err := termbox.Init() 
    if err != nil { 
        panic(err) 
    } 
    defer termbox.Close() 

    termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse) 

    loop: 
    for { 
        switch ev := termbox.PollEvent(); ev.Type { 
            case termbox.EventKey:
            input = input + string(ev.Ch)
            //fmt.Printf("%c - %d\n", ev.Ch, ev.Key) 
            if ev.Key == termbox.KeyEnter { 
            break loop 
        } 
    //case termbox.EventResize: 
    //fmt.Println(ev.Width, ev.Height) 

    } 
} 



    return input
}