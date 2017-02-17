//package main

//import "fmt"

//import "time"
//import "strings"

//func main() {

//ввод интеджера
/*fmt.Print("Enter integer of characters: ")
  var input int
  _, err := fmt.Scanf("%d", &input)
  if err != nil {
   fmt.Println("Error in integer input")
  }*/
//  createRooster()
//go station()

//	fmt.Println("******************************************")

//}

package main

import (
	tm "github.com/buger/goterm"
	"github.com/nsf/termbox-go"
)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	tm.Clear()
	//readKeybord()
	// Create Box with 30% width of current screen, and height of 20 lines
	//playerInfo := tm.NewBox(20|tm.PCT, 80|tm.PCT, 0)
	//playerInfo.Border = "— | ┌ ┐ └ ┘"

	// Add some content to the box
	// Note that you can add ANY content, even tables
	//fmt.Fprint(playerInfo, "Some box content")
	station()
	createRooster()

	// Move Box to approx center of the screen
	//tm.Print(tm.MoveTo(playerInfo.String(), 0, 1))
	//tm.Println(tm.Height())
	/*outputBox("test")
	fmt.Println(tm.Height())
		tm.Flush()
		fmt.Println(tm.Height())
		fmt.Println(tm.Width())*/
	defer termbox.Close()
}

/*func outputBox (s string, ) {
	//fmt.Println = tm.Height()
	playerInfo := tm.NewBox(20|tm.PCT, 150|tm.PCT, 0)
	playerInfo.Border = "— | ┌ ┐ └ ┘"
	//tm.Println(150|tm.PCT)
	s = "lalalal kjfdhgkdfjgh fgksdfhg"
	//s1 := split
	//fmt.Println(tm.Height())
	fmt.Fprint(playerInfo, s)
    //output := "\033[32m" + strings.ToUpper(letter[i]) + "\033[0m"
    time.Sleep(time.Millisecond * 18)
	tm.Print(tm.MoveTo(playerInfo.String(), 0, 1))
    tm.Println("")
	//fmt.Println(tm.Height())

}*/
